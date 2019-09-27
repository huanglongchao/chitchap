package command

import (
	"chitchat/data"
	"log"
	"math/rand"
)

/**
	注册用户
 */

func RegisterWrap(){

	randStr := string(randNum(100000000,1000000))
	user := data.User{
		Name:     "Robot" + randStr,
		Email:    randStr + "@qq.com",
		Password: "944792",
	}
	if err := user.Create(); err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
}

func randNum(max,min int) int{
	return rand.Intn(max-min)
}
