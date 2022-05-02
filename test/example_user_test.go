package test

import (
	"log"
	"testing"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/examples/user_mgr"
)

func initialize_exampleuser() {
	basic.InitLogger()

	//db
	err := sqldb_plugin.Init(sqldb_plugin.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "testdb",
		UserName: "root",
		Password: "123456",
	}, basic.Logger)
	if err != nil {
		log.Fatalln("db init err", err)
	}

	// auto migrate table in db
	// please create table by yourself in real project
	sqldb_plugin.GetInstance().AutoMigrate(&user_mgr.ExampleUserModel{})

	//redis
	err = redis_plugin.Init(redis_plugin.Config{
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
	err = reference_plugin.Init()
	if err != nil {
		log.Fatalln("reference init err", err)
	}
}

func Test_UserDB(t *testing.T) {
	initialize_exampleuser()
	//important! Please create db table before you run this test.
	//add
	log.Println("---add user---")
	newUser := &user_mgr.ExampleUserModel{
		Status: "normal",
		Name:   "userName",
		Email:  "mail@email.com",
	}
	newUserInfo, err := user_mgr.CreateUser(newUser)
	if err != nil {
		log.Println("InsertUser error:", err)
		return
	}
	log.Println("newUserInfo:", newUserInfo)
	ID := newUserInfo.Id

	//get
	log.Println("---get user---")
	queryResult, err := user_mgr.QueryUser(&ID, nil, nil, nil, 0, 0, false, true)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	if len(queryResult.Users) == 0 {
		log.Println("GetUserById user not exist")
	} else {
		log.Println("userInfo:", queryResult.Users[0])
	}

	//update
	log.Println("---update user---")
	newData := map[string]interface{}{
		"status": "error",
		"name":   "userName2",
		"email":  "mail2@email.com",
	}
	err = user_mgr.UpdateUser(newData, ID)
	if err != nil {
		log.Println("UpdateUser error:", err)
		return
	}
	//get
	queryResult, err = user_mgr.QueryUser(&ID, nil, nil, nil, 0, 0, false, true)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	if len(queryResult.Users) == 0 {
		log.Println("GetUserById user not exist")
	} else {
		log.Println("userInfo:", queryResult.Users[0])
	}

	//delete
	log.Println("---delete user---")
	err = user_mgr.DeleteUser(ID)
	if err != nil {
		log.Println("DeleteUser error:", err)
		return
	}
	queryResult, err = user_mgr.QueryUser(&ID, nil, nil, nil, 0, 0, false, true)
	if err != nil {
		log.Println("GetUserById error:", err)
		return
	}
	if len(queryResult.Users) == 0 {
		log.Println("GetUserById user not exist")
	} else {
		log.Println("userInfo:", queryResult.Users[0])
	}

}
