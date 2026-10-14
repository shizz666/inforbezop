package main

import "fmt"

const (
	USD = 80
	EUR = 94
)

func main() {
	var rub float64

	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&rub)

	usd := rub / USD
	eur := rub / EUR
	usdToEur := usd * USD / EUR

	fmt.Println("Рубли в доллары:", usd, "USD")
	fmt.Println("Доллары в евро:", usdToEur, "EUR")
	fmt.Println("Рубли в евро:", eur, "EUR")

}
