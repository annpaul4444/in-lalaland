package main

import (
	_ "embed" // This is for HTML stuff, _ is needed when stuff is not called for whatever is imported
	"html/template"
	"log"
	"math/rand/v2" // This is for math stuff
	"os"           // This is for printing stuff
	"os/exec"
	"path/filepath"
)

// The below line is an instruction, not comment
//
//go:embed popup.html
var popupHTML string

// Ingredient - 2, Popup Stuff
var page = template.Must(template.New("popup").Parse(popupHTML))

// Ingredient - 1, BS generator
func buzzword() string {
	verb := []string{"Leverage", "Harmonize", "Accelerate", "Expedite", "Harness"}
	fancyWhats := []string{"Automation of", "Orchestration of", "Integration of", "Optimization of", "Power of"}
	fakeWhats := []string{"Agentic Vision", "Agents", "Auto-Bots", "AI Workflows", "Operational Excellence"}
	// Bit like lt_len[len]
	sentence := verb[rand.IntN(len(verb))] + " " + fancyWhats[rand.IntN(len(fancyWhats))] + " " + fakeWhats[rand.IntN(len(fakeWhats))]
	return sentence
}

// Where everything is cooked
func main() {

	text := buzzword()

	// Builds the path - tempdir is temp folder, filepath join makes sure of slashes
	path := filepath.Join(os.TempDir(), "buzz.html")
	f, err := os.Create(path)
	// bit like sy-subrc check
	if err != nil {
		log.Fatal(err) // Program stops here, no need for else
	}

	// Fill in the template onto the page
	err = page.Execute(f, text)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Asks windows to open default program which is browser, start launches it without waiting to finish
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	if err != nil {
		log.Fatal(err)
	}

}
