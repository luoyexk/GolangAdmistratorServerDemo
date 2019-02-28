package main

/**
建表
CREATE TABLE `userinfo` (
    `autid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `password` VARCHAR(64) NULL DEFAULT NULL,
    `uid` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);
*/
import (
	_ "GolangAdmistratorServerDemo/routers"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER = "root:zwLyykl8@tcp(localhost:3306)/testDB?charset=utf-8"
)

func main() {
	openDB()
	//beego.Run()
}

/**
打开数据库
*/
func openDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_DRIVER)
	if err != nil {
		isOpen = false
		fmt.Println("open failed")
	} else {
		isOpen = true
		fmt.Println("open success")
	}
	return isOpen, db
}

/**
增
*/
func add(db *sql.DB) {

}
