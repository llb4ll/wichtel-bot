package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
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

func (wichtelGroup *WichtelGroup) readWichtelInput(filePath string) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

	jsonErr := json.Unmarshal(content, wichtelGroup)

	if jsonErr != nil {
		log.Fatal("Wichtel unmarshal failed: ", jsonErr)
	}
}
