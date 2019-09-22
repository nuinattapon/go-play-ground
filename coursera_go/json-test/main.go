package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var user = User{
	Id:    1,
	Name:  "Nattapon Sub-Anake",
	Email: "nattapon.subanake@gmail.com",
	Phone: "+66815517772",
}

func main() {

	barr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	fmt.Println(string(barr))

	str := `{"id":2,"name":"Wassana Sub-Anake","email":"nuchinter@hotmail.com","phone":"+6686122052"}`
	user2 := User{}
	err = json.Unmarshal([]byte(str), &user2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
	fmt.Println(user2)

}
