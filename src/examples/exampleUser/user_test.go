package exampleUser

import (
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/sqldb"
)

func init() {
	//db
	err := sqldb.Init(sqldb.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "testdb",
		UserName: "root",
		Password: "123456",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//redis
	err = redisClient.Init(redisClient.Config{
		Address:   "127.0.0.1",
		UserName:  "",
		Password:  "",
		Port:      6379,
		KeyPrefix: "userTest:",
		UseTLS:    false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//cache
	err = cache.Init()
	if err != nil {
		log.Fatalln(err)
	}

}

func Test_UserDB(t *testing.T) {
	//add
	log.Println("---add user---")
	newUser := &ExampleUserModel{
		Status: "normal",
		Name:   "userName",
		Email:  "mail@email.com",
	}
	newUserInfo, err := InsertUser(newUser)
	if err != nil {
		log.Println("InsertUser error:", err)
		return
	}
	log.Println("newUserInfo:", newUserInfo)
	ID := newUserInfo.ID

	//get
	log.Println("---get user---")
	userInfo, err := GetUserById(ID, false)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	log.Println("userInfo:", userInfo)

	//update
	log.Println("---update user---")
	newData := map[string]interface{}{
		"status": "error",
		"name":   "userName2",
		"email":  "mail2@email.com",
	}
	err = UpdateUser(newData, ID)
	if err != nil {
		log.Println("UpdateUser error:", err)
		return
	}
	//get
	userInfo, err = GetUserById(ID, false)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	log.Println("userInfo:", userInfo)

	//delete
	log.Println("---delete user---")
	err = DeleteUser(ID)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return
	}
	userInfo, err = GetUserById(ID, false)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	log.Println("userInfo:", userInfo)
}
