// Фасад (Facade) — это структурный паттерн, который предоставляет простой интерфейс, упрощающий
// использование сложной системы объектов, библиотек или фреймворков.

// Когда имеется сложная система, и необходимо упростить с ней работу, фасад позволит определить
// одну точку взаимодействия между клиентом и системой.
// Когда надо уменьшить количество зависимостей между клиентом и сложной системой,
// фасадные объекты позволяют отделить, изолировать компоненты системы от клиента и развивать
// и работать с ними независимо.
// Создание фасадов для компонентов каждой отдельной подсистемы позволит упростить взаимодействие
// между ними и повысить их независимость друг от друга.

// Минус паттерна фасад,структура фасада может разрастись до объекта-бога c большим количеством
// методов.

//  Если приложение взаимодействует с различными сторонними сервисами и библиотеками,
// фасад может упростить интеграцию, предоставляя чистый и удобный интерфейс
// для взаимодействия с ними.

package main

import (
	"fmt"
	"log"
)

// сложная структура спрятана за фасадом Wallet - он позволяет клиенту работать с десятками 
// компонентов, используя при этом простой интерфейс. Клиенту необходимо лишь ввести свои данные:
// имя аккаунта и код безопасности. Фасад управляет дальнейшей коммуникацией между различными 
// компонентами без контакта клиента со сложными внутренними механизмами.

type WalletFacade struct {
    account      *Account
    wallet       *Wallet
    securityCode *SecurityCode
    notification *Notification
    ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
    fmt.Println("Starting create account")
    walletFacacde := &WalletFacade{
        account:      newAccount(accountID),
        securityCode: newSecurityCode(code),
        wallet:       newWallet(),
        notification: &Notification{},
        ledger:       &Ledger{},
    }
    fmt.Println("Account created")
    return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting add money to wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }
    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    w.wallet.creditBalance(amount)
    w.notification.sendWalletCreditNotification()
    w.ledger.makeEntry(accountID, "credit", amount)
    return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting debit money from wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    err = w.wallet.debitBalance(amount)
    if err != nil {
        return err
    }
    w.notification.sendWalletDebitNotification()
    w.ledger.makeEntry(accountID, "debit", amount)
    return nil
}

type Account struct {
    name string
}

func newAccount(accountName string) *Account {
    return &Account{
        name: accountName,
    }
}

func (a *Account) checkAccount(accountName string) error {
    if a.name != accountName {
        return fmt.Errorf("Account Name is incorrect")
    }
    fmt.Println("Account Verified")
    return nil
}

type SecurityCode struct {
    code int
}

func newSecurityCode(code int) *SecurityCode {
    return &SecurityCode{
        code: code,
    }
}

func (s *SecurityCode) checkCode(incomingCode int) error {
    if s.code != incomingCode {
        return fmt.Errorf("Security Code is incorrect")
    }
    fmt.Println("SecurityCode Verified")
    return nil
}

type Wallet struct {
    balance int
}

func newWallet() *Wallet {
    return &Wallet{
        balance: 0,
    }
}

func (w *Wallet) creditBalance(amount int) {
    w.balance += amount
    fmt.Println("Wallet balance added successfully")
    return
}

func (w *Wallet) debitBalance(amount int) error {
    if w.balance < amount {
        return fmt.Errorf("Balance is not sufficient")
    }
    fmt.Println("Wallet balance is Sufficient")
    w.balance = w.balance - amount
    return nil
}

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
    fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
    return
}

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
    fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
    fmt.Println("Sending wallet debit notification")
}

func main() {
    fmt.Println()
    walletFacade := newWalletFacade("abc", 1234)
    fmt.Println()

    err := walletFacade.addMoneyToWallet("abc", 1234, 10)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }

    fmt.Println()
    err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }
}