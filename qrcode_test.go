// go-qrcode
// Copyright 2014 Tom Harwood

package qrcode

import (
	"strings"
	"testing"
)

func TestQRCodeMaxCapacity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestQRCodeCapacity")
	}

	tests := []struct {
		char     byte
		maxChars int
	}{
		{
			'0',
			7089,
		},
		{
			'A',
			4296,
		},
		{
			'#',
			2953,
		},
	}

	for _, test := range tests {
		_, err := New(strings.Repeat(string(test.char), test.maxChars), Low)

		if err != nil {
			t.Errorf("%d x '%c' got %s expected success", test.maxChars, test.char, err.Error())
		}
	}

	for _, test := range tests {
		_, err := New(strings.Repeat(string(test.char), test.maxChars+1), Low)

		if err == nil {
			t.Errorf("%d x '%c' chars encodable, expected not encodable",
				test.maxChars+1, test.char)
		}
	}
}

func TestQRCodeVersionCapacity(t *testing.T) {
	tests := []struct {
		version         int
		level           RecoveryLevel
		maxNumeric      int
		maxAlphanumeric int
		maxByte         int
	}{
		{
			1,
			Low,
			41,
			25,
			17,
		},
		{
			2,
			Low,
			77,
			47,
			32,
		},
		{
			2,
			Highest,
			34,
			20,
			14,
		},
		{
			40,
			Low,
			7089,
			4296,
			2953,
		},
		{
			40,
			Highest,
			3057,
			1852,
			1273,
		},
	}

	for i, test := range tests {
		numericData := strings.Repeat("1", test.maxNumeric)
		alphanumericData := strings.Repeat("A", test.maxAlphanumeric)
		byteData := strings.Repeat("#", test.maxByte)

		var n *QRCode
		var a *QRCode
		var b *QRCode
		var err error

		n, err = New(numericData, test.level)
		if err != nil {
			t.Fatal(err.Error())
		}

		a, err = New(alphanumericData, test.level)
		if err != nil {
			t.Fatal(err.Error())
		}

		b, err = New(byteData, test.level)
		if err != nil {
			t.Fatal(err.Error())
		}

		if n.VersionNumber != test.version {
			t.Fatalf("Test #%d numeric has version #%d, expected #%d", i,
				n.VersionNumber, test.version)
		}

		if a.VersionNumber != test.version {
			t.Fatalf("Test #%d alphanumeric has version #%d, expected #%d", i,
				a.VersionNumber, test.version)
		}

		if b.VersionNumber != test.version {
			t.Fatalf("Test #%d byte has version #%d, expected #%d", i,
				b.VersionNumber, test.version)
		}
	}
}

func BenchmarkQRCodeMinimumSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New("1", Low)
	}
}

func BenchmarkQRCodeURLSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New("http://www.example.org", Medium)
	}
}

func BenchmarkQRCodeMaximumSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// 7089 is the maximum encodable number of numeric digits.
		New(strings.Repeat("0", 7089), Low)
	}
}
