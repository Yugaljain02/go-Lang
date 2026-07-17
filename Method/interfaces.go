package main
 import "fmt"

   type payment interface{
	pay()
   }

    type  UPI struct{
	 upi_no  string
    }

     func  (c UPI) pay(){
		fmt.Println("payment done with upi")
	 }

	type Card  struct{
		card_no  string
	} 

	func (c Card) pay(){
		fmt.Println("payment done with card")
	}
    type  Paypal struct{
		Email  string
	} 
	func (c Paypal) pay(){
		fmt.Println("payment done with paypal")
	}

   func Makepayment(p payment){
	p.pay()
   }


   func main(){
	upi :=UPI{ 
	 upi_no : "yugaljain",
	}

	card :=Card{
		card_no : "12343565",
	}

	paypal :=Paypal{
		Email : "apna@gmail.com",
	}

	Makepayment(paypal)
	Makepayment(upi)
	Makepayment(card)
	
 }