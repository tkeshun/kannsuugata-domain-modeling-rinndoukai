package main

import "fmt"

// Contactの定義
type Contact struct {
	Name        string
	ContactInfo ContactInfo
}

// ContactInfoインターフェース
type ContactInfo interface {
	contactInfo()
}

// パターン1: Email + Postal
type EmailAndPostal struct {
	Email  string
	Postal string
}

func (e EmailAndPostal) contactInfo() {}

// パターン2: Emailのみ
type EmailOnly struct {
	Email string
}

func (e EmailOnly) contactInfo() {}

// パターン3: Postalのみ
type PostalOnly struct {
	Postal string
}

func (p PostalOnly) contactInfo() {}

func main() {
	// OK: パターン1
	c1 := Contact{
		Name: "Alice",
		ContactInfo: EmailAndPostal{
			Email:  "alice@example.com",
			Postal: "12345",
		},
	}

	// OK: パターン2
	c2 := Contact{
		Name: "Bob",
		ContactInfo: EmailOnly{
			Email: "bob@example.com",
		},
	}

	// OK: パターン3
	c3 := Contact{
		Name: "Charlie",
		ContactInfo: PostalOnly{
			Postal: "67890",
		},
	}

	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%+v\n", c3)
}
