package utils

import "errors"

var session = ""

func Login(username, password string) (string, error) {
	if username == "admin" && password == "123" {
		session = "admin-session"
		return session, nil
	}
	return "", errors.New("invalid credentials")
}

func CheckSession(s string) bool {
	return s == session
}
