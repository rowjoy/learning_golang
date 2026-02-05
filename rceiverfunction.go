//https://youtu.be/D0ptgkOWtaM?si=UOtpERb0OUEMVfDD

package main

import "fmt"

type UserInfo struct {
	Name    string
	Age     float64
	Village string
}

func printUserInfo(usr UserInfo) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
	fmt.Println(usr.Village)
}

func main() {
	userinfo := UserInfo{
		Name:    "Jamirul islam",
		Age:     30,
		Village: "Dhaka - Rumpura",
	}

	printUserInfo(userinfo)
	userinfo2 := UserInfo{
		Name:    "Mukbul mia",
		Age:     30,
		Village: "Dhaka - Rumpura",
	}
	printUserInfo(userinfo2)

}

//https://youtu.be/D0ptgkOWtaM?si=lLIq1CZWXxPKeMON
