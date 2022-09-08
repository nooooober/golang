package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string   `json:"name"`
	Nge    int      `json:"age"`
	Friend []string `json:"friend"`
}

func main() {
	usr := User{"Bob", 18, []string{"Tom", "Jim"}}
	b, _ := json.Marshal(usr)
	fmt.Println(string(b))
	var nux interface{}
	json.Unmarshal(b, &nux)
	fmt.Println(nux)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("fdjfasdjfjksda")))
	fmt.Println()
}
