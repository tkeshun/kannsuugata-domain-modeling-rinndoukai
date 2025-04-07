package main

import (
    "fmt"
    "errors"
)

// MoneyTransfer は送金を表します
type MoneyTransfer struct {
    Id          string
    ToAccount   string
    FromAccount string
    Amount      float64
}

// Account は口座を表します
type Account struct {
    Name           string
    AccountId      string
    AccountBalance float64
}

// NewAccount は新しい口座を作成します
func NewAccount(name, accountId string, balance float64) *Account {
    return &Account{Name: name, AccountId: accountId, AccountBalance: balance}
}

// SendMoney は送金を行います
func SendMoney(fromAccount *Account, transfer MoneyTransfer) error {
    if fromAccount.AccountId != transfer.FromAccount {
        return errors.New("from account ID mismatch")
    }

    if fromAccount.AccountBalance < transfer.Amount {
        return errors.New("insufficient funds")
    }

    fromAccount.AccountBalance -= transfer.Amount
    fmt.Printf("Sent %.2f from %s to %s\n", transfer.Amount, transfer.FromAccount, transfer.ToAccount)
    return nil
}

// ReceiveMoney は受金を行います
func ReceiveMoney(toAccount *Account, transfer MoneyTransfer) error {
    if toAccount.AccountId != transfer.ToAccount {
        return errors.New("to account ID mismatch")
    }

    toAccount.AccountBalance += transfer.Amount
    fmt.Printf("Received %.2f to %s from %s\n", transfer.Amount, transfer.ToAccount, transfer.FromAccount)
    return nil
}

func main() {
    // 口座を作成
    accounts := map[string]*Account{
        "acc1": NewAccount("Alice", "acc1", 1000.0),
        "acc2": NewAccount("Bob", "acc2", 500.0),
    }

    // 送金を作成
    transfer := MoneyTransfer{
        Id:          "trans1",
        ToAccount:   "acc2",
        FromAccount: "acc1",
        Amount:      200.0,
    }

    // 送金を実行
    err := SendMoney(accounts["acc1"], transfer)
    if err != nil {
        fmt.Println("Send failed:", err)
    }

    // 任意のタイミングで受金を実行
    err = ReceiveMoney(accounts["acc2"], transfer)
    if err != nil {
        fmt.Println("Receive failed:", err)
    }

    // 口座の残高を表示
    for _, account := range accounts {
        fmt.Printf("Account %s (%s): Balance %.2f\n", account.Name, account.AccountId, account.AccountBalance)
    }
}
