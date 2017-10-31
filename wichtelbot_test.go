package main

import (
	"testing"
)

func TestWichtel(t *testing.T) {

	// construct new wichtel
	wichtel1 := Wichtel{"Test Wichtel", "testmail@test.de", nil}

	wichtelName := wichtel1.Name
	if (wichtelName) != "Test Wichtel" {
		t.Error("Expected Test Wichtel, got ", wichtelName)
	}

	wichtelEmail := wichtel1.Email
	if (wichtelEmail) != "testmail@test.de" {
		t.Error("Expected testmail@test.de, got ", wichtelEmail)
	}
}

func TestWichtelJson(t *testing.T) {

	var wichtelGroup WichtelGroup

	wichtelGroup.readWichtelInput("wichtel.json")

	if (len(wichtelGroup) != 4) {
		t.Error("Expected 4 wichtel, got ", wichtelGroup)
	}
}

func TestWichtelAssignment(t *testing.T) {
	wichtel1 := Wichtel{"Test Wichtel 1", "testmail1", nil}
	wichtel2 := Wichtel{"Test Wichtel 2", "testmail2", nil}

	wichtelGroup := WichtelGroup{wichtel1, wichtel2}

	wichtelMap := make(WichtelMap)
	err, wichtelMap := wichtelGroup.assignWichtel()

	if err == nil {
		if (len(wichtelGroup) != len(wichtelMap)) {
			t.Error("Assignment of wichtel failed - not everybody got a wichtel: ", wichtelGroup, wichtelMap)
		}

		for santa, wichtel := range wichtelMap {
			if ((*santa).Email == wichtel.Email) {
				t.Error("Assignment failed - secret santa was assigned himself as a wichtel: ", (*santa).Name)
			}
		}
	}
}

func TestCheckWichtel(t *testing.T) {
	santa := Wichtel{"Test Santa", "testsantamail", nil}
	santaWithRestrictions := Wichtel{"Test Santa", "testsantamail", []string{"testwichtelmail"}}
	wichtel := Wichtel{"Test Wichtel", "testwichtelmail", nil}

	if (santa.checkWichtel(santa) != false) {
		t.Error("Check wichtel failed. Expected santa not being a valid wichtel for himself")
	}

	if (santaWithRestrictions.checkWichtel(wichtel) != false) {
		t.Error("Check wichtel failed. Expected santa not being a valid santa for wichtel in his restricitons list.")
	}

	if (santa.checkWichtel(wichtel) != true) {
		t.Error("Check wichtel failed. Expected santa to be a match for other wichtel")
	}
}