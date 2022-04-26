package main

import (
	"Todo_back_end/DB/structs"
	Handling "Todo_back_end/api"
	"fmt"
)
func main() {
	all := Handling.Handling{}
	if all.User.Create(structs.User{Name: "Dan" ,Email: "daniel081009a@gmail.com",Password: "skymkey08"}) != 1 {
		fmt.Println("1 err")
	}
	user := all.User.Login("daniel081009a@gmail.com","skymkey08")
	if user == nil {
		fmt.Println("2 err")
	}
	if all.User.Edit(2,"daniel081009@naver.com",user.Id) != 1 {
		fmt.Println("3 err")
	}

	_,block_1 := all.Block.Create(structs.Block{Id: "Non",Typed: 0,Data: "hi",User_id:user.Id,Final_Date: "2022-10-09"})
	if block_1 == nil {
		fmt.Println("4 err")
	}	
	
	if all.Block.Get.User_id(user.Id) == nil {
		fmt.Println("5 err")
	}
	if all.Block.Get.Custom(block_1.Id,"final_date > ?","2021-10-09") == nil {
		fmt.Println("6 err")
	}
	block_1.Data = "Dan"
	if all.Block.Edit(user.Id,block_1.Id,1,"1") != 1 {
		fmt.Println("7 err")
	}

	if all.Block.Delete(block_1.Id) != 1 {
		fmt.Println("8 err")
	}
	if all.User.Delete(user.Id) != 1 {
		fmt.Println("9 err")
	}
	fmt.Println("Ok")
}