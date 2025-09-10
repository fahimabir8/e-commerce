package main

import (
	"ecommerce/util"
	"fmt"
)

// "ecommerce/cmd"

func main() {
	// cmd.Serve()


	// s := "ta"
	//
	// fmt.Println(s)
	// arr := []byte(s)
	//
	// fmt.Println(arr)
	//
	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	//
	// base := enc.EncodeToString(arr)
	// fmt.Println(base)

	// lekha := []byte("Jibon")
	// d := sha256.Sum256(lekha)
	// fmt.Println(d)

	jwt,err := util.CreateJwt("dog",util.Payload{
		Sub: 10,
		FirstName: "tarique",
		LastName: "adil",
		Email: "tarique@yahoo.com",
		IsShopOwner: true,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jwt)
}
