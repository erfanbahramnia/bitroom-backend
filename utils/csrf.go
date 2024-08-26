package utils

import (
	"encoding/base64"
)

func GenerateCSRFToken() (string, error) {
	b := make([]byte, 32)
	return base64.URLEncoding.EncodeToString(b), nil
}
