package bank

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClock struct {
	mock.Mock
}

func (m *MockClock) GetNow() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func TestNewAccount(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	assert.Equal(t, "123", acc.ID) // todo добавить меседж
	assert.Equal(t, "John Doe", acc.Owner)
	assert.Equal(t, 100.0, acc.Balance)
	assert.True(t, acc.IsActive)
	assert.Equal(t, expectedTime, acc.CreatedAt)
	assert.Len(t, acc.GetTransactionHistory(), 1)

	mockClock.AssertExpectations(t)
}

func TestDeposit(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	tests := []struct {
		name            string
		initialBalance  float64
		depositAmount   float64
		expectedBalance float64
		expectedError   error
	}{
		{"Positive deposit", 100.0, 50.0, 150.0, nil},
		{"Zero deposit", 100.0, 0.0, 100.0, errors.New("deposit amount must be positive")},
		{"Negative deposit", 100.0, -10.0, 100.0, errors.New("deposit amount must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.Deposit(tс.depositAmount, "test deposit")

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance)

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2)
				assert.Equal(t, "deposit", history[1].Type)
				assert.Equal(t, tс.depositAmount, history[1].Amount)
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1)
			}
		})
	}

	t.Run("Deposit to closed account", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false
		err := acc.Deposit(50.0, "test")
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 100.0, acc.Balance)
	})
}

func TestWithdraw(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	tests := []struct {
		name            string
		initialBalance  float64
		withdrawAmount  float64
		expectedBalance float64
		expectedError   error
	}{
		{"Successful withdrawal", 100.0, 50.0, 50.0, nil},
		{"Insufficient funds", 100.0, 150.0, 100.0, ErrInsufficientFunds},
		{"Zero withdrawal", 100.0, 0.0, 100.0, errors.New("withdrawal amount must be positive")},
		{"Negative withdrawal", 100.0, -10.0, 100.0, errors.New("withdrawal amount must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.Withdraw(tс.withdrawAmount, "test withdrawal")

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance)

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2)
				assert.Equal(t, "withdrawal", history[1].Type)
				assert.Equal(t, tс.withdrawAmount, history[1].Amount)
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1)
			}
		})
	}

	t.Run("Withdrawal from closed account", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false
		err := acc.Withdraw(50.0, "test")
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 100.0, acc.Balance)
	})
}

func TestTransfer(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	t.Run("Successful transfer", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

		err := acc1.Transfer(acc2, 30.0, "rent")
		assert.NoError(t, err)
		assert.Equal(t, 70.0, acc1.Balance)
		assert.Equal(t, 80.0, acc2.Balance)

		history1 := acc1.GetTransactionHistory()
		history2 := acc2.GetTransactionHistory()
		assert.Len(t, history1, 2)
		assert.Len(t, history2, 2)
		assert.Equal(t, "transfer", history1[1].Type)
		assert.Equal(t, -30.0, history1[1].Amount)
		assert.Equal(t, "transfer", history2[1].Type)
		assert.Equal(t, 30.0, history2[1].Amount)
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

		err := acc1.Transfer(acc2, 150.0, "rent")
		assert.EqualError(t, err, ErrInsufficientFunds.Error())
		assert.Equal(t, 100.0, acc1.Balance)
		assert.Equal(t, 50.0, acc2.Balance)
		assert.Len(t, acc1.GetTransactionHistory(), 1)
		assert.Len(t, acc2.GetTransactionHistory(), 1)
	})

	t.Run("Transfer to closed account", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
		acc2.IsActive = false

		err := acc1.Transfer(acc2, 30.0, "rent")
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 100.0, acc1.Balance)
		assert.Equal(t, 50.0, acc2.Balance)
	})

	t.Run("Transfer from closed account", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
		acc1.IsActive = false

		err := acc1.Transfer(acc2, 30.0, "rent")
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 100.0, acc1.Balance)
		assert.Equal(t, 50.0, acc2.Balance)
	})
}

func TestCloseAccount(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	t.Run("Successful closure", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		balance, err := acc.Close()

		assert.NoError(t, err)
		assert.Equal(t, 100.0, balance)
		assert.Equal(t, 0.0, acc.Balance)
		assert.False(t, acc.IsActive)

		history := acc.GetTransactionHistory()
		assert.Len(t, history, 2)
		assert.Equal(t, "closure", history[1].Type)
		assert.Equal(t, -100.0, history[1].Amount)
	})

	t.Run("Already closed", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false

		balance, err := acc.Close()
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 0.0, balance)
		assert.Equal(t, 100.0, acc.Balance)
	})
}

func TestApplyInterest(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	tests := []struct {
		name            string
		initialBalance  float64
		rate            float64
		expectedBalance float64
		expectedError   error
	}{
		{"Positive rate", 1000.0, 5.0, 1050.0, nil},
		{"Zero rate", 1000.0, 0.0, 1000.0, errors.New("interest rate must be positive")},
		{"Negative rate", 1000.0, -1.0, 1000.0, errors.New("interest rate must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.ApplyInterest(tс.rate)

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance)

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2)
				assert.Equal(t, "interest", history[1].Type)
				assert.Equal(t, tс.expectedBalance-tс.initialBalance, history[1].Amount)
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1)
			}
		})
	}

	t.Run("Interest on closed account", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 1000.0, mockClock)
		acc.IsActive = false
		err := acc.ApplyInterest(5.0)
		assert.EqualError(t, err, ErrAccountClosed.Error())
		assert.Equal(t, 1000.0, acc.Balance)
	})
}
