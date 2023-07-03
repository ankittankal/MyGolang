package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	// Open a connection to the MySQL database using GORM
	d, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/book_store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = d
	fmt.Println("Connected to the MySQL database")
	d.LogMode(true) /*----The line d.LogMode(true) enables the log mode in GORM, which outputs the generated SQL statements and their respective parameters to the console.

	By default, GORM logs are disabled. However, by calling d.LogMode(true), where d is the *gorm.DB instance,
	you can enable the log mode and instruct GORM to print the generated SQL statements.    */

	// Use the 'db' variable of type *gorm.DB for performing database operations with GORM
}

func GetDB() *gorm.DB {
	return db
}

/* // Ping the database to verify the connection
err = db.Ping()
if err != nil {
	log.Fatal(err)
}
It is used to check the connectivity or availability of the database server.

In the context of GORM and database connections, db.Ping() is a method that sends a ping request to the database server and waits for a response.
If the database server responds successfully, it indicates that the connection to the database server is active and valid.
On the other hand, if there is an error or no response from the server, it suggests that the connection might be broken or the server is unreachable.

By calling db.Ping(), you can verify whether the database server is accessible before performing any database operations.
It's a good practice to check the connectivity to the database server to handle any potential connection errors or issues upfront.
*/
