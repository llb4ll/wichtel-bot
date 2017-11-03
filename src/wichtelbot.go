package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"errors"
	"net/smtp"
)

// wichtel-bot settings
type Settings struct {
	SmtpServer      string
	SmtpPort        string
	SmtpUser        string
	SmtpPassword    string
	SenderEmail     string
	MaximumDrawRuns int
}

// A wichtel for the wichtel draw - wichtel are identified by their email
type Wichtel struct {
	Name                  string
	Email                 string
	ExcludeWichtelByEmail []string
}

type WichtelGroup []Wichtel

type WichtelMap map[*Wichtel]Wichtel

const MAXIMUM_DRAW_RUNS = 100

// Main function running the wichtel draw
func main() {
	fmt.Println("Hello wichtel admin!")

	var settings Settings

	settings.readSettings("config/settings.json")

	var wichtelGroup WichtelGroup

	wichtelGroup.readWichtelInput("config/wichtel.json")

	fmt.Println("You are about to draw wichtel from: ", wichtelGroup)

	var runs = 1
	for ; runs < MAXIMUM_DRAW_RUNS; {
		err, wichtelMap := wichtelGroup.assignWichtel()
		if err == nil {
			fmt.Println("Wichtel successfully assigned: ")
			for santa, wichtel := range wichtelMap {
				fmt.Println((*santa).Email, wichtel.Email)
				sendEmail(settings, wichtel, *santa)
			}
			return
		}
	}

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
func sendEmail(settings Settings, wichtel Wichtel, santa Wichtel) {

	if(settings.SmtpServer == "mail.example.com"){
		fmt.Println("Please provide smtp settings to send out wichtel mails.")
		return
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", settings.SmtpUser, settings.SmtpPassword, settings.SmtpServer + ":" + settings.SmtpPort)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{santa.Email}
	msg := []byte("Dein Wichtel ist: " + wichtel.Email + "\r\n" +
		"Happy Wichteling!\r\n" +
		"\r\n" +
		"¨¨¨¨¨¨¨¨¨. *\r\n" +
		"¨¨¨¨¨¨¨¨¨ **\r\n" +
		"¨¨¨¨¨¨¨¨¨*o*\r\n" +
		"¨¨¨¨¨¨¨¨*♥*o*\r\n" +
		"¨¨¨¨¨¨¨***o***\r\n" +
		"¨¨¨¨¨¨**o**♥*o*\r\n" +
		"¨¨¨¨¨**♥**o**o**\r\n" +
		"¨¨¨¨**o**♥***♥*o*\r\n" +
		"¨¨¨*****♥*o**o****\r\n" +
		"¨¨**♥**o*****o**♥**\r\n" +
		"¨******o*****♥**o***\r\n" +
		"	****o***♥**o***o***♥ *\r\n" +
		"¨¨¨¨¨____!_!____\r\n" +
		"¨¨¨¨¨\\_________/¨¨\r\n")
	err := smtp.SendMail(settings.SmtpServer + ":" + settings.SmtpPort, auth, settings.SenderEmail, to, msg)
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Mail sent successfully.")
	}
}

func (settings *Settings) readSettings(filePath string) {

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(content, settings)

	if jsonErr != nil {
		log.Fatal("Settings unmarshal failed: ", jsonErr)
	}
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

func (inputGroup WichtelGroup) assignWichtel() (err error, wichtelMap WichtelMap) {

	wichtelMap = make(WichtelMap)
	var santaGroup = append([]Wichtel(nil), inputGroup...)
	var wichtelGroup WichtelGroup = append([]Wichtel(nil), inputGroup...)

	for len(wichtelGroup) > 0 {
		var wichtel Wichtel;
		var santa Wichtel;

		wichtel, wichtelGroup = wichtelGroup.drawRandomWichtel()
		santa, santaGroup = santaGroup[0], santaGroup[1:]
		if !santa.checkWichtel(wichtel) {
			err = errors.New("Draw did not succeed")
			return
		}

		wichtelMap[&santa] = wichtel
	}

	return
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