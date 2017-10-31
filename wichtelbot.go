package main

import "fmt"

// A wichtel for the wichtel draw - wichtel are identified by their email
type Wichtel struct {
	Name  string
	Email string
	ExcludeWichtel []Wichtel
}

// Main function running the wichtel draw
func main() {
	fmt.Println("Hello wichtel admin!")
}

// Draw a random wichtel from the wichtel pool
func drawRandomWichtel() {

}

// Send out wichtel emails
func sendEmail() {

}