package block

import (
	db "Todo_back_end/DB"
	"Todo_back_end/DB/structs"
	get "Todo_back_end/api/block/get"
	"Todo_back_end/utils"
	"fmt"
)

type Block struct{
	Get get.Get
}

func (Block) Create(block structs.Block) (int,*structs.Block){
	db := db.DB()
	block.Id = utils.RandStringBytes(10)
	if !block.Top_id.Valid {
		_, err := db.Exec("INSERT into block(user_id,id,type,data,final_date) values(?,?,?,?,?);",block.User_id,block.Id,block.Typed,block.Data,block.Final_Date)
		if err != nil {
			fmt.Println(err)
			return 0,nil
    	}
	}else {
		_, err := db.Exec("INSERT into block(user_id,id,type,data,top_id,final_date) values(?,?,?,?,?,?);",block.User_id,block.Id,block.Typed,block.Data,block.Top_id.String,block.Final_Date)
		if err != nil {
			fmt.Println(err)
			return 0,nil
    	}
	}
    defer db.Close()
	return 200,&block
}

func (Block) Edit(user_id string, block_id string,typed int,data string) int {
	db := db.DB()
	query := ""
    if typed == 1 { // type
        query = "UPDATE User set type = ? where user_id =? and id = ?"
    }else if typed == 2 { // data
        query = "UPDATE User set data = ? where user_id =? and id = ?"
    }else if typed == 3 { // top_id
        query = "UPDATE User set top_id = ? where user_id =? and id = ?"
    }else if typed == 4 { // final_date
        query = "UPDATE User set final_date = ? where user_id =? and id = ?"
    }
    _,err := db.Exec(query,data,user_id,block_id)
    if err != nil {
		fmt.Println(err)
		return 0
    }
	defer db.Close()	
	return 1	
}

func (Block) Delete(id string) int {
	db := db.DB()
	_, err := db.Exec("delete from block where id = ?",id)
	defer db.Close()
    if err != nil {
		fmt.Println(err)
		return 0
    }
	return 1	
}