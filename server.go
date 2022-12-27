package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

// RedirectConfig struct {
//   // Skipper defines a function to skip middleware.
//   Skipper Skipper

//   // Status code to be used when redirecting the request.
//   // Optional. Default value http.StatusMovedPermanently.
//   Code int `json:"code"`
// }

// var DefaultRedirectConfig = RedirectConfig{
//   Skipper: DefaultSkipper,
//   Code:    http.StatusMovedPermanently,
// }

var connStr string = `
user=postgres
password=123
host=localhost
dbname=postgres`

var db, err = sql.Open("postgres", connStr)

// var products = []Product
var products []Product

func sendDataToDBpostgres() {

	stmt, err := db.Prepare(`INSERT INTO products(name_product,description_product,price,imgpath)Values($1,$2,$3,$4)`)
	if err != nil {
		log.Fatal(err)
	}
	for i, item := range products {
		fmt.Println(item, i)

		res, err := stmt.Exec(item.Name, item.Description, item.Price, item.ImgPath)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Добавлено %d\n", rowCnt)
	}
}
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
func encrtyptPasswords(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	// fmt.Println(b)
	str := base64.StdEncoding.EncodeToString(b)
	return str
}
func loadFromDB() {

	var price int
	var name string
	var description string
	var imgpath string

	// rows, err := db.Query()
	stmt, err := db.Prepare(`select name_product,description_product,price,imgpath from products`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close() // закрытие соединения
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &description, &price, &imgpath)
		if err != nil {
			panic(err)
		}

		products = append(products, Product{
			Name:        name,
			Description: description,
			Price:       price,
			ImgPath:     imgpath,
		})
		// fmt.Println(name, description, price, imgpath)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
func clear_image() {
	path1 := "/home/zakhar/sandbox/lab3/public/assets/img"
	path2 := "/home/zakhar/sandbox/lab3/"
	os.Chdir(path1)
	files, err := os.ReadDir(path1)
	if err != nil {
		fmt.Println(err)
	}
	for index, file := range files {
		if len([]rune(strings.Split(file.Name(), ".")[0])) == 36 {
			os.Remove(file.Name())
			fmt.Println(index, file.Name())
		}
	}

	os.Chdir(path2)
}
func remove(slice []Product, index string) []Product {
	for i, product := range slice {
		if product.Name == index {
			// fmt.Println(slice[i:])
			// fmt.Println(slice[i+1:])
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
func main() {

	// id := uuid.New()
	// fmt.Println(id.String())
	// fmt.Println(anime1.first + anime1.second)
	defer db.Close()

	e := echo.New()
	products = make([]Product, 0)
	// fillProductsToArray()
	// saveProductsToJSON()
	// loadProductsFromJSON()
	// sendDataToDBpostgres()
	loadFromDB()
	// clear_image()
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
		// Отвечает за сессию
		stmt, err := db.Prepare(`select passw from users where email=$1`)
		if err != nil {
			log.Fatal(err)
		}
		// defer stmt.Close()
		// тоже за сессию
		var hash_from_db string
		sess, _ := session.Get("session", c)
		email, err1 := sess.Values["email"]

		fmt.Println(email, err1)
		// если пользователя нет в базе данных то делай это
		err = stmt.QueryRow(email).Scan(&hash_from_db)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Тут ничего нет! В sql-запросе ничего найденно не было!")
			} else {
				log.Fatal(err)
			}
			return c.JSON(http.StatusOK, products[0:0])
		} else {
			// иначе делай это
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
		}
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
	e.POST("/addUser", func(c echo.Context) error {
		firstname := c.FormValue("name")
		lastname := c.FormValue("description")
		email := c.FormValue("email")
		password1 := c.FormValue("password1")
		password2 := c.FormValue("password2")
		fmt.Println(firstname, lastname, email, password1, password2)
		if password1 != password2 {
			return c.Redirect(http.StatusFound, "/reg")
		}
		password := encrtyptPasswords(password1)
		if err != nil {
			fmt.Println("Произошла ошибка с ковертацией price!")
			return err
		}

		err = db.Ping()
		stmt, err := db.Prepare(`INSERT INTO users(firstname,lastname,email,passw)Values($1,$2,$3,$4)`)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Длина хэша: ", len(password))
		res, err := stmt.Exec(firstname, lastname, email, password)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Добавлено %d\n", rowCnt)
		/*
			json_map := make(map[string]interface{})
			err := json.NewDecoder(c.Request().Body).Decode(&json_map)
			if err != nil {
				return err
			}

			fmt.Println(json_map)

			if err != nil {
				panic(err)
			}





			firstname := json_map["firstname"].(string)
			lastname := json_map["lastname"].(string)
			email := json_map["email"].(string)
			password := json_map["password"].(string)



		*/

		return c.Redirect(http.StatusFound, "/auth")
	})
	e.POST("/sendProduct", func(c echo.Context) error {

		// читаем данные из формы
		fmt.Println(c)
		name := c.FormValue("name")
		description := c.FormValue("description")
		price, err := strconv.Atoi(c.FormValue("price"))
		fmt.Println(name, description, price)
		if err != nil {
			fmt.Println("Произошла ошибка с ковертацией price!")
			return err
		}

		var fdefault bool = true
		var filename string
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("файл не загружен, отсутсвует, другие ошибки с файлом")
			//return err
			fdefault = false

		}

		if fdefault {
			src, err := file.Open()
			if err != nil {
				fmt.Println("тут ошибка 2")
				return err
			}
			defer src.Close()
			// imgpath := fmt.Sprintf("public/assets/img/%v", file.Filename)
			// fmt.Printf(imgpath)

			chunks := strings.Split(file.Filename, ".")
			ext := chunks[len(chunks)-1]
			filename = fmt.Sprintf("%v.%v", uuid.New().String(), ext)
			dst, err := os.Create(fmt.Sprintf("public/assets/img/%v", filename))
			if err != nil {
				fmt.Println("тут ошибка 3")
				return err

			}
			defer dst.Close()
			//Copy
			if _, err = io.Copy(dst, src); err != nil {
				fmt.Println("тут ошибка 4")
				return err

			}
		} else {
			filename = "cart.png"
		}
		// тут мы добавляем в базу данных нужную инфу
		stmt, err := db.Prepare(`INSERT INTO products(name_product,description_product,price,imgpath)Values($1,$2,$3,$4)`)
		if err != nil {
			log.Fatal(err)
		}

		imgpath := fmt.Sprintf("assets/img/%v", filename)
		res, err := stmt.Exec(name, description, price, imgpath)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Добавлено %d\n", rowCnt)

		// а это для vue шки
		fmt.Printf("name: %v\ndescription: %v\nprice: %v\n", name, description, price)
		products = append(products, Product{
			Name:        name,
			Description: description,
			Price:       price,
			ImgPath:     imgpath,
		})

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
		/*
			json_map := make(map[string]interface{})
			err := json.NewDecoder(c.Request().Body).Decode(&json_map)
			if err != nil {
				return err
			}

			fmt.Println(json_map)
			name := json_map["name"].(string)
			description := json_map["description"].(string)
			price := json_map["price"].(float64)

		*/
		// return c.JSON(http.StatusOK, "OK")

		// return c.Redirect(http.StatusOK, "/sendProduct")
		//return c.Redirect(http.StatusFound, "/products")
		// return c.JSON(http.StatusOK, products)
	})

	// для post с multipart/form-data работает норм - перезагрузка с использованием button с submit
	// e.POST("/sendProduct", func(c echo.Context) error {
	// 	// читаем данные из формы
	// 	fmt.Println(c)
	// 	name := c.FormValue("name")
	// 	description := c.FormValue("description")
	// 	price, err := strconv.Atoi(c.FormValue("price"))
	// 	fmt.Println(name, description, price)
	// 	if err != nil {
	// 		fmt.Println("Произошла ошибка с ковертацией price!")
	// 		return err
	// 	}
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		fmt.Println("тут ошибка 1")
	// 		return err
	// 	}
	// 	src, err := file.Open()
	// 	if err != nil {
	// 		fmt.Println("тут ошибка 2")
	// 		return err
	// 	}
	// 	defer src.Close()
	// 	// imgpath := fmt.Sprintf("public/assets/img/%v", file.Filename)
	// 	// fmt.Printf(imgpath)
	// 	dst, err := os.Create(fmt.Sprintf("public/assets/img/%v", file.Filename))
	// 	if err != nil {
	// 		fmt.Println("тут ошибка 3")
	// 		return err
	// 	}
	// 	defer dst.Close()
	// 	//Copy
	// 	if _, err = io.Copy(dst, src); err != nil {
	// 		fmt.Println("тут ошибка 4")
	// 		return err
	// 	}
	// 	// тут мы добавляем в базу данных нужную инфу
	// 	stmt, err := db.Prepare(`INSERT INTO products(name_product,description_product,price,imgpath)Values($1,$2,$3,$4)`)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	imgpath := fmt.Sprintf("assets/img/%v", file.Filename)
	// 	res, err := stmt.Exec(name, description, price, imgpath)
	// 	rowCnt, err := res.RowsAffected()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("Добавлено %d\n", rowCnt)
	// 	// а это для vue шки
	// 	fmt.Printf("name: %v\ndescription: %v\nprice: %v\n", name, description, price)
	// 	products = append(products, Product{
	// 		Name:        name,
	// 		Description: description,
	// 		Price:       price,
	// 		ImgPath:     imgpath,
	// 	})
	// 	/*
	// 		json_map := make(map[string]interface{})
	// 		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(json_map)
	// 		name := json_map["name"].(string)
	// 		description := json_map["description"].(string)
	// 		price := json_map["price"].(float64)
	// 	*/
	// 	// return c.JSON(http.StatusOK, "OK")
	// 	// return c.Redirect(http.StatusOK, "/sendProduct")
	// 	return c.Redirect(http.StatusFound, "/products")
	// })

	e.POST("/t2", func(c echo.Context) error {
		fmt.Println(c)
		// json_map := make(map[string]interface{})
		// err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		// if err != nil {
		// return err
		// }

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, "OK")
	})
	// Пример обработчика запроса POST с получением параметров
	e.POST("/authentication", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["email"] = "no"
		fmt.Println(c)
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}
		email := json_map["email"].(string)
		password := json_map["password"].(string)
		hash := encrtyptPasswords(password)
		fmt.Println(email)
		fmt.Println(hash)

		stmt, err := db.Prepare(`select passw from users where email=$1`)
		if err != nil {
			log.Fatal(err)
		}
		// defer stmt.Close()

		var hash_from_db string
		err = stmt.QueryRow(email).Scan(&hash_from_db)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Тут ничего нет!")
			} else {
				log.Fatal(err)
			}
		}
		/*
			row1 := stmt.QueryRow(email)

			err = row1.Scan(&hash_from_db)
			if err = row1.Err(); err != nil {
				log.Fatal(err)
			}
		*/
		fmt.Println(hash_from_db)

		if hash_from_db == hash {
			fmt.Println("Авторизация успешно пройдена!")

			fmt.Println(sess)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			sess.Values["email"] = email
			sess.Values["password"] = password
			fmt.Println(sess.Values)
			sess.Save(c.Request(), c.Response())

		} else {
			fmt.Println("Пошёл нахуй!")
		}
		return c.JSON(http.StatusOK, "OK")
	})
	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
		// return 123
	})
	e.POST("/delrec", func(c echo.Context) error {
		name := c.QueryParam("name")

		// удаление записи из бд
		stmt, err := db.Prepare(`delete from products where name_product=$1`)
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(name)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Было затронуто: ", rowCnt)
		fmt.Println(777, len(products))
		products = remove(products, name)
		fmt.Println(777, len(products))
		// подготовка отправки данных
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

		return c.JSON(http.StatusOK, productsPage)
	})
	e.Logger.Fatal(e.Start(":1323"))
}