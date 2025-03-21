package main

import (
    "fmt"
)

// verifiedEmail 構造体
type verifiedEmail struct {
    address string
}

// UnverifiedEmail 構造体
type UnverifiedEmail struct {
    Address string
}

// Verify メソッド
func (u UnverifiedEmail) Verify() *verifiedEmail {
    return &verifiedEmail{address: u.Address}
}

// SendPasswordResetEmail メソッド
func (v verifiedEmail) SendPasswordResetEmail() {
    fmt.Printf("Sending password reset email to %s\n", v.address)
}

func main() {
	unverified := UnverifiedEmail{Address: "hoge@example.com"}
	verified := unverified.Verify()
	verified.SendPasswordResetEmail()
}
