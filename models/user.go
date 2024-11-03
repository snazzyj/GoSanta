package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserModel struct {
	ID              int32    `json:"userId"`
	Name            string   `json:"name"`
	ActivePoolIds   []int32  `json:"activePoolIds"`
	InactivePoolIds []int32  `json:"inactivePoolIds"`
	Interests       []string `json:"interests"`
}

func GetUserJSONFile() []UserModel {
	file, err := os.Open("users.json")
	if err != nil {
		return nil
	}
	var loadedUsers []UserModel
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&loadedUsers)
	if err != nil {
		fmt.Print("Error decoding JSON data\n")
		return nil
	}
	return loadedUsers
}
func AddUser(newUser UserModel) bool {
	fmt.Println("Adding new user...", newUser)

	newUser.ID = GenerateRandomNumber()
	users := GetUserJSONFile()
	if users != nil {
		users = append(users, newUser)
	} else {
		users = make([]UserModel, 0)
		users = append(users, newUser)
	}

	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println("err creating json: ", err)
		return false
	}

	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("err creating json file: ", err)
		return false
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("err writing to json file", err)
		return false
	}
	fmt.Println("successfully wrote to users.json")
	return true
}
