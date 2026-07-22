package main
import (
	"errors" 
       "fmt"
)

var ErrDb   = errors.New("Database connection failed")
func database() error {
	return ErrDb 
}
func saveuser()error {
	err := database()
	if err!=nil{
		return fmt.Errorf("user is not saved %w", err)
	}
	return nil
}

func main(){
	err := saveuser()
	if err!=nil {
		fmt.Println("Actual Error")
		fmt.Println(err)
        fmt.Println()

		if errors.Is(err,ErrDb){
		fmt.Println("Database error occured")
	}

}
}