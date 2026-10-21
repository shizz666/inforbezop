package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const (
	USD = 90.0  // 1 доллар = 90 рублей
	EUR = 100.0 // 1 евро = 100 рублей
)

// Данные, которые передаём в index.html
type PageData struct {
	Input    string // что ввёл пользователь
	Show     bool   // показывать ли результат
	Error    string
	USD      int
	EUR      int
	UsdToEur int
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	input := r.URL.Query().Get("rub")
	if input != "" {
		data.Input = input

		rub, err := strconv.ParseFloat(input, 64)
		if err != nil {
			data.Error = "Введите число, например 1000"
		} else {
			usd := rub / USD
			eur := rub / EUR
			usdToEur :=  usd * USD / EUR*0.9

			data.Show = true
			data.USD = int(usd)
			data.EUR = int(eur)
			data.UsdToEur = int(usdToEur)
		}
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Не найден index.html: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)

	port := ":8080"
	fmt.Println("Сервер запущен на http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
