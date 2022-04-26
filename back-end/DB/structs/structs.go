package structs

import "database/sql"

type User struct {
	Id string
	Name string
	Email string
	Password string
}
type Block struct {
	Id string
    Typed int
    Data string
    Top_id sql.NullString 
	User_id string
	Create_Date string
	Final_Date string
}