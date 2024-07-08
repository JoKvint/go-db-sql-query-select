package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	// Объявление пустого слайса sales типа Sale, который будет заполнен данными о продажах.
	var sales []Sale

	// напишите код здесь
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	// Выполнение SQL-запроса для выбора данных о продажах, где client равен переданному идентификатору клиента.
	// Используется именованная переменная :client для безопасного вставления значения.
	// Если возникает ошибка, функция возвращает nil и ошибку.
	// defer rows.Close() гарантирует закрытие результата запроса после использования.
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Цикл for перебирает все строки результата запроса. Для каждой строки создается новый экземпляр Sale.
	for rows.Next() {
		sale := Sale{}

		// Считывание данных из текущей строки результата запроса в поля структуры sale.
		// Если происходит ошибка, функция возвращает nil и ошибку.
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			return nil, err
		}

		// Добавление считанной продажи в слайс sales.
		sales = append(sales, sale)
	}

	// Возвращение слайса sales с данными о продажах и nil в случае отсутствия ошибок.
	return sales, nil
}

func main() {
	client := 208

	// Вызов функции selectSales с переданным идентификатором клиента. Проверка на наличие ошибки и вывод ее в случае возникновения.
	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Цикл for перебирает слайс sales и выводит информацию о каждой продаже.
	for _, sale := range sales {
		fmt.Println(sale)
	}
}
