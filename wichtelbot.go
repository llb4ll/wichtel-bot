package main

import (
	"fmt"
	"encoding/json"
)

// A wichtel for the wichtel draw - wichtel are identified by their email
type Wichtel struct {
	Name                  string
	Email                 string
	ExcludeWichtelByEmail []string
}

type WichtelGroup []Wichtel

// Main function running the wichtel draw
func main() {
	fmt.Println("Hello wichtel admin!")

}

// Draw a random wichtel from the wichtel pool
func (wichtelOut WichtelGroup) drawRandomWichtel(wichtelIn WichtelGroup) {

}

// Send out wichtel emails
func sendEmail() {

}

func (wichtelGroup *WichtelGroup) readWichtelInput() {

	jsonInput := []byte(`[{
		  "Name": "Test Wichtel 1",
		  "Email": "testwichtel1@test.de"
		}]`);

	err := json.Unmarshal(jsonInput, wichtelGroup)

	if (err != nil) {
		fmt.Println("Wichtel unmarshal failed.", err)
	}
}
