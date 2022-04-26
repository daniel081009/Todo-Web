package user

import (
	db "Todo_back_end/DB"
	"Todo_back_end/DB/structs"
	"Todo_back_end/utils"
)

type User struct {}


func (User)Create(user structs.User) int{
	db := db.DB()
	user.Id = utils.RandStringBytes(10) 
	_, err := db.Exec("INSERT into User(id,name,email,password) values(?,?,?,?)",user.Id,user.Name,user.Email,user.Password)
    if err != nil {
		return 0
    }
    defer db.Close()
	return 1
}
func (User)Login(email string,password string) *structs.User {
	db := db.DB()
    user := structs.User{}
    err := db.QueryRow("select * from User where email = ?",email).Scan(&user.Id,&user.Name,&user.Email,&user.Password)
    if err != nil {
        return nil
    }else if user.Password != password {
        return nil
    }
    defer db.Close()
    return &user
}
func (User)Edit(typed int,data string,id string) int {
	db := db.DB()
    query := ""
    if typed == 1 { // name
        query = "UPDATE User set name = ? where id = ?"
    }else if typed == 2 { // email
        query = "UPDATE User set email = ? where id = ?"
    }else if typed == 3 { // passwrod
        query = "UPDATE User set password = ? where id = ?"
    }
    _,err := db.Exec(query,data,id)
    defer db.Close()
    if err != nil {
        return 0
    }
    return 1 
}
func (User)Delete(id string) int {
	db := db.DB()
    data,err := db.Exec("DELETE from User where id=?",id)
    n,_ := data.RowsAffected()
    defer db.Close()
    if err != nil || n == 0{
        return 0
    }
    return 1
}