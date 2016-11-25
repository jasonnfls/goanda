package goanda

import (
	"errors"
)

var (
	a1 = errors.New("Account Not Found")
	a2 = errors.New("Cannot Delete Current Account")
	a3 = errors.New("No Current Account")
)

//a for account related errors
