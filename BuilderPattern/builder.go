package BuilderPattern

import (
	"strconv"
)

type account struct {
	owner          string
	account_number string
	balance        uint
}

type Account interface {
	Deposit(uint) string
	Withdrawal(uint) string
	Payment(uint) string
	ShowInfo() string
}

type accountBuilder struct {
	owner          string
	account_number string
	balance        uint
}

type AccountBuilder interface {
	Owner(string) AccountBuilder
	AccountNumber(string) AccountBuilder
	Build() Account
}

func (ab *accountBuilder) Owner(owner string) AccountBuilder {
	ab.owner = owner
	return ab
}

func (ab *accountBuilder) AccountNumber(account_number string) AccountBuilder {
	ab.account_number = account_number
	return ab
}

func (ab *accountBuilder) Build() Account {
	return &account{
		owner:          ab.owner,
		account_number: ab.account_number,
		balance:        0,
	}
}

func New() AccountBuilder {
	return &accountBuilder{}
}

func (a *account) Deposit(amount uint) string {
	a.balance = a.balance + amount
	return a.owner + "'s balance is: " + strconv.FormatUint(uint64(a.balance), 10)
}

func (a *account) Withdrawal(amount uint) string {
	if a.balance < amount {
		return "Insufficient withdrawal balance"
	}
	a.balance = a.balance - amount
	return a.owner + "'s balance is: " + strconv.FormatUint(uint64(a.balance), 10)
}

func (a *account) Payment(amount uint) string {
	if a.balance < amount {
		return "Insufficient withdrawal balance"
	}
	a.balance = a.balance - amount
	return a.owner + "'s balance is: " + strconv.FormatUint(uint64(a.balance), 10)
}

func (a account) ShowInfo() string {
	return a.owner + "'s AccountNumber is " + a.account_number + " And Balance is " + strconv.FormatUint(uint64(a.balance), 10)
}
