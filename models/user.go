package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type UserModel struct {
	ID              int32    `json:"userId"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	ActivePoolIds   []int32  `json:"activePoolIds"`
	InactivePoolIds []int32  `json:"inactivePoolIds"`
	Interests       []string `json:"interests"`
}

func GetUserById(id int32, users []UserModel) (UserModel, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return UserModel{}, errors.New("user not found")
}

func GetUserByEmail(incomingEmail string, users []UserModel) bool {
	for _, user := range users {
		if strings.EqualFold(user.Email, incomingEmail) {
			return true
		}
	}
	return false
}

func GetUserJSONFile() []UserModel {
	// file, err := os.Open("users.json")
	// if err != nil {
	// 	return nil
	// }
	// var loadedUsers []UserModel
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&loadedUsers)
	// if err != nil {
	// 	fmt.Print("Error decoding JSON data\n")
	// 	return nil
	// }
	// return loadedUsers
	// usersFromJSONFile := DecodeData[UserModel](OpenFile("users.json"))
	// return usersFromJSONFile
	return DecodeData[UserModel](OpenFile("users.json"))
}
func AddUser(newUser UserModel) (bool, error) {
	fmt.Println("Adding new user...", newUser)

	newUser.ID = GenerateRandomNumber()
	users := GetUserJSONFile()
	if users != nil {
		if !GetUserByEmail(newUser.Email, users) {
			users = append(users, newUser)
		} else {
			return false, errors.New("user already exist.  Please try a different email")
		}

	} else {
		users = make([]UserModel, 0)
		users = append(users, newUser)
	}

	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println("err creating json: ", err)
		return false, errors.New("error creating json format")
	}

	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("err creating json file: ", err)
		return false, errors.New("error creating json file to os")
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("err writing to json file", err)
		return false, errors.New("error writing to the json file")
	}
	fmt.Println("successfully wrote to users.json")
	return true, nil
}
