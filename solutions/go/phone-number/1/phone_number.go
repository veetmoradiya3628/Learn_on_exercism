package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Number cleans up the formatting of a phone number and validates its digits.
func Number(phoneNumber string) (string, error) {
	var digits []rune
	allowedPunctuation := " +()-. \t"
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		} else if unicode.IsLetter(r) {
			return "", errors.New("letters not permitted")
		} else if !strings.ContainsRune(allowedPunctuation, r) {
			return "", errors.New("punctuations not permitted")
		}
	}
	num := string(digits)
	length := len(num)
	if length < 10 {
		return "", errors.New("must not be fewer than 10 digits")
	}
	if length > 11 {
		return "", errors.New("must not be greater than 11 digits")
	}
	if length == 11 {
		if num[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		num = num[1:]
	}

	if num[0] == '0' {
		return "", errors.New("area code cannot start with zero")
	}
	if num[0] == '1' {
		return "", errors.New("area code cannot start with one")
	}
	if num[3] == '0' {
		return "", errors.New("exchange code cannot start with zero")
	}
	if num[3] == '1' {
		return "", errors.New("exchange code cannot start with one")
	}
	return num, nil
}

// AreaCode returns the area code (first 3 digits) of a phone number.
func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format returns the phone number formatted as (NXX) NXX-XXXX.
func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}