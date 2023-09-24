package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	var n string

	for _, v := range phoneNumber {
		if v >= '0' && v <= '9' {
			n += string(v)
		}
	}

	l := len(n)
	switch {
	case l < 10 || l > 11:
		return "", errors.New("invalid size phone number")
	case l == 11 && n[:1] != "1":
		return "", errors.New("invalid international code")
	case n[l-10] < '2' || n[l-7] < '2':
		return "", errors.New("invalid phone number")
	}

	return n[l-10:], nil
}

func AreaCode(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return n[:3], nil
}

func Format(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
