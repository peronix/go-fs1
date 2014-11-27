// +build integration

package fs1

import (
	"log"
	"testing"
)

func TestIntegration(t *testing.T) {
	f := NewFsOneInterface("", "", "", "", true)
	f.SetChurchCode("")
	f.SetAccessToken("", "")
	hd := Household{}
	hd.Household.HouseholdName = "Nathan Scott"
	//var h Household
	if hid, err := f.CreateHousehold(hd); err == nil {
		hd.Household.Id = hid
		p := Person{}
		p.Person.HouseholdMemberType.Id = "1"
		p.Person.HouseholdMemberType.Uri = "https://demo.fellowshiponeapi.com/v1/People/HouseholdMemberTypes/1"
		p.Person.HouseholdId = hd.Household.Id
		p.Person.FirstName = "Nathan"
		p.Person.LastName = "Scott"
		p.Person.Status.Id = "1"
		p.Person.IsAuthorized = "true"
		if personId, err := f.CreatePerson(p); err == nil {
			log.Println("personid", personId)
		} else {
			log.Fatalf("%d:%s", 2, err)
		}

	} else {
		log.Fatalf("%d:%s", 1, err)
	}
}
