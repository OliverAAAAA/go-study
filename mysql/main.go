package main

import (
	"database/sql"
	"fmt"
)

func main(){
	db,err:=sql.Open("mysql" ,"root:123456@tcp(127.0.0.1:3306)/oliver?charset=utf8")
	if err != nil{
		fmt.Printf("connect mysql fail ! [%s]",err)
	}else{
		fmt.Println("connect to mysql success")
	}
	defer func() {
		if db!=nil{
			db.Close()
		}
	}()

	stmt,_:=db.Prepare("insert into table_name column_1 values ('111') ")
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	r,_ := stmt.Exec()
	count,_:= r.RowsAffected()
	fmt.Printf(string(count));
}

