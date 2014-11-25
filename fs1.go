package fs1

import (
	"io/ioutil"
	"net/http"

	"github.com/peronix/go-simplejson"
	"github.com/peronix/oauth"
)

type FsOneInterface struct {
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

func makeRequest(response *http.Response, err error) (*simplejson.Json, error) {
	if err != nil {
		return &simplejson.Json{}, err
	}
	defer response.Body.Close()
	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &simplejson.Json{}, err
	}
	json, err := simplejson.NewJson(bits)
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

	json, err := makeRequest(fs.consumer.Get(
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

func (fs *FsOneInterface) FindPerson(name, address string) (string, error) {
	url := fs.basePath + "/v1/People/Search.json"

	json, err := makeRequest(fs.consumer.Get(
		url, map[string]string{
			"searchFor": name,
			"address":   address,
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

func (fs *FsOneInterface) NewHousehold() (*simplejson.Json, error) {
	return fs.newObject("/v1/Households/New.json")
}

func (fs *FsOneInterface) NewPerson() (*simplejson.Json, error) {
	return fs.newObject("/v1/People/New.json")
}

func (fs *FsOneInterface) NewAddress() (*simplejson.Json, error) {
	return fs.newObject("/v1/Addresses/New.json")
}

func (fs *FsOneInterface) NewCommunication() (*simplejson.Json, error) {
	return fs.newObject("/v1/Communications/New.json")
}

func (fs *FsOneInterface) NewContribution() (*simplejson.Json, error) {
	return fs.newObject("/giving/v1/contributionreceipts/new.json")
}

func (fs *FsOneInterface) newObject(requestUrl string) (*simplejson.Json, error) {
	url := fs.basePath + requestUrl

	json, err := makeRequest(fs.consumer.Get(
		url, map[string]string{}, fs.accessToken,
	))
	if err != nil {
		return &simplejson.Json{}, err
	}
	return json, nil
}

func (fs *FsOneInterface) CreateHousehold(data *simplejson.Json) (string, error) {
	return fs.createObject("/v1/Households.json", "person", data)
}

func (fs *FsOneInterface) CreatePerson(data *simplejson.Json) (string, error) {
	return fs.createObject("/v1/People.json", "person", data)
}

func (fs *FsOneInterface) CreateAddress(data *simplejson.Json) (string, error) {
	return fs.createObject("/v1/Addresses.json", "address", data)
}

func (fs *FsOneInterface) CreateCommunication(data *simplejson.Json) (string, error) {
	return fs.createObject("/v1/Communications.json", "communication", data)
}

func (fs *FsOneInterface) CreateContribution(data *simplejson.Json) (string, error) {
	return fs.createObject("/giving/v1/contributionreceipts.json", "contributionReceipt", data)
}

func (fs *FsOneInterface) createObject(requestUrl, objectName string, data *simplejson.Json) (string, error) {
	url := fs.basePath + requestUrl

	dataStr, err := data.Encode()
	if err != nil {
		return "", err
	}
	json, err := makeRequest(fs.consumer.PostJson(
		url, string(dataStr), fs.accessToken,
	))
	if err != nil {
		return "", err
	}
	return json.GetPath(objectName, "@id").MustString(""), nil
}
