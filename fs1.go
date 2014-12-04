package fs1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/peronix/go-simplejson"
	"github.com/peronix/oauth"
)

type FsOneInterface struct {
	debug        bool
	churchCode   string
	basePath     string
	callbackUrl  string
	consumer     *oauth.Consumer
	accessToken  *oauth.AccessToken
	requestToken *oauth.RequestToken
}

type Fund struct {
	Id   string
	Name string
}

func NewFsOneInterface(consumerKey, consumerSecret, consumerCode, callbackUrl string, debug bool) (fs FsOneInterface) {
	fs.debug = debug
	fs.callbackUrl = callbackUrl
	fs.basePath = "https://" + consumerCode + ".fellowshiponeapi.com"
	fs.consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://" + consumerCode + ".fellowshiponeapi.com/v1/Tokens/RequestToken",
			AuthorizeTokenUrl: "https://" + consumerCode + ".fellowshiponeapi.com/v1/PortalUser/Login",
			AccessTokenUrl:    "https://" + consumerCode + ".fellowshiponeapi.com/v1/Tokens/AccessToken",
		},
	)
	fs.consumer.Debug(debug)
	return fs
}

func (fs *FsOneInterface) makeRequest(response *http.Response, err error) (*simplejson.Json, error) {
	if err != nil {
		if fs.debug {
			fmt.Printf("Respone: %#v\n", err.Error())
		}
		return &simplejson.Json{}, err
	}
	if fs.debug {
		fmt.Printf("Respone: %#v\n", response)
	}
	json, err := simplejson.NewFromReader(response.Body)
	if err != nil {
		return &simplejson.Json{}, err
	}
	return json, nil
}

func (fs *FsOneInterface) SetChurchCode(code string) {
	fs.churchCode = code
	fs.basePath = "https://" + fs.churchCode + ".fellowshiponeapi.com"
}

func (fs *FsOneInterface) SetAccessToken(token, secret string) {
	fs.accessToken = &oauth.AccessToken{
		Token:  token,
		Secret: secret,
	}
}

func (fs *FsOneInterface) SetRequestToken(token, secret string) {
	fs.requestToken = &oauth.RequestToken{
		Token:  token,
		Secret: secret,
	}
}

func (fs *FsOneInterface) GetRequestTokenAndUrl() (string, string, string, error) {
	token, url, err := fs.consumer.GetRequestTokenAndUrl(fs.callbackUrl)
	if err != nil {
		return "", "", "", err
	}
	fs.requestToken = token
	return token.Token, token.Secret, url, nil
}

func (fs *FsOneInterface) GetAccessToken(verificationCode string) (string, string, error) {
	token, err := fs.consumer.AuthorizeToken(fs.requestToken, verificationCode)
	if err != nil {
		return "", "", err
	}
	fs.accessToken = token
	return token.Token, token.Secret, nil
}

func (fs *FsOneInterface) GetFundList() ([]Fund, error) {
	url := fs.basePath + "/giving/v1/funds.json"

	json, err := fs.makeRequest(fs.consumer.Get(
		url, map[string]string{}, fs.accessToken,
	))
	if err != nil {
		return nil, err
	}

	funds := make([]Fund, 0)
	json.GetPath("funds", "fund").ArrayEach(func(v *simplejson.Json) {
		funds = append(funds, Fund{
			Id:   v.Get("@id").MustString(""),
			Name: v.Get("name").MustString(""),
		})
	})

	return funds, nil
}

func (fs *FsOneInterface) FindPerson(name, email string) (string, error) {
	url := fs.basePath + "/v1/People/Search.json"

	json, err := fs.makeRequest(fs.consumer.Get(
		url, map[string]string{
			"searchFor":     name,
			"communication": email,
		}, fs.accessToken,
	))
	if err != nil {
		return "", err
	}

	json = json.Get("results")

	recordCount := json.Get("@totalRecords").MustIntFromString(0)
	if recordCount < 1 {
		return "", nil
	}

	return json.Get("person").GetIndex(0).Get("@id").MustString(""), nil
}

func (fs *FsOneInterface) CreateHousehold(data *Household) error {
	return fs.createObject("/v1/households", &data)
}

func (fs *FsOneInterface) CreatePerson(data *Person) error {
	return fs.createObject("/v1/people", &data)
}

func (fs *FsOneInterface) CreateAddress(data *Address) error {
	return fs.createObject("/v1/addresses", &data)
}

func (fs *FsOneInterface) CreateCommunication(data *Communication) error {
	return fs.createObject("/v1/communications", &data)
}

func (fs *FsOneInterface) CreateContribution(data *ContributionReceipt) error {
	return fs.createObject("/giving/v1/contributionreceipts", &data)
}

func (fs *FsOneInterface) createObject(requestUrl string, data interface{}) error {
	url := fs.basePath + requestUrl + ".json"

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if fs.debug {
		fmt.Println("\n" + string(dataBytes))
	}
	r, err := fs.consumer.PostJson(
		url, string(dataBytes), fs.accessToken,
	)
	defer r.Body.Close()
	if fs.debug {
		fmt.Printf("Respone: %#v\n", r)
	}
	bits, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bits, &data)
	return err
}

func (fs *FsOneInterface) EditHousehold(id string) (Household, error) {
	data := Household{}
	err := fs.editObject("/v1/households", id, &data)
	return data, err
}

func (fs *FsOneInterface) EditPerson(id string) (Person, error) {
	data := Person{}
	err := fs.editObject("/v1/people", id, &data)
	return data, err
}

func (fs *FsOneInterface) EditAddress(id string) (Address, error) {
	data := Address{}
	err := fs.editObject("/v1/addresses", id, &data)
	return data, err
}

func (fs *FsOneInterface) EditCommunication(id string) (Communication, error) {
	data := Communication{}
	err := fs.editObject("/v1/communications", id, &data)
	return data, err
}

func (fs *FsOneInterface) EditContribution(id string) (ContributionReceipt, error) {
	data := ContributionReceipt{}
	err := fs.editObject("/giving/v1/contributionreceipts", id, &data)
	return data, err
}

func (fs *FsOneInterface) editObject(requestUrl, id string, object interface{}) error {
	url := fs.basePath + requestUrl + "/" + id + "/edit.json"

	r, err := fs.consumer.Get(url, map[string]string{}, fs.accessToken)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if fs.debug {
		fmt.Printf("Respone: %#v\n", r)
	}
	bits, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bits, &object)

	return err
}

func (fs *FsOneInterface) UpdateHousehold(data *Household) error {
	return fs.updateObject("/v1/households", data.Household.Id, &data)
}

func (fs *FsOneInterface) UpdatePerson(data *Person) error {
	return fs.updateObject("/v1/people", data.Person.Id, &data)
}

func (fs *FsOneInterface) UpdateAddress(data *Address) error {
	return fs.updateObject("/v1/addresses", data.Address.Id, &data)
}

func (fs *FsOneInterface) UpdateCommunication(data *Communication) error {
	return fs.updateObject("/v1/communications", data.Communication.Id, &data)
}

func (fs *FsOneInterface) UpdateContribution(data *ContributionReceipt) error {
	return fs.updateObject("/giving/v1/contributionreceipts", data.ContributionReceipt.Id, &data)
}

func (fs *FsOneInterface) updateObject(requestUrl, id string, data interface{}) error {
	url := fs.basePath + requestUrl + "/" + id + ".json"

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if fs.debug {
		fmt.Println("\n" + string(dataBytes))
	}
	r, err := fs.consumer.PostJson(
		url, string(dataBytes), fs.accessToken,
	)
	defer r.Body.Close()
	if fs.debug {
		fmt.Printf("Respone: %#v\n", r)
	}
	bits, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bits, &data)
	return err
}
