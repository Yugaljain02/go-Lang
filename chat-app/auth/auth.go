package auth

import (
	"GoProject/ecommerce/internal/database"
	"fmt"
)

func Login() {
	database.ConnectDB()
	fmt.Println("User Logged In")
}
