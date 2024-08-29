package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xlzd/gotp"
)

const Version string = "v0.1"

func checkTOTP(s string) string {
	totp := gotp.NewDefaultTOTP(s)
	totpResult := totp.Now()

	return totpResult
}

func main() {
	var authSecret string
	var callVersion bool
	flag.StringVar(&authSecret, "secret", authSecret, "the auth secret (ex: ABCDEFGHIJKLMNOP)")
	flag.BoolVar(&callVersion, "version", false, "displays version and exits")
	flag.Parse()

	if callVersion {
		fmt.Printf("totps %s\n", Version)
		os.Exit(0)
	}

	if len(authSecret) <= 1 {
		println("Auth secret required (minimum 2 characters)")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("TOTP: %s\n", checkTOTP(authSecret))
}
