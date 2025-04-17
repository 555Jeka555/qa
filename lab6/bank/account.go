package bank

import (
	"errors"
	"time"
)

type TransactionType = string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
	Closure    TransactionType = "closure"
	Interest   TransactionType = "interest"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrAccountClosed     = errors.New("account is closed")
)

type Clock interface {
	GetNow() time.Time
}

func NewClock() Clock {
	return &clock{}
}

type clock struct {
}

func (c *clock) GetNow() time.Time {
	return time.Now()
}

type Transaction struct {
	Type      TransactionType
	Amount    float64
	Timestamp time.Time
	Message   string
}

func NewAccount(
	id,
	owner string,
	initialDeposit float64,
	clock Clock,
) *Account {
	return &Account{
		ID:        id,
		Owner:     owner,
		Balance:   initialDeposit,
		IsActive:  true,
		CreatedAt: clock.GetNow(),
		transactions: []Transaction{{
			Type:      Deposit,
			Amount:    initialDeposit,
			Timestamp: clock.GetNow(),
			Message:   "Initial deposit",
		}},
		clock: clock,
	}
}

type Account struct {
	ID           string
	Owner        string
	Balance      float64
	IsActive     bool
	CreatedAt    time.Time
	transactions []Transaction
	clock        Clock
}

func (a *Account) Deposit(amount float64, message string) error {
	if !a.IsActive {
		return ErrAccountClosed
	}
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	a.Balance += amount
	a.transactions = append(a.transactions, Transaction{
		Type:      Deposit,
		Amount:    amount,
		Timestamp: a.clock.GetNow(),
		Message:   message,
	})

	return nil
}

func (a *Account) Withdraw(amount float64, message string) error {
	if !a.IsActive {
		return ErrAccountClosed
	}
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	a.transactions = append(a.transactions, Transaction{
		Type:      Withdrawal,
		Amount:    amount,
		Timestamp: a.clock.GetNow(),
		Message:   message,
	})

	return nil
}

func (a *Account) Transfer(to *Account, amount float64, message string) error {
	if !a.IsActive || !to.IsActive {
		return ErrAccountClosed
	}
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}

	a.Balance -= amount
	to.Balance += amount

	now := a.clock.GetNow()
	a.transactions = append(a.transactions, Transaction{
		Type:      Transfer,
		Amount:    -amount,
		Timestamp: now,
		Message:   "Outgoing: " + message,
	})
	to.transactions = append(to.transactions, Transaction{
		Type:      Transfer,
		Amount:    amount,
		Timestamp: now,
		Message:   "Incoming: " + message,
	})

	return nil
}

func (a *Account) Close() (float64, error) {
	if !a.IsActive {
		return 0, ErrAccountClosed
	}

	remainingBalance := a.Balance
	a.Balance = 0
	a.IsActive = false
	a.transactions = append(a.transactions, Transaction{
		Type:      Closure,
		Amount:    -remainingBalance,
		Timestamp: a.clock.GetNow(),
		Message:   "Account closed",
	})

	return remainingBalance, nil
}

func (a *Account) GetTransactionHistory() []Transaction {
	return a.transactions
}

func (a *Account) ApplyInterest(rate float64) error {
	if !a.IsActive {
		return ErrAccountClosed
	}
	if rate <= 0 {
		return errors.New("interest rate must be positive")
	}

	interest := a.Balance * rate / 100
	a.Balance += interest
	a.transactions = append(a.transactions, Transaction{
		Type:      Interest,
		Amount:    interest,
		Timestamp: a.clock.GetNow(),
		Message:   "Interest applied",
	})

	return nil
}
