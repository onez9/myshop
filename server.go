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
	Id          int    `json:"id"`
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
var cart_products []Product

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

	var id int
	var price int
	var name string
	var description string
	var imgpath string

	// rows, err := db.Query()
	stmt, err := db.Prepare(`select id, name_product,description_product,price,imgpath from products`)
	if err != nil {
		panic(err)
	}

	defer stmt.Close() // закрытие соединения
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products = products[:0] // перед каждой загрузкой элементов в массив, отчищаем его полностью

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &price, &imgpath)
		if err != nil {
			panic(err)
		}

		products = append(products, Product{
			Id:          id,
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
func remove_i(slice []Product, index int) []Product {
	for i, product := range slice {
		if product.Id == index {
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
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// e.Use(middleware.Recover()) // промежуточные обработчики корневого уровня

	products = make([]Product, 0)
	cart_products = make([]Product, 0)

	// fillProductsToArray()
	// saveProductsToJSON()
	// loadProductsFromJSON()
	// sendDataToDBpostgres()
	loadFromDB() // Загрузка данных из БД
	// clear_image()
	// e := echo.New()
	//g := e.Group("/admin")     // промежуточные обработчики для группы
	//g.Use(middleware.Logger()) // промежуточные обработчики это функция,включенная в http-запрос-ответ с доступом к Echo.Context

	e.File("/favicon.ico", "favicon.ico")

	// Промежуточные обработчики для маршрута
	/*
			track := func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					println("request to /users")
					return next(c)
				}
			}


		e.GET("/users", func(c echo.Context) error {
			return c.String(http.StatusOK, "/users")
		}, track)
	*/

	// e.GET("/", func(c echo.Context) error {
	// 	sess, _ := session.Get("session", c)
	// 	sess.Options = &sessions.Options{
	// 		Path:     "/",
	// 		MaxAge:   86400 * 7,
	// 		HttpOnly: true,
	// 	}
	// 	sess.Values["foo"] = "bar"
	// 	sess.Save(c.Request(), c.Response())
	// 	return c.NoContent(http.StatusOK)
	// })

	e.GET("/whoami", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		email := sess.Values["email"]

		return c.JSON(http.StatusOK, email)
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
			//return c.JSON(http.StatusOK, products[0:0])
		}
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

	})

	// Загрузка при старте пользовательской корзины с продуктами картинками :)
	e.GET("/getcartitems", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		fmt.Println("session(addtocart) id: ", sess.Values["id"])
		fmt.Println("session(addtocart) email: ", sess.Values["email"])
		fmt.Println("session(addtocart) password: ", sess.Values["password"])

		query := `select p.id,p.name_product,p.description_product,p.price,p.imgpath from users as u
		left join users_products on u.id=users_products.user_id
		inner join products as p on users_products.product_id=p.id
		where u.id=$1`
		stmt, err := db.Prepare(query)
		if err != nil {
			panic(err)
		}

		defer stmt.Close() // закрытие соединения
		rows, err := stmt.Query(sess.Values["id"])
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var id int
		var price int
		var name string
		var description string
		var imgpath string

		cart_products = cart_products[:0] // перед каждым запросом к базе данных убираем все элементы
		for rows.Next() {
			err := rows.Scan(&id, &name, &description, &price, &imgpath)
			if err != nil {
				panic(err)
			}

			cart_products = append(cart_products, Product{
				Id:          id,
				Name:        name,
				Description: description,
				Price:       price,
				ImgPath:     imgpath,
			})
			fmt.Println(id, name, description, price, imgpath)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, cart_products)
		// return c.JSON(http.StatusOK, "OK")
	})

	// количество товаров
	e.GET("/getproductscount", func(c echo.Context) error {
		fmt.Println(len(products))
		return c.JSON(http.StatusOK, len(products))
	})
	// Регистрируем маршруты для статичных файлов
	e.Static("/public", "public")        // Пример будет обслуживать любой файл из каталога ресурсов по пути /public/*
	e.Static("/assets", "public/assets") // Пример будет обслуживать любой файл из каталога ресурсов по пути /public/assets/*
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
		// name := c.QueryParam("name")
		// fmt.Println(name)
		// type H map[string]interface{}
		return c.Redirect(http.StatusFound, "/home")
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
		loadFromDB()

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

	e.POST("/addtocart", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		fmt.Println("session(addtocart: ", sess.Values["id"])
		fmt.Println("session(addtocart: ", sess.Values["email"])
		fmt.Println("session(addtocart: ", sess.Values["password"])

		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		user_id := sess.Values["id"]
		product_id := int(json_map["product_id"].(float64))

		fmt.Println("user_id: ", user_id)
		fmt.Println("product_id: ", product_id)

		// тут мы добавляем в базу данных нужную инфу
		stmt, err := db.Prepare(`INSERT INTO users_products(user_id,product_id)Values($1,$2)`)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(user_id, product_id)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Добавлено %d\n", rowCnt)

		return c.JSON(http.StatusOK, "OK")
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
		sess.Values["email"] = ""
		fmt.Println(c)
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}
		email := json_map["email"].(string)
		password := json_map["password"].(string)
		hash := encrtyptPasswords(password)

		fmt.Println("email(authentication): ", email)
		fmt.Println("hash(authentication):: ", hash)

		stmt, err := db.Prepare(`select id, passw from users where email=$1`)
		if err != nil {
			log.Fatal(err)
		}
		// defer stmt.Close()

		var id int
		var hash_from_db string
		err = stmt.QueryRow(email).Scan(&id, &hash_from_db)
		fmt.Println(id)
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
			sess.Values["id"] = id
			sess.Values["email"] = email
			sess.Values["password"] = password
			fmt.Println(sess.Values)
			sess.Save(c.Request(), c.Response())

		} else {
			fmt.Println("Пошёл нахуй!")
		}
		return c.JSON(http.StatusOK, "OK")
	})

	e.POST("/authentication1", func(c echo.Context) error {
		// данные из формы авторизации
		email := c.FormValue("email")
		password := c.FormValue("password")
		hash := encrtyptPasswords(password)

		stmt, err := db.Prepare(`select id, passw from users where email=$1`)
		if err != nil {
			log.Fatal(err)
		}

		var id int
		var hash_from_db string // переменная для записи хешированного пароля из базы данных
		err = stmt.QueryRow(email).Scan(&id, &hash_from_db)
		fmt.Println(id)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Тут ничего нет!")
			} else {
				log.Fatal(err)
			}
		}

		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		sess.Values["id"] = ""
		sess.Values["email"] = ""
		sess.Values["password"] = ""
		sess.Save(c.Request(), c.Response())

		if hash_from_db == hash {
			fmt.Println("Авторизация успешно пройдена!")

			fmt.Println(sess)

			sess.Values["id"] = id
			sess.Values["email"] = email
			sess.Values["password"] = password
			fmt.Println(sess.Values)
			sess.Save(c.Request(), c.Response())

		} else {
			fmt.Println("Пошёл нахуй!")
		}
		return c.Redirect(http.StatusFound, "/home")
	})

	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
		// return 123
	})

	e.POST("/updaterec", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}
		id := json_map["id"].(float64)
		name := json_map["name"].(string)
		description := json_map["description"].(string)
		price := json_map["price"].(float64)
		// удаление записи из бд
		stmt, err := db.Prepare(`update products set name_product=$1, description_product=$2, price=$3 where id=$4`)
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(name, description, price, id) // тут мы записываем данные в БД
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Обновлено строк: ", rowCnt)

		// Обновляем массив для vue-шек
		for i, item := range products {

			if item.Id == int(id) {
				products[i].Name = name
				products[i].Description = description
				products[i].Price = int(price)

			}
		}

		return c.JSON(http.StatusOK, "OK")
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

	e.POST("/del_rec_in_cart", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		product_id, err := strconv.Atoi(c.QueryParam("id")) // product_id - получаем из url запроса
		user_id := sess.Values["id"]                        // user_id - получаем из сессии пользователя
		if err != nil {
			fmt.Println("Проблемы с конвертаций из строки в число маршруте - del_rec_in_cart")
			log.Fatal(err)
		}
		// удаление записи из бд корзины а именно связи иногие ко многим в промежуточной корзине
		stmt, err := db.Prepare(`delete from users_products where user_id=$1 and product_id=$2`)
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(user_id, product_id)
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Было затронуто: ", rowCnt)
		fmt.Println(777, len(products))
		cart_products = remove_i(cart_products, product_id)
		fmt.Println(777, len(cart_products))
		// подготовка отправки данных

		return c.JSON(http.StatusOK, cart_products)
	})

	e.POST("/logout", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["id"] = ""
		sess.Values["email"] = ""
		sess.Values["password"] = ""
		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, "OK")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
