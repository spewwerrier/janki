package db

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	username1        = "dummyuser1"
	username2        = "dummyuser2"
	username3        = "dummyuser3"
	knob_description = "A theme of Arknights I like very much"
)

func TestKnobDuplications(t *testing.T) {
	d := ConnectTestInstance(t)
	defer d.Close()

	api1, err := d.CreateNewUser(username1, passwd)
	if err != nil {
		t.Fatalf("failed to create new user %v", err)
		return
	}
	k := KnobDescription{
		Knob: Knob{
			KnobName: "Understanding atomic configuration of atomic bombs",
			IsPublic: true,
		},
	}

	_, err = d.CreateNewKnob(api1, k)
	if err != nil {
		t.Fatalf("failed to create new knob %v", err)
		return
	}

	// user creating a knob with same name twice results in error
	_, err = d.CreateNewKnob(api1, k)
	if err == nil {
		t.Fatal("should throw multiple knobs exists but did not")
		return
	}

	api2, err := d.CreateNewUser(username2, passwd)
	if err != nil {
		t.Fatal(err)
		return
	}

	// another user creating knob with same name as other user does not results in error
	_, err = d.CreateNewKnob(api2, k)
	if err != nil {
		t.Fatalf("failed to create new knob %v", err)
		return
	}
	KnobForkTests(t, d, api1, api2)
}

func TestKnob(t *testing.T) {
	d := ConnectTestInstance(t)
	defer d.Close()
	api, err := d.CreateNewUser(username3, passwd)
	if err != nil {
		t.Fatal(err)
		return
	}

	send_knob := KnobDescription{
		Knob: Knob{
			KnobName: "siracusano II",
			IsPublic: true,
		},
	}
	_, err = d.CreateNewKnob(api, send_knob)
	if err != nil {
		t.Fatal(err)
		return
	}

	recv_knob, err := d.GetUserKnobs(api)
	if err != nil {
		t.Fatalf("failed to get user knob %v", err)
		return
	}
	if recv_knob[0].KnobName != "siracusano II" {
		t.Fatalf("failed to verify knobs")
		return
	}

	knob_id, err := d.GetKnobId(api, "siracusano II")
	if err != nil {
		t.Fatal(err)
		return
	}

	identifier := recv_knob[0].Identifier

	test_knob_id, err := d.GetKnobIdFromIdentifier(identifier)
	if err != nil {
		t.Fatal(err)
		return
	}

	if knob_id != test_knob_id {
		t.Fatal("id should be same")
		return
	}

	err = d.UpdateKnobDescriptions(api, recv_knob[0].Identifier, "description", knob_description)
	if err != nil {
		t.Fatal(err)
		return
	}

	k, err := d.GetKnobDescriptions(api, identifier)
	if err != nil {
		t.Fatal(err)
		return
	}

	if k.Description != knob_description {
		t.Fatalf("WTMOOO %s", k.Description)
		return
	}

	err = d.UpdateKnobDescriptions(api, recv_knob[0].Identifier, "ques", "how to cook a carrot")
	if err != nil {
		t.Fatal("failed to update ques property")
	}
	err = d.UpdateKnobDescriptions(api, recv_knob[0].Identifier, "ques", "how to cook a carrot part 2")
	if err != nil {
		t.Fatal("failed to update ques property")
	}

	k, err = d.GetKnobDescriptions(api, recv_knob[0].Identifier)
	if err != nil {
		t.Fatal(err)
	}
	var v pgtype.Array[string]

	v.Elements = []string{
		"how to cook a carrot",
		"how to cook a carrot part 2",
	}

	for i := range len(v.Elements) {
		if v.Elements[i] != k.Ques.Elements[i] {
			t.Fatal("knob elements are not same, they are not updated")
		}
	}
}

func KnobForkTests(t *testing.T, db *Database, api1 string, api2 string) {
	k := KnobDescription{
		Knob: Knob{
			KnobName: "Learning C++",
			IsPublic: true,
		},
		Description: "By user 1",
	}
	identifier, err := db.CreateNewKnob(api1, k)
	if err != nil {
		t.Fatal(err)
	}

	err = db.ForkKnob(api2, identifier)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.GetKnobId(api2, "Learning C++")
	if err != nil {
		t.Fatalf("failed to get knobid: %v", err)
	}
}
