package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAcount(test *testing.T) {
	acnt, err := newAccount("firstName", "lastName", "testNewAccount")
	assert.Nil(test, err)

	fmt.Printf("%+v\n", acnt)
}
