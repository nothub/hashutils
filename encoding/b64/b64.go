// Package b64 does standard Base64 (RFC 4648, sect 4) encoding and decoding,
// except that padding (=) signs are omitted, and whitespace are not allowed.
// b64 is intended to be used for handling PHC format serialized strings.
// https://github.com/P-H-C/phc-string-format/blob/master/phc-sf-spec.md#b64
package b64

import (
	"encoding/base64"
	"strings"
)

var encoding = base64.RawStdEncoding

func Encode(bytes []byte) string {
	buf := make([]byte, encoding.EncodedLen(len(bytes)))
	encoding.Encode(buf, bytes)

	return strings.TrimSpace(string(buf))
}

func Decode(enc string) ([]byte, error) {
	bytes := []byte(enc)

	buf := make([]byte, encoding.DecodedLen(len(bytes)))
	_, err := encoding.Decode(buf, bytes)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
