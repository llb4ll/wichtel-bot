package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
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
func (wichtelGroup WichtelGroup) drawRandomWichtel() (wichtel Wichtel, remainingWichtelGroup WichtelGroup) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randomIndex = r.Intn(len(wichtelGroup))
	wichtel = wichtelGroup[randomIndex]
	remainingWichtelGroup = append(wichtelGroup[:randomIndex], wichtelGroup[randomIndex + 1:]...)
	return
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

func (wichtelMap *WichtelMap) assignWichtel(wichtelGroup WichtelGroup) (bool) {

	var santaGroup = append([]Wichtel(nil), wichtelGroup...)

	for len(wichtelGroup) > 0 {
		var wichtel Wichtel;
		var santa Wichtel;

		wichtel, wichtelGroup = wichtelGroup.drawRandomWichtel()
		santa, santaGroup = santaGroup[0], santaGroup[1:]
		if !santa.checkWichtel(wichtel) {
			return false
		}

		(*wichtelMap)[&santa] = wichtel
	}

	return true
}

func (santa Wichtel) checkWichtel(wichtel Wichtel) (bool) {
	if (santa.Email == wichtel.Email) {
		return false
	}

	for _, mail := range santa.ExcludeWichtelByEmail {
		if mail == wichtel.Email {
			return false
		}
	}

	return true
}