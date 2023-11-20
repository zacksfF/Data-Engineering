package main

import (
	"fmt"
	"time"
)

func main() {

	Start := time.Now()
	userName := fetchUser()
	respch := make(chan any)
	
	go fetchUserLikes(userName)
	//go fetchUserLikes(userName)
	go fetchUserMatch(userName)

	//fmt.Println("likes", likes)
	fmt.Println("matches", match)
	fmt.Println("matches", match)
	fmt.Println("Tooks", time.Since(Start))

}
func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Saif"
}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150)
	return 11
}

func fetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 200)
	return "Zack"
}
