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

type WichtelMap map[*Wichtel]Wichtel


// Main function running the wichtel draw
func main() {
	fmt.Println("Hello wichtel admin!")

	var wichtelGroup WichtelGroup;

	wichtelGroup.readWichtelInput("wichtel.json")

	fmt.Println("You are about to draw wichtel from: ", wichtelGroup)

	var wichtelMap WichtelMap;

	wichtelMap.assignWichtel(wichtelGroup)

}

// Draw a random wichtel from the wichtel pool
func (wichtel Wichtel) drawRandomWichtel(wichtelIn WichtelGroup) {

}

// Send out wichtel emails
func (wichtel Wichtel) sendEmail() {
	// send out wichtel mail
}

func (wichtelGroup *WichtelGroup) readWichtelInput(filePath string) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(content, wichtelGroup)

	if jsonErr != nil {
		log.Fatal("Wichtel unmarshal failed: ", jsonErr)
	}
}

func (wichtelMap *WichtelMap) assignWichtel(wichtelGroup WichtelGroup) {
	for index, wichtel := range wichtelGroup {
		(*wichtelMap)[&(wichtelGroup[index])] = wichtel
	}
}
