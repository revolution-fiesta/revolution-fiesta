package main
/////

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "270153"
	dbname   = "mydb"
)

func main() {
	// 连接字符串
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// 打开数据库连接
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50),
        email VARCHAR(50),
		password VARCHAR(50),
		gender VARCHAR(5)
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully!")
	//追加表行信息
	alterTableSQL := `ALTER TABLE users ADD COLUMN WeChat_number VARCHAR(30);`
	_, err = db.Exec(alterTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// // 插入数据
	// insertSQL := `INSERT INTO users (name, age) VALUES ($1, $2)`
	// _, err = db.Exec(insertSQL, "Alice", 25)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// _, err = db.Exec(insertSQL, "Bob", 30)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println("Data inserted successfully!")

	// 查询数据
	// querySQL := `SELECT id, username, email,password,gender FROM users`
	// rows, err := db.Query(querySQL)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	//     var id int
	//     var name string
	//     var age int
	//     err = rows.Scan(&id, &name, &age)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	// }

	// // 更新数据
	// updateSQL := `UPDATE users SET age = $1 WHERE name = $2`
	// _, err = db.Exec(updateSQL, 26, "Alice")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println("Data updated successfully!")

	// // 删除数据
	// deleteSQL := `DELETE FROM users WHERE name = $1`
	// _, err = db.Exec(deleteSQL, "Bob")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println("Data deleted successfully!")
}
