package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Demo() {
	Select()
	Insert()
	Update()
	Transaction()
}

func Select() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	type User struct {
		Id   int64
		Name string
		Age  int64
		Sex  int64
	}
	users := make([]User, 0, 10)

	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, *user)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println(users)
}

func Insert() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ret, err := db.Exec("insert users(user_name, age, sex) value ('mike', 18, 1)")
	if err != nil {
		panic(err.Error())
	}

	lastInsertId, err := ret.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertId)
}

func Update() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ret, err := db.Exec("update users set age = age + 1 where id = 2")
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffected)
}

func Transaction() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback()

	ret, err := tx.Exec("update users set age = age + 1 where id = 2")
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(ret.RowsAffected())
}