package main

import (
	"encoding/json"
	"fmt"
)

type Common struct {
	Id   int64
	Name string
}

type User struct {
	Common
	Age int64
}

func main() {
	user := new(User)
	user.Id = 1
	user.Name = "张三"
	user.Age = 10
	ub, _ := json.Marshal(user)

	userResult := new(User)
	json.Unmarshal(ub, userResult)
	fmt.Println(string(ub))
}
