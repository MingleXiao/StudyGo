package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

func main()  {
	db, err :=sql.Open("mysql","root@root@/test?charset=utf8")

	checkErr(err)

	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)

	checkErr(err)
	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("插入成功")
	fmt.Println(id)
}

func checkErr(err error){
	if err!=nil{
		panic(err)
	}

}
