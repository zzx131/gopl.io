package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

	// using an encoder:
	file, _ := os.OpenFile("user.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(user)
	if err != nil {
		log.Println("Error in encoding json")
	}
	/// using an decoder
	fileDe, _ := os.Open("user.json")
	dec := json.NewDecoder(fileDe)
	dec.Decode(user)
	fmt.Println(user.Name)
}
