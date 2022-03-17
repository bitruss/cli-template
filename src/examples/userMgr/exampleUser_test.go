package userMgr

import (
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
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
		log.Fatalln("db init err", err)
	}

	// auto migrate table in db
	// please create table by yourself in real project
	sqldb.GetInstance().AutoMigrate(&ExampleUserModel{})

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
		log.Fatalln("redis init err", err)
	}

	//reference
	err = reference.Init()
	if err != nil {
		log.Fatalln("reference init err", err)
	}
}

func Test_UserDB(t *testing.T) {
	//important! Please create db table before you run this test.
	//add
	log.Println("---add user---")
	newUser := &ExampleUserModel{
		Status: "normal",
		Name:   "userName",
		Email:  "mail@email.com",
	}
	newUserInfo, err := CreateUser(newUser)
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

	userInfo, err = GetUserById(ID, false)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	log.Println("userInfo:", userInfo)
}

func Test_UserArray(t *testing.T) {
	// user array
	//for i := 0; i < 10; i++ {
	//	newUser := &ExampleUserModel{
	//		Status: "normal",
	//		Name:   "userName" + strconv.Itoa(i),
	//		Email:  "mail@email.com",
	//	}
	//	if i > 5 {
	//		newUser.Status = "forbidden"
	//	}
	//	_, err := CreateUser(newUser)
	//	if err != nil {
	//		log.Println("InsertUser error:", err)
	//		return
	//	}
	//}

	userList, err := GetUsersByStatus("forbidden", false)
	if err != nil {
		log.Println("GetUsersByStatus error:", err)
		return
	}
	log.Println(userList)

	userList, err = GetUsersByStatus("forbidden", false)
	if err != nil {
		log.Println("GetUsersByStatus error:", err)
		return
	}
	log.Println(userList)
}
