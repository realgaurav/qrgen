package qr

import (
	"bytes"
	"fmt"

	gqrcode "github.com/skip2/go-qrcode"
)

const (
	contentFmt = "WIFI:T:WPA;S:%s;P:%s;;"
	size       = 256
)

// Generate creates a QR code PNG file named from the input SSID.
func Generate(ssid string, pw []byte) error {
	// WIFI:T:WPA;S:<SSID>;P:<PASSWORD>;;
	//   from https://bit.ly/3P1SP8v
	content := fmt.Sprintf(contentFmt, ssid, escape(pw))
	qrc, err := gqrcode.New(content, gqrcode.High)
	if err != nil {
		return fmt.Errorf("error generating QRCode: %w", err)
	}

	err = qrc.WriteFile(size, ssid+".png")
	if err != nil {
		fmt.Errorf("error creating QRCode file: %w", err)
	}

	return nil
}

// escape provides QR code format specific escaping.
func escape(pw []byte) (escaped []byte) {
	escaped = bytes.ReplaceAll(pw, []byte("\\"), []byte("\\\\"))
	escaped = bytes.ReplaceAll(escaped, []byte(";"), []byte("\\;"))
	escaped = bytes.ReplaceAll(escaped, []byte(","), []byte("\\,"))
	escaped = bytes.ReplaceAll(escaped, []byte("\""), []byte("\\\""))
	return bytes.ReplaceAll(escaped, []byte(":"), []byte("\\:"))
}
