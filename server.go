package main

import (
	// "encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // загружается анонимно
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImgPath     string `json:"imgpath"`
}

// var products = []Product
var products []Product

func saveProductsToJSON() {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bytes, err := json.Marshal(products) // преобразуем массив в bytes
	file.Write(bytes)                    // записываем в файл
}

func loadProductsFromJSON() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		log.Fatal(err)
	}
}

func fillProductsToAqrray() {
	products = append(products, Product{
		Name:        "Товар1",
		Description: "Приносит удачу 1",
		Price:       1,
		ImgPath:     "assets/img/1.jpg",
	})
	products = append(products, Product{
		Name:        "Товар2",
		Description: "Приносит удачу 2",
		Price:       2,
		ImgPath:     "assets/img/2.jpg",
	})
	products = append(products, Product{
		Name:        "Товар3",
		Description: "Приносит удачу 3",
		Price:       3,
		ImgPath:     "assets/img/3.jpg",
	})
	products = append(products, Product{
		Name:        "Товар4",
		Description: "Приносит удачу 4",
		Price:       4,
		ImgPath:     "assets/img/4.jpg",
	})
	products = append(products, Product{
		Name:        "Товар5",
		Description: "Приносит удачу 5",
		Price:       5,
		ImgPath:     "assets/img/5.jpg",
	})
	products = append(products, Product{
		Name:        "Товар6",
		Description: "Приносит удачу 6",
		Price:       6,
		ImgPath:     "assets/img/6.jpg",
	})
	products = append(products, Product{
		Name:        "Товар7",
		Description: "Приносит удачу 7",
		Price:       7,
		ImgPath:     "assets/img/7.jpg",
	})
	products = append(products, Product{
		Name:        "Товар8",
		Description: "Приносит удачу 8",
		Price:       8,
		ImgPath:     "assets/img/8.jpg",
	})
	products = append(products, Product{
		Name:        "Товар9",
		Description: "Приносит удачу 9",
		Price:       9,
		ImgPath:     "assets/img/9.jpg",
	})
	products = append(products, Product{
		Name:        "Товар10",
		Description: "Приносит удачу 10",
		Price:       10,
		ImgPath:     "assets/img/10.jpg",
	})
	products = append(products, Product{
		Name:        "Эскимосы",
		Description: "Владеет магией снега",
		Price:       1000999,
		ImgPath:     "assets/img/11.jpg",
	})
}

func main() {
	e := echo.New()
	products = make([]Product, 0)
	// fillProductsToArray()
	// saveProductsToJSON()
	loadProductsFromJSON()
	connStr := `
	user=postgres
	password=123
	host=localhost
	dbname=postgres`

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()

	var price string
	var name string
	var customer_id string
	rows, err := db.Query(`select name,price,customer_id from products;`)
	if err != nil {
		panic(err)
	}

	defer rows.Close() // закрытие соединения

	for rows.Next() {
		err := rows.Scan(&name, &price, &customer_id)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, price, customer_id)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})

	// Маршрут для products
	e.GET("/getproducts", func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("p"))
		if err != nil {
			log.Fatal(err)
		}
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			log.Fatal(err)
		}

		fromIndex := page * limit     // начальный индекс товара
		toIndex := page*limit + limit // конечный индекс товара
		if toIndex > len(products) {
			toIndex = len(products)
		}
		productsPage := products[fromIndex:toIndex]

		fmt.Println(page, limit)
		return c.JSON(http.StatusOK, productsPage)
	})

	// количество товаров
	e.GET("/getproductscount", func(c echo.Context) error {
		fmt.Println(len(products))
		return c.JSON(http.StatusOK, len(products))
	})
	// Регистрируем маршруты для статичных файлов
	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	// Пример обработчика GET с получением параметров
	e.GET("/testget", func(c echo.Context) error {
		name := c.QueryParam("name")
		fmt.Println(name)
		type H map[string]interface{}
		return c.JSON(http.StatusOK, H{
			"a1": 123,
			"a2": 123,
			"a3": 123,
		})
	})

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		fmt.Println(name)
		type H map[string]interface{}
		return c.JSON(http.StatusOK, H{
			"t1": "1",
			"t2": "2",
			"t3": "3",
		})
	})

	// Пример обработчика запроса POST с получением параметров
	e.POST("/testpost", func(c echo.Context) error {
		fmt.Println(c)
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		fmt.Println(json_map)
		name := json_map["name"].(string)
		v := map[string]interface{}{
			"response": "Добрый день, " + name,
		}
		fmt.Println(v)
		return c.JSON(http.StatusOK, v)
	})

	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
		// return 123
	})

	e.Logger.Fatal(e.Start(":1323"))
}
