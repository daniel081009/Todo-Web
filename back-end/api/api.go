package api

import (
	"Todo_back_end/api/block"
	"Todo_back_end/api/user"
)

type Handling struct {
	User user.User 
	Block block.Block
}