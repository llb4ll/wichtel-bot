package main

import "testing"

func TestWichtel(t *testing.T) {

	// construct new wichtel
	wichtel1 := Wichtel{"Test Wichtel", "testmail@test.de"}

	wichtelName := wichtel1.Name
	if(wichtelName) != "Test Wichtel"{
		t.Error("Expected Test Wichtel, got ", wichtelName)
	}

	wichtelEmail := wichtel1.Email
	if(wichtelEmail) != "testmail@test.de"{
		t.Error("Expected testmail@test.de, got ", wichtelEmail)
	}
}