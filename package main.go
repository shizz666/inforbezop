package main

import "fmt"

const (
	USD = 90.0  // 1 доллар = 90 рублей
	EUR = 100.0 // 1 евро = 100 рублей
)

func main() {
	var rub float64

	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&rub)

	usd := rub / USD
	eur := rub / EUR
	usdToEur := usd / EUR

	fmt.Println("Рубли в доллары:", usd, "USD")
	fmt.Println("Рубли в евро:", eur, "EUR")
	fmt.Println("Доллары в евро:", usdToEur, "EUR")
}
