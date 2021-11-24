package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	// "time"
)

var (
	counter   int32          // counter is a variable incremented by all goroutines.
	waitGroup sync.WaitGroup // wg is used to wait for the program to finish.
	mutex     sync.Mutex     // mutex is used to define a critical section of code.
)

func main() {

	// go responseSize("https://coderwall.com")
	// go responseSize("https://stackoverflow.com")

	//Calling time.Sleep(10 * time.Second) will make the main goroutine to sleep for 10 seconds.
	//time.Sleep(30 * time.Second)

	//using an unbuffered channel to recieve values from goroutines
	nums := make(chan int)

	waitGroup.Add(3)
	fmt.Println("Start main")

	go responseSize("https://www.golangprograms.com", nums)

	go responseSizeWait("https://www.golangprograms.com")

	go increment("Go Programming Language")

	//Calling waitGroup.Wait() will make sure all the goroutines finish before the function is excecuted
	//so there's no need to make the function sleep by calling sleep
	waitGroup.Wait()

	//Note: Seems to access channels, the `<-` sign is always used
	fmt.Println(<-nums)

	fmt.Println("End main")
	//Closing the channel
	close(nums)
}

//Added the go keyword before each call of function responseSize.
//The three responseSize goroutines starts up concurrently and three calls to http.
//Get are made concurrently as well. The program doesn't wait until one response comes back before sending out the next request.
//As a result the three response sizes are printed much sooner using goroutines.

func responseSize(url string, nums chan int) {

	defer waitGroup.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	nums <- len(body)
	fmt.Println("Step4: ", len(body))
}

//Waiting for a gorouting to stop using WaitGroup
// var waitGroup sync.WaitGroup

func responseSizeWait(url string) {
	defer waitGroup.Done()

	fmt.Println("Getting url: ", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Defer closing the response body")
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))

}

//sycn.Mutex is used to define critical sections that must must be carried out by one
//goroutine at a time
func increment(lang string) {
	defer waitGroup.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}
