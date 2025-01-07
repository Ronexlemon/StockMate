package main

import (
	"context"
	"fmt"
	"microservice/auth_service/database"
	"microservice/auth_service/models"
	"microservice/auth_service/service"
	"time"
)


func main(){
	fmt.Println("Starting connecting .....")
	 database.Connect()
	 user:= models.User{
		Username: "john",
		Email: "ronexlemon@gmail.com",
		Password: "John4712#",
		Role: "Admin",
	 }
	 ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
	 defer cancel()
	 //insert user //to be moved to handler
	 result,err := service.Register(user,ctx)
	 if err != nil {
		fmt.Println(err)
		}else{
			fmt.Println(result)
			}

}