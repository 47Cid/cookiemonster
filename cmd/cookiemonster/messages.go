package main

import (
	"fmt"
	"os"

	"github.com/iangcarroll/cookiemonster/pkg/monster"
)

var outputFile *os.File

// Say hello!
func sayHello() {
	fmt.Fprintf(outputFile, "🍪 CookieMonster %s\n", version)
}

// Output a sadder failure message if we cannot decode the cookie.
func failureMessage(message string) {
	fmt.Fprintf(outputFile, ColorRed+"❌ %s"+ColorReset+"\n", message)
	os.Exit(1)
}

// Output a nice success message if we decode the cookie.
func keyDiscoveredMessage(cookie *monster.Cookie) {
	_, key, decoder := cookie.Result()

	if isASCII(string(key)) {
		fmt.Fprintf(outputFile, ColorGreen+"✅ Success! I discovered the key for this cookie with the %s decoder; it is \"%s\".\n"+ColorReset, decoder, string(key))
	} else {
		fmt.Fprintf(outputFile, ColorGreen+"✅ Success! I discovered the key for this cookie with the %s decoder; it is (in base64): \"%s\".\n"+ColorReset, decoder, base64Key(key))
	}
}

// Output a nice success message if we decode the cookie.
func resignedMessage(out string) {
	fmt.Fprintf(outputFile, ColorGreen+"✅ I resigned this cookie for you; the new one is: %s\n"+ColorReset, out)
}

func warningMessage(message string) {
	fmt.Fprintf(outputFile, ColorYellow+"⚠️ %s\n"+ColorReset, message)
}
