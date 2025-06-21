package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"html/template"
	"image/png"
	"log"
	"net/http"
)

// Page Структура Page для встраивания
type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", pageHandler)               //Обработка главной страницы
	http.HandleFunc("/generator/", viewCodeHandler) //Страница генерации картинки
	fmt.Printf("Запуск сервера на порту 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil)) //Запуск прослушивания сервера на входящие запросы
}

// Генерация страницы для ввода данных
func pageHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Генератор QR-кода"}         //Переменная
	t, _ := template.ParseFiles("generator.html") //Шаблон
	t.Execute(w, p)                               //Встраивание переменной в шаблон
}

// Генерация QR-кода в виде картинки
func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString") // Получение параметра и реквеста

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto) //Генерация QR-кода по реквесту
	qrCode, _ = barcode.Scale(qrCode, 512, 512)       //Генерация картинки
	png.Encode(w, qrCode)                             //Кодирование QR-кода в картинку и встраивание в ответ
}
