package fs1

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestMarshalling(t *testing.T) {
	p := Person{}
	a := Address{}
	c := Communication{}

	pb, _ := json.Marshal(&p)
	ab, _ := json.Marshal(&a)
	ac, _ := json.Marshal(&c)

	if string(pb) != string(rawperson()) {
		log.Fatalf("\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", string(rawperson()), string(pb))
	}
	if string(ab) != string(rawaddress()) {
		log.Fatalf("\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", string(rawaddress()), string(ab))
	}
	if string(ac) != string(rawcommunication()) {
		log.Fatalf("\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", string(rawcommunication()), string(ac))
	}
}

func rawperson() []byte {
	raw := `{
	    "person": {
	        "@id": "",
	        "@uri": "",
	        "@imageURI": "",
	        "@oldID": "",
	        "@iCode": "",
	        "@householdID": "",
	        "@oldHouseholdID": "",
	        "title": "",
	        "salutation": "",
	        "prefix": "",
	        "firstName": "",
	        "lastName": "",
	        "suffix": "",
	        "middleName": "",
	        "goesByName": "",
	        "formerName": "",
	        "gender": "",
	        "dateOfBirth": "",
	        "maritalStatus": "",
	        "householdMemberType": {
	            "@id": "",
	            "@uri": "",
	            "name": ""
	        },
	        "isAuthorized": "",
	        "status": {
	            "@id": "",
	            "@uri": "",
	            "name": "",
	            "comment": "",
	            "date": "",
	            "subStatus": {
	                "@id": "",
	                "@uri": "",
	                "name": ""
	            }
	        },
	        "occupation": {
	            "@id": "",
	            "@uri": "",
	            "name": "",
	            "description": ""
	        },
	        "employer": "",
	        "school": {
	            "@id": "",
	            "@uri": "",
	            "name": ""
	        },
	        "denomination": {
	            "@id": "",
	            "@uri": "",
	            "name": ""
	        },
	        "formerChurch": "",
	        "barCode": "",
	        "memberEnvelopeCode": "",
	        "defaultTagComment": "",
	        "weblink": {
	            "userID": "",
	            "passwordHint": "",
	            "passwordAnswer": ""
	        },
	        "solicit": "",
	        "thank": "",
	        "firstRecord": "",
	        "lastMatchDate": "",
	        "createdDate": "",
	        "lastUpdatedDate": ""
	    }
	}`
	return stripToBytes(raw)
}
func rawaddress() []byte {
	raw := `{
	    "address": {
	        "@id": "",
	        "@uri": "",
	        "household": {
	            "@id": "",
	            "@uri": ""
	        },
	        "person": {
	            "@id": "",
	            "@uri": ""
	        },
	        "addressType": {
	            "@id": "",
	            "@uri": "",
	            "name": ""
	        },
	        "address1": "",
	        "address2": "",
	        "address3": "",
	        "city": "",
	        "postalCode": "",
	        "county": "",
	        "country": "",
	        "stProvince": "",
	        "carrierRoute": "",
	        "deliveryPoint": "",
	        "addressDate": "",
	        "addressComment": "",
	        "uspsVerified": "",
	        "addressVerifiedDate": "",
	        "lastVerificationAttemptDate": "",
	        "createdDate": "",
	        "lastUpdatedDate": ""
	    }
	}`
	return stripToBytes(raw)
}
func rawcommunication() []byte {
	raw := `{
	    "communication": {
	        "@id": "",
	        "@uri": "",
	        "household": {
	            "@id": "",
	            "@uri": ""
	        },
	        "person": {
	            "@id": "",
	            "@uri": ""
	        },
	        "communicationType": {
	            "@id": "",
	            "@uri": "",
	            "name": ""
	        },
	        "communicationGeneralType": "",
	        "communicationValue": "",
	        "searchCommunicationValue": "",
	        "preferred": "",
	        "communicationComment": "",
	        "createdDate": "",
	        "lastUpdatedDate": ""
	    }
	}`
	return stripToBytes(raw)
}
func stripToBytes(raw string) []byte {
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)
	return []byte(raw)
}
