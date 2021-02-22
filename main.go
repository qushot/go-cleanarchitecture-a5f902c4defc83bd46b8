package main

import (
	"fmt"

	"go-clean-architecture-sample/di"
)

func main() {
	c := di.NewContainer()
	userController := c.UserController

	fmt.Println("Enter the name of the new user.")
	fmt.Println("username:")
	fmt.Printf(">")
	var userName string
	fmt.Scan(&userName)

	if err := userController.CreateUser(userName); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}
