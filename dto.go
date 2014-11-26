package fs1

type Primary struct {
	Id  string `json:"@id"`
	Uri string `json:"@uri"`
}
type PrimaryAndName struct {
	Primary
	Name string `json:"name"`
}
type Audit struct {
	CreatedDate     string `json:"createdDate"`
	LastUpdatedDate string `json:"lastUpdatedDate"`
}

// type Security struct {
// 	LastSecurityAuthorization string `json:"lastSecurityAuthorization"`
// 	LastActivityDate          string `json:"lastActivityDate"`
// }

type Household struct {
	Household struct {
		Primary
		OldId                     string `json:"@oldID"`
		HCode                     string `json:"@hCode"`
		HouseholdName             string `json:"householdName"`
		HouseholdSortName         string `json:"householdSortName"`
		HouseholdFirstName        string `json:"householdFirstName"`
		LastSecurityAuthorization string `json:"lastSecurityAuthorization"`
		LastActivityDate          string `json:"lastActivityDate"`
		CreatedDate               string `json:"createdDate"`
		LastUpdatedDate           string `json:"lastUpdatedDate"`
	} `json:"household"`
}

type Person struct {
	Person struct {
		Primary
		ImageURI            string         `json:"@imageURI"`
		OldId               string         `json:"@oldID"`
		ICode               string         `json:"@iCode"`
		HouseholdId         string         `json:"@householdID"`
		OldHouseholdId      string         `json:"@oldHouseholdID"`
		Title               string         `json:"title"`
		Salutation          string         `json:"salutation"`
		Prefix              string         `json:"prefix"`
		FirstName           string         `json:"firstName"`
		LastName            string         `json:"lastName"`
		Suffix              string         `json:"suffix"`
		MiddleName          string         `json:"middleName"`
		GoesByName          string         `json:"goesByName"`
		FormerName          string         `json:"formerName"`
		Gender              string         `json:"gender"`
		DateOfBirth         string         `json:"dateOfBirth"`
		MaritalStatus       string         `json:"maritalStatus"`
		HouseholdMemberType PrimaryAndName `json:"householdMemberType"`
		IsAuthorized        string         `json:"isAuthorized"`
		Status              struct {
			PrimaryAndName
			Comment   string         `json:"comment"`
			Date      string         `json:"date"`
			SubStatus PrimaryAndName `json:"subStatus"`
		} `json:"status"`
		Occupation struct {
			PrimaryAndName
			Description string `json:"description"`
		} `json:"occupation"`
		Employer           string         `json:"employer"`
		School             PrimaryAndName `json:"school"`
		Denomination       PrimaryAndName `json:"denomination"`
		FormerChurch       string         `json:"formerChurch"`
		BarCode            string         `json:"barCode"`
		MemberEnvelopeCode string         `json:"memberEnvelopeCode"`
		DefaultTagComment  string         `json:"defaultTagComment"`
		Weblink            struct {
			UserId         string `json:"userID"`
			PasswordHint   string `json:"passwordHint"`
			PasswordAnswer string `json:"passwordAnswer"`
		} `json:"weblink"`
		Solicit       string `json:"solicit"`
		Thank         string `json:"thank"`
		FirstRecord   string `json:"firstRecord"`
		LastMatchDate string `json:"lastMatchDate"`
		Audit
	} `json:"person"`
}

type Address struct {
	Address struct {
		Primary
		Household               Primary        `json:"household"`
		Person                  Primary        `json:"person"`
		AddressType             PrimaryAndName `json:"addressType"`
		Address1                string         `json:"address1"`
		Address2                string         `json:"address2"`
		Address3                string         `json:"address3"`
		City                    string         `json:"city"`
		PostalCode              string         `json:"postalCode"`
		County                  string         `json:"county"`
		Country                 string         `json:"country"`
		StProvince              string         `json:"stProvince"`
		CarrierRoute            string         `json:"carrierRoute"`
		DeliveryPoint           string         `json:"deliveryPoint"`
		AddressDate             string         `json:"addressDate"`
		AddressComment          string         `json:"addressComment"`
		USPSVerified            string         `json:"uspsVerified"`
		AddressVerifiedDate     string         `json:"addressVerifiedDate"`
		LastVerificationAttempt string         `json:"lastVerificationAttempt"`
		Audit
	} `json:"address"`
}
type Communication struct {
	Communication struct {
		Primary
		Household                Primary        `json:"household"`
		Person                   Primary        `json:"person"`
		CommunicationType        PrimaryAndName `json:"communicationType"`
		CommunicationGeneralType string         `json:"communicationGeneralType"`
		ComminicationValue       string         `json:"communicationValue"`
		SearchComminicationValue string         `json:"searchCommunicationValue"`
		Preferred                string         `json:"preferred"`
		CommunicationComment     string         `json:"communicationComment"`
		Audit
	} `json:"communication"`
}

type ContributionReceipt struct {
	ContributionReceipt struct {
		Primary
		OldId               string         `json:"@oldID"`
		AccountReference    string         `json:"@accountReference"`
		Amount              string         `json:"@amount"`
		Fund                PrimaryAndName `json:"fund"`
		SubFund             PrimaryAndName `json:"subFund"`
		PledgeDrive         PrimaryAndName `json:"pledgeDrive"`
		Household           Primary        `json:"household"`
		Person              Primary        `json:"person"`
		Account             Primary        `json:"account"`
		ReferenceImage      Primary        `json:"referenceImage"`
		Batch               PrimaryAndName `json:"batch"`
		ActivityInstance    Primary        `json:"activityInstance"`
		ContributionType    PrimaryAndName `json:"contributionType"`
		ContributionSubType PrimaryAndName `json:"contributionSubType"`
		ReceivedDate        string         `json:"@receivedDate"`
		TransmitDate        string         `json:"@transmitDate"`
		ReturnDate          string         `json:"@returnDate"`
		RetransmitDate      string         `json:"@retransmitDate"`
		GlPostDate          string         `json:"@glPostDate"`
		IsSplit             string         `json:"@isSplit"`
		AddressVerification string         `json:"@addressVerification"`
		Memo                string         `json:"@memo"`
		StatedValue         string         `json:"@statedValue"`
		TrueValue           string         `json:"@trueValue"`
		Thank               string         `json:"@thank"`
		ThankedDate         string         `json:"@thankedDate"`
		IsMatched           string         `json:"@isMatched"`
		CreatedDate         string         `json:"@createdDate"`
		CreatedByPerson     Primary        `json:"createdByPerson"`
		LastUpdatedDate     string         `json:"@lastUpdatedDate"`
		LastUpdatedByPerson Primary        `json:"lastUpdatedByPerson"`
	} `json:"contributionReceipt"`
}
