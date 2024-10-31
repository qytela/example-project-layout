package utils

import (
	"encoding/base64"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ValidateRequest(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	return nil
}

func IsPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func IsFileExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func Base64Encode(text string) string {
	var encodedText = make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(encodedText, []byte(text))

	return string(encodedText)
}

func Base64Decode(encodedText []byte) (string, error) {
	var decodedText = make([]byte, base64.StdEncoding.DecodedLen(len(encodedText)))
	if _, err := base64.StdEncoding.Decode(decodedText, encodedText); err != nil {
		return "", err
	}

	return string(decodedText), nil
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ComparePassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}

func Join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		result += s
		if i < len(strs)-1 {
			result += sep
		}
	}

	return result
}
