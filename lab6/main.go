package main

import (
	"fmt"
	"lab6/bank"
	"time"
)

func main() {
	clock := bank.NewClock()

	acc1 := bank.NewAccount("acc1", "Иван Иванов", 1000.0, clock)
	acc2 := bank.NewAccount("acc2", "Петр Петров", 500.0, clock)

	fmt.Println("=== Операции с аккаунтом 1 ===")
	if err := acc1.Deposit(200.0, "Зарплата"); err != nil {
		fmt.Println("Ошибка при пополнении:", err)
	}

	if err := acc1.Withdraw(150.0, "Покупка в магазине"); err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	if err := acc1.Transfer(acc2, 300.0, "Возврат долга"); err != nil {
		fmt.Println("Ошибка при переводе:", err)
	}

	if err := acc1.ApplyInterest(5.0); err != nil {
		fmt.Println("Ошибка при начислении процентов:", err)
	}

	printAccountInfo(acc1)
	printAccountInfo(acc2)

	// Закрываем второй аккаунт
	if balance, err := acc2.Close(); err != nil {
		fmt.Println("Ошибка при закрытии счета:", err)
	} else {
		fmt.Printf("Аккаунт %s закрыт, возвращено: %.2f\n", acc2.ID, balance)
	}
}

func printAccountInfo(acc *bank.Account) {
	fmt.Printf("\nСостояние аккаунта %s (%s):\n", acc.ID, acc.Owner)
	fmt.Printf("Баланс: %.2f\n", acc.Balance)
	fmt.Printf("Статус: %v\n", acc.IsActive)

	fmt.Println("\nИстория операций:")
	for _, t := range acc.GetTransactionHistory() {
		fmt.Printf("- [%s] %s: %.2f (%s) в %v\n",
			t.Type, t.Message, t.Amount, t.Timestamp.Format(time.RFC3339))
	}
}
