package main

import (
	"bytes"
	"fmt"
	"log"

	gqrcode "github.com/skip2/go-qrcode"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var ssid string
	const (
		content = "WIFI:T:WPA;S:%s;P:%s;;"
		size    = 256
	)

	fmt.Print("SSID: ")
	fmt.Scanf("%s", &ssid)
	fmt.Print("Password: ")
	pw, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalf("error reading password: %v", err)
	}

	qrc, err := gqrcode.New(string(escape(pw)), gqrcode.High)
	if err != nil {
		log.Fatalf("error generating QRCode: %v", err)
	}
	err = qrc.WriteFile(size, ssid+".png")
	if err != nil {
		log.Fatalf("error creating QRCode file: %v", err)
	}

	fmt.Println("Successfully generated QRCode image" + ssid + ".png")
}

func escape(pw []byte) (escaped []byte) {
	escaped = bytes.ReplaceAll(pw, []byte("\\"), []byte("\\\\"))
	escaped = bytes.ReplaceAll(escaped, []byte(";"), []byte("\\;"))
	escaped = bytes.ReplaceAll(escaped, []byte(","), []byte("\\,"))
	escaped = bytes.ReplaceAll(escaped, []byte("\""), []byte("\\\""))
	return bytes.ReplaceAll(escaped, []byte(":"), []byte("\\:"))
}
