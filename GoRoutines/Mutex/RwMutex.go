package main

import (
	"sync"

	"fmt"
)

var price int = 500

var rw sync.RWMutex

func writeprice(Updateprice int, wg *sync.WaitGroup) {

	defer wg.Done()

	rw.Lock()

	price = Updateprice

	fmt.Println("writeprice", price)

	rw.Unlock()

}

func readprice(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	rw.Lock()

	fmt.Println("read price ", id, "Price is ", price)

	rw.Unlock()

}

func main() {

	var wg sync.WaitGroup

	wg.Add(4)

	go readprice(1, &wg)

	go readprice(2, &wg)

	go writeprice(500, &wg)

	go writeprice(500, &wg)

	wg.Wait()
}
