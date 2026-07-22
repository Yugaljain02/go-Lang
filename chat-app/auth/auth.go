package auth

import (
	"fmt"
	"GoProject/ecommerce/internal/database"
)

func Login() {
	database.ConnectDB()
	fmt.Println("User Logged In")
}