package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/realgaurav/qrgen/pkg/qr"
)

func main() {
	var (
		ssid     string
		pw       bool
		password []byte
		err      error
	)

	flag.StringVar(&ssid, "s", "", "WiFi SSID")
	flag.BoolVar(&pw, "p", false, "Prompt for WiFi password")
	flag.Parse()

	if ssid == "" {
		fmt.Print("SSID: ")
		fmt.Scanf("%s", &ssid)
	}

	if pw {
		fmt.Print("Password: ")
		password, err = terminal.ReadPassword(0)
		if err != nil {
			log.Fatalf("error reading password: %v", err)
		}
		fmt.Println()
	} else {
		password = make([]byte, 20)
		_, err = rand.Read(password)
		if err != nil {
			log.Fatalf("error generating password: %v", err)
		}
	}

	err = qr.Generate(ssid, password)
	if err != nil {
		log.Fatalf("error generating QR code: %v", err)
	}

	fmt.Println("Successfully generated QRCode image " + ssid + ".png")
}
