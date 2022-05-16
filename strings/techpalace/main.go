/*
	https://exercism.org/tracks/go/exercises/welcome-to-tech-palace
*/
package techpalace
import (
    "strings"
    "fmt"
)
const WelcomeMsg = "Welcome to the Tech Palace"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return WelcomeMsg +", "+ strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var b strings.Builder

    fmt.Fprintf(&b, 
        "%s\n%s\n%s",
        strings.Repeat("*", numStarsPerLine),
    	welcomeMsg,
    	strings.Repeat("*", numStarsPerLine),
        )
    return b.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim( strings.ReplaceAll(oldMsg, "*", ""), " \n")
}
