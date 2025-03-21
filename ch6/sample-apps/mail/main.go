package main

import (
	"mail/send"
)

func main() {
	unverified := send.UnverifiedEmail{Address: "hoge@example.com"}
	verified := unverified.Verify()
	verified.SendPasswordResetEmail()
}
