package get

import (
	db "Todo_back_end/DB"
	"Todo_back_end/DB/structs"
	"database/sql"
	"log"
)

type Get struct {}

func (Get) User_id(user_id string) []structs.Block {
	db := db.DB()
	var Dan []structs.Block
	rows, err := db.Query("SELECT * FROM block where user_id = ?", user_id)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() 
    for rows.Next() {
		temp := structs.Block{}
        err := rows.Scan(&temp.User_id,&temp.Id,&temp.Typed,&temp.Data,&temp.Top_id,&temp.Create_Date,&temp.Final_Date)
        if err != nil {
            log.Fatal(err)
        }
		Dan = append(Dan, temp)
    }
    defer db.Close()
	return Dan
}
func (Get) Block_id(user_id string,block_id string) []structs.Block {
	db := db.DB()
	var Dan []structs.Block
	rows, err := db.Query("SELECT * FROM block where user_id = ? and id = ?", user_id,block_id)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() 
    for rows.Next() {
		temp := structs.Block{}
        err := rows.Scan(&temp.User_id,&temp.Id,&temp.Typed,&temp.Data,&temp.Top_id,&temp.Create_Date,&temp.Final_Date)
        if err != nil {
            log.Fatal(err)
        }
		Dan = append(Dan, temp)
    }
    defer db.Close()
	return Dan
}
func (Get) Custom(user_id string,where string,data string) []structs.Block {
	db := db.DB()
	var Dan []structs.Block
    var rows *sql.Rows
    var err error
    if data == "" {
        rows, err = db.Query("SELECT * FROM block where user_id = ? and "+where, user_id,sql.NullString{"",false})
    }else {
        rows, err = db.Query("SELECT * FROM block where user_id = ? and "+where, user_id,sql.NullString{data,true})
    }
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() 

    for rows.Next() {
		temp := structs.Block{}
        err := rows.Scan(&temp.User_id,&temp.Id,&temp.Typed,&temp.Data,&temp.Top_id,&temp.Create_Date,&temp.Final_Date)
        if err != nil {
            log.Fatal(err)
        }
		Dan = append(Dan, temp)
    }
    defer db.Close()
	return Dan
}