package main

import (
	"testing"
)

func TestWichtel(t *testing.T) {

	// construct new wichtel
	wichtel1 := Wichtel{"Test Wichtel", "testmail@test.de", []string{}}

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

	wichtelGroup.readWichtelInput()

	if (len(wichtelGroup) != 1) {
		t.Error("Expected 1 wichtel, got ", wichtelGroup)
	}
}