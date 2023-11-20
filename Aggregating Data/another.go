package main

import (
	"fmt"
	"time"
)

func main() {

	Start := time.Now()
	userName := fetchUser()
	respch := make(chan any,20 )
	
	go fetchUserLikes(userName, respch)
	//go fetchUserLikes(userName)
	go fetchUserMatch(userName, respch)

	for resp := range respch{
		fmt.Println("resp", resp)
	}


	//fmt.Println("likes", likes)
	//fmt.Println("matches", match)
	//fmt.Println("matches", match)
	fmt.Println("Tooks", time.Since(Start))

}
func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Saif"
}

func fetchUserLikes(userName string, respch chan any) {
	time.Sleep(time.Millisecond * 150)
	respch <-  11
}

func fetchUserMatch(userName string, respch chan any)  {
	time.Sleep(time.Millisecond * 200)
	respch <-  "Zack"
}
