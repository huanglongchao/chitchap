package command

import (
	"chitchat/data"
	"log"
	"math/rand"
	"time"
)

/**
	注册用户
 */

func RegisterWrap() data.User{

	randStr := string(randNum(100000000,1000000))
	user := data.User{
		Name:     "Robot" + randStr,
		Email:    randStr + "@qq.com",
		Password: "944792",
	}
	if err := user.Create(); err != nil {
		log.Println("Cannot generate UUID", err)
	}
	return user
}

func randNum(max,min int) int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min)
}
