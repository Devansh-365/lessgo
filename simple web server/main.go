package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"log"
	"net/http"
)

func main() {

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	// for {
	// 	secret := getRandomNumber()

	// 	//Get User Input
	// 	guess := getUserInput()

	// 	//Comparing 
	// 	matching := isMatching(secret, guess)
	// 	fmt.Println(matching)
	// }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successfull \n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "hello")
}

// func isMatching(secret, guess int) bool {
// 	if guess > secret {
// 		fmt.Println("Your guess is too big")
// 		return false
// 	} else if guess < secret {
// 		fmt.Println("Your guess is too small")
// 		return false
// 	} else {
// 		fmt.Println("Your guess is correct")
// 		return true
// 	}
// }

// func getUserInput() int {
// 	fmt.Print("Please input the guess number : ")
// 	var input int
// 	_, err := fmt.Scan(&input)
// 	if err != nil {
// 		fmt.Println("Something is wrong")
// 	} else {
// 		fmt.Println("Your guessed number : ", input)
// 	}

// 	return input
// }

// func getRandomNumber() int {
// 	rand.Seed(time.Now().Unix())
// 	return rand.Int() % 11
// }