package get

import (
	"Todo_back_end/DB/structs"
	Handling "Todo_back_end/api"
	"database/sql"
)

var all Handling.Handling

type Blocks struct {
	Id string
	Typed int
	Data string
	Top_id sql.NullString 
	User_id string
	Create_Date string
	Final_Date string
	Down []Blocks
}

func Change(data *structs.Block) *Blocks {
	return &Blocks{Id :data.Id,Typed: data.Typed,Top_id: data.Top_id,User_id: data.User_id,Data: data.Data, Create_Date: data.Create_Date,Final_Date: data.Final_Date}
}
func Create(user_id string, root *Blocks) {
	origin := all.Block.Get.Custom(user_id,"top_id = ?",root.Id)
	for i,data := range origin {
		root.Down = append(root.Down, *Change(&data))
		if data.Typed != 3 {
			Create(user_id,&root.Down[i])
		}
	}
}
