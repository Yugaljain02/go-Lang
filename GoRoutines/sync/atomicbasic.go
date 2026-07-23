func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 2000; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	wg.Wait()

	fmt.Println("Final Counter =", counter)
}