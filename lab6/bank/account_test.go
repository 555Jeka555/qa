package bank

import (
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
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	// Действие
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Проверка
	assert.Equal(t, "123", acc.ID, "ID аккаунта должен соответствовать")
	assert.Equal(t, "John Doe", acc.Owner, "Владелец аккаунта должен соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен соответствовать начальному значению")
	assert.True(t, acc.IsActive, "Новый аккаунт должен быть активным")
	assert.Equal(t, expectedTime, acc.CreatedAt, "Время создания аккаунта должно соответствовать")
	assert.Len(t, acc.GetTransactionHistory(), 1, "Новый аккаунт должен иметь одну начальную транзакцию")
	mockClock.AssertExpectations(t)
}

func TestDeposit_PositiveAmount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Deposit(50.0, "тестовый депозит")

	// Проверка
	assert.NoError(t, err, "Не должно быть ошибки для валидного депозита")
	assert.Equal(t, 150.0, acc.Balance, "Баланс аккаунта должен обновиться корректно")

	history := acc.GetTransactionHistory()
	assert.Len(t, history, 2, "История транзакций должна содержать новую запись")
	assert.Equal(t, "deposit", history[1].Type, "Тип транзакции должен быть 'deposit'")
	assert.Equal(t, 50.0, history[1].Amount, "Сумма транзакции должна соответствовать депозиту")
}

func TestDeposit_ZeroAmount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Deposit(0.0, "тестовый депозит")

	// Проверка
	assert.EqualError(t, err, "deposit amount must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта не должен измениться")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
}

func TestDeposit_NegativeAmount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Deposit(-10.0, "тестовый депозит")

	// Проверка
	assert.EqualError(t, err, "deposit amount must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта не должен измениться")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
}

func TestDeposit_ClosedAccount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)
	acc.IsActive = false

	// Действие
	err := acc.Deposit(50.0, "тест")

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 100.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
}

func TestWithdraw_Success(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Withdraw(50.0, "тестовое снятие")

	// Проверка
	assert.NoError(t, err, "Не должно быть ошибки для валидного снятия")
	assert.Equal(t, 50.0, acc.Balance, "Баланс аккаунта должен обновиться корректно")

	history := acc.GetTransactionHistory()
	assert.Len(t, history, 2, "История транзакций должна содержать новую запись")
	assert.Equal(t, "withdrawal", history[1].Type, "Тип транзакции должен быть 'withdrawal'")
	assert.Equal(t, 50.0, history[1].Amount, "Сумма транзакции должна соответствовать снятию")
}

func TestWithdraw_InsufficientFunds(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Withdraw(150.0, "тестовое снятие")

	// Проверка
	assert.EqualError(t, err, ErrInsufficientFunds.Error(), "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен остаться прежним")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться")
}

func TestWithdraw_ZeroAmount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Withdraw(0.0, "тестовое снятие")

	// Проверка
	assert.EqualError(t, err, "withdrawal amount must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен остаться прежним")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться")
}

func TestWithdraw_NegativeAmount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	err := acc.Withdraw(-10.0, "тестовое снятие")

	// Проверка
	assert.EqualError(t, err, "withdrawal amount must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен остаться прежним")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться")
}

func TestWithdraw_ClosedAccount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)
	acc.IsActive = false

	// Действие
	err := acc.Withdraw(50.0, "тест")

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 100.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
}

func TestTransfer_Success(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
	acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

	// Действие
	err := acc1.Transfer(acc2, 30.0, "аренда")

	// Проверка
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
}

func TestTransfer_InsufficientFunds(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
	acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)

	// Действие
	err := acc1.Transfer(acc2, 150.0, "аренда")

	// Проверка
	assert.EqualError(t, err, ErrInsufficientFunds.Error(), "Должна возвращаться ошибка недостатка средств")
	assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
	assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
	assert.Len(t, acc1.GetTransactionHistory(), 1, "История транзакций исходного аккаунта не должна изменяться")
	assert.Len(t, acc2.GetTransactionHistory(), 1, "История транзакций целевого аккаунта не должна изменяться")
}

func TestTransfer_ToClosedAccount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
	acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
	acc2.IsActive = false

	// Действие
	err := acc1.Transfer(acc2, 30.0, "аренда")

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
	assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
}

func TestTransfer_FromClosedAccount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc1 := NewAccount("123", "John Doe", 100.0, mockClock)
	acc2 := NewAccount("456", "Jane Doe", 50.0, mockClock)
	acc1.IsActive = false

	// Действие
	err := acc1.Transfer(acc2, 30.0, "аренда")

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 100.0, acc1.Balance, "Баланс исходного аккаунта не должен изменяться")
	assert.Equal(t, 50.0, acc2.Balance, "Баланс целевого аккаунта не должен изменяться")
}

func TestCloseAccount_Success(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)

	// Действие
	balance, err := acc.Close()

	// Проверка
	assert.NoError(t, err, "Не должно быть ошибки для активного аккаунта")
	assert.Equal(t, 100.0, balance, "Должен возвращаться текущий баланс")
	assert.Equal(t, 0.0, acc.Balance, "Баланс аккаунта должен быть нулевым после закрытия")
	assert.False(t, acc.IsActive, "Аккаунт должен быть неактивным после закрытия")

	history := acc.GetTransactionHistory()
	assert.Len(t, history, 2, "История транзакций должна содержать запись о закрытии")
	assert.Equal(t, "closure", history[1].Type, "Тип транзакции должен быть 'closure'")
	assert.Equal(t, -100.0, history[1].Amount, "Сумма транзакции должна соответствовать балансу")
}

func TestCloseAccount_AlreadyClosed(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 100.0, mockClock)
	acc.IsActive = false

	// Действие
	balance, err := acc.Close()

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 0.0, balance, "Должен возвращаться нулевой баланс для уже закрытого аккаунта")
	assert.Equal(t, 100.0, acc.Balance, "Баланс аккаунта должен оставаться неизменным")
}

func TestApplyInterest_PositiveRate(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 1000.0, mockClock)

	// Действие
	err := acc.ApplyInterest(5.0)

	// Проверка
	assert.NoError(t, err, "Не должно быть ошибки для валидной процентной ставки")
	assert.Equal(t, 1050.0, acc.Balance, "Баланс аккаунта должен обновиться с учетом процентов")

	history := acc.GetTransactionHistory()
	assert.Len(t, history, 2, "История транзакций должна содержать запись о начислении процентов")
	assert.Equal(t, "interest", history[1].Type, "Тип транзакции должен быть 'interest'")
	assert.Equal(t, 50.0, history[1].Amount, "Сумма транзакции должна соответствовать начисленным процентам")
}

func TestApplyInterest_ZeroRate(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 1000.0, mockClock)

	// Действие
	err := acc.ApplyInterest(0.0)

	// Проверка
	assert.EqualError(t, err, "interest rate must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 1000.0, acc.Balance, "Баланс аккаунта должен оставаться неизменным")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
}

func TestApplyInterest_NegativeRate(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 1000.0, mockClock)

	// Действие
	err := acc.ApplyInterest(-1.0)

	// Проверка
	assert.EqualError(t, err, "interest rate must be positive", "Сообщение об ошибке должно соответствовать")
	assert.Equal(t, 1000.0, acc.Balance, "Баланс аккаунта должен оставаться неизменным")
	assert.Len(t, acc.GetTransactionHistory(), 1, "История транзакций не должна изменяться при ошибке")
}

func TestApplyInterest_ClosedAccount(t *testing.T) {
	// Подготовка
	mockClock := new(MockClock)
	expectedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock.On("GetNow").Return(expectedTime)
	acc := NewAccount("123", "John Doe", 1000.0, mockClock)
	acc.IsActive = false

	// Действие
	err := acc.ApplyInterest(5.0)

	// Проверка
	assert.EqualError(t, err, ErrAccountClosed.Error(), "Должна возвращаться ошибка закрытого аккаунта")
	assert.Equal(t, 1000.0, acc.Balance, "Баланс не должен изменяться для закрытого аккаунта")
}
