package main

// Example application using the github.com/briandowns/spinner API

import (
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner
	time.Sleep(4 * time.Second)                                 // Run for some time to simulate work

	s.UpdateCharSet(spinner.CharSets[9])  // Update spinner to use a different character set
	s.UpdateSpeed(100 * time.Millisecond) // Update the speed the spinner spins at
	s.Restart()                           // Restart the spinner
	time.Sleep(4 * time.Second)

	s.Reverse() // Reverse the direction the spinner is spinning
	s.Restart()
	time.Sleep(4 * time.Second)

	s.Prefix = "prefixed text: " // Prefix text before the spinner
	time.Sleep(4 * time.Second)
	s.Prefix = ""
	s.Suffix = "  :appended text" // Append text after the spinner
	time.Sleep(4 * time.Second)
	s.Stop() // Stop the spinner
}
