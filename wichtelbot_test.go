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

	if (len(wichtelGroup) != 3) {
		t.Error("Expected 3 wichtel, got ", wichtelGroup)
	}
}

func TestWichtelAssignment(t *testing.T) {
	wichtel1 := Wichtel{"Test Wichtel 1", "testmail1", nil}
	wichtel2 := Wichtel{"Test Wichtel 2", "testmail2", nil}

	wichtelGroup := WichtelGroup{wichtel1, wichtel2}

	wichtelMap := make(WichtelMap)
	wichtelMap.assignWichtel(wichtelGroup)

	if(len(wichtelGroup) != len(wichtelMap)){
		t.Error("Assignment of wichtel failed - not everybody got a wichtel: ", wichtelGroup, wichtelMap)
	}

	for wichtelPointer, wichtel := range wichtelMap{
		if((*wichtelPointer).Email == wichtel.Email){
			t.Error("Assignment failed - secret santa was assigned himself as a wichtel: ", (*wichtelPointer).Name)
		}
	}
}