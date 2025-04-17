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

	assert.Equal(t, "123", acc.ID, "ID аккаунта должен соответствовать")
	assert.Equal(t, "John Doe", acc.Owner, "Владелец аккаунта должен соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен соответствовать начальному значению")
	assert.True(t, acc.IsActive, "Новый аккаунт должен быть активным")
	assert.Equal(t, expectedTime, acc.CreatedAt, "Время создания аккаунта должно соответствовать")
	assert.Len(t, acc.GetTransactionHistory(), 1, "Новый аккаунт должен иметь одну начальную транзакцию")

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
		{"Положительный депозит", 100.0, 50.0, 150.0, nil},
		{"Нулевой депозит", 100.0, 0.0, 100.0, errors.New("deposit amount must be positive")},
		{"Отрицательный депозит", 100.0, -10.0, 100.0, errors.New("deposit amount must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.Deposit(tс.depositAmount, "тестовый депозит")

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error(), "Сообщение об ошибке должно соответствовать")
			} else {
				assert.NoError(t, err, "Не должно быть ошибки для валидного депозита")
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance, "Баланс аккаунта должен обновиться корректно")

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2, "История транзакций должна содержать новую запись")
				assert.Equal(t, "deposit", history[1].Type, "Тип транзакции должен быть 'deposit'")
				assert.Equal(t, tс.depositAmount, history[1].Amount, "Сумма транзакции должна соответствовать депозиту")
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
			}
		})
	}

	t.Run("Депозит в закрытый аккаунт", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false
		err := acc.Deposit(50.0, "тест")
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 100.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
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
		{"Успешное снятие", 100.0, 50.0, 50.0, nil},
		{"Недостаточно средств", 100.0, 150.0, 100.0, ErrInsufficientFunds},
		{"Нулевое снятие", 100.0, 0.0, 100.0, errors.New("withdrawal amount must be positive")},
		{"Отрицательное снятие", 100.0, -10.0, 100.0, errors.New("withdrawal amount must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.Withdraw(tс.withdrawAmount, "тестовое снятие")

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error(), "Сообщение об ошибке должно соответствовать")
			} else {
				assert.NoError(t, err, "Не должно быть ошибки для валидного снятия")
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance, "Баланс аккаунта должен обновиться корректно")

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2, "История транзакций должна содержать новую запись")
				assert.Equal(t, "withdrawal", history[1].Type, "Тип транзакции должен быть 'withdrawal'")
				assert.Equal(t, tс.withdrawAmount, history[1].Amount, "Сумма транзакции должна соответствовать снятию")
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
			}
		})
	}

	t.Run("Снятие с закрытого аккаунта", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false
		err := acc.Withdraw(50.0, "тест")
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 100.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
	})
}

func TestTransfer(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	t.Run("Успешный перевод", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

		err := acc1.Transfer(acc2, 30.0, "аренда")
		assert.NoError(t, err, "Не должно быть ошибки для валидного перевода")
		assert.Equal(t, 70.0, acc1.Balance, "Баланс исходного аккаунта должен уменьшиться")
		assert.Equal(t, 80.0, acc2.Balance, "Баланс целевого аккаунта должен увеличиться")

		history1 := acc1.GetTransactionHistory()
		history2 := acc2.GetTransactionHistory()
		assert.Len(t, history1, 2, "Исходный аккаунт должен иметь новую транзакцию")
		assert.Len(t, history2, 2, "Целевой аккаунт должен иметь новую транзакцию")
		assert.Equal(t, "transfer", history1[1].Type, "Тип транзакции исходного аккаунта должен быть 'transfer'")
		assert.Equal(t, -30.0, history1[1].Amount, "Сумма транзакции исходного аккаунта должна быть отрицательной")
		assert.Equal(t, "transfer", history2[1].Type, "Тип транзакции целевого аккаунта должен быть 'transfer'")
		assert.Equal(t, 30.0, history2[1].Amount, "Сумма транзакции целевого аккаунта должна быть положительной")
	})

	t.Run("Недостаточно средств", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

		err := acc1.Transfer(acc2, 150.0, "аренда")
		assert.EqualError(t, err, ErrInsufficientFunds.Error(), "Должна возвращаться ошибка недостатка средств")
		assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
		assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
		assert.Len(t, acc1.GetTransactionHistory(), 1, "История транзакций исходного аккаунта не должна изменяться")
		assert.Len(t, acc2.GetTransactionHistory(), 1, "История транзакций целевого аккаунта не должна изменяться")
	})

	t.Run("Перевод в закрытый аккаунт", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
		acc2.IsActive = false

		err := acc1.Transfer(acc2, 30.0, "аренда")
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
		assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
	})

	t.Run("Перевод из закрытого аккаунта", func(t *testing.T) {
		acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
		acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
		acc1.IsActive = false

		err := acc1.Transfer(acc2, 30.0, "аренда")
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
		assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
	})
}

func TestCloseAccount(t *testing.T) {
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)

	t.Run("Успешное закрытие", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		balance, err := acc.Close()

		assert.NoError(t, err, "Не должно быть ошибки для активного аккаунта")
		assert.Equal(t, 100.0, balance, "Должен возвращаться текущий баланс")
		assert.Equal(t, 0.0, acc.Balance, "Баланс аккаунта должен быть нулевым после закрытия")
		assert.False(t, acc.IsActive, "Аккаунт должен быть неактивным после закрытия")

		history := acc.GetTransactionHistory()
		assert.Len(t, history, 2, "История транзакций должна содержать запись о закрытии")
		assert.Equal(t, "closure", history[1].Type, "Тип транзакции должен быть 'closure'")
		assert.Equal(t, -100.0, history[1].Amount, "Сумма транзакции должна соответствовать балансу")
	})

	t.Run("Уже закрытый аккаунт", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 100.0, mockClock)
		acc.IsActive = false

		balance, err := acc.Close()
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 0.0, balance, "Должен возвращаться нулевой баланс для уже закрытого аккаунта")
		assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен оставаться неизменным")
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
		{"Положительная ставка", 1000.0, 5.0, 1050.0, nil},
		{"Нулевая ставка", 1000.0, 0.0, 1000.0, errors.New("interest rate must be positive")},
		{"Отрицательная ставка", 1000.0, -1.0, 1000.0, errors.New("interest rate must be positive")},
	}

	for _, tс := range tests {
		t.Run(tс.name, func(t *testing.T) {
			acc := NewAccount("123", "John Doe", tс.initialBalance, mockClock)
			err := acc.ApplyInterest(tс.rate)

			if tс.expectedError != nil {
				assert.EqualError(t, err, tс.expectedError.Error(), "Сообщение об ошибке должно соответствовать")
			} else {
				assert.NoError(t, err, "Не должно быть ошибки для валидной процентной ставки")
			}
			assert.Equal(t, tс.expectedBalance, acc.Balance, "Баланс аккаунта должен обновиться с учетом процентов")

			if tс.expectedError == nil {
				history := acc.GetTransactionHistory()
				assert.Len(t, history, 2, "История транзакций должна содержать запись о начислении процентов")
				assert.Equal(t, "interest", history[1].Type, "Тип транзакции должен быть 'interest'")
				assert.Equal(t, tс.expectedBalance-tс.initialBalance, history[1].Amount, "Сумма транзакции должна соответствовать начисленным процентам")
			} else {
				assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
			}
		})
	}

	t.Run("Начисление процентов на закрытый аккаунт", func(t *testing.T) {
		acc := NewAccount("123", "John Doe", 1000.0, mockClock)
		acc.IsActive = false
		err := acc.ApplyInterest(5.0)
		assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
		assert.Equal(t, 1000.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
	})
}
