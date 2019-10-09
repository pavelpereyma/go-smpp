package pdutext

import (
	"github.com/pavelpereyma/go-smpp/smpp/encoding"
	"golang.org/x/text/transform"
)

// FlashSMS flash
type FlashSMS []byte

// Type implements the Codec interface.
func (s FlashSMS) Type() DataCoding {
	return 0xc0
}

// Encode to GSM 7-bit (unpacked)
func (s FlashSMS) Encode() []byte {
	e := encoding.GSM7(false).NewEncoder()
	es, _, err := transform.Bytes(e, s)
	if err != nil {
		return s
	}
	return es
}

// Decode from GSM 7-bit (unpacked)
func (s FlashSMS) Decode() []byte {
	e := encoding.GSM7(false).NewDecoder()
	es, _, err := transform.Bytes(e, s)
	if err != nil {
		return s
	}
	return es
}
