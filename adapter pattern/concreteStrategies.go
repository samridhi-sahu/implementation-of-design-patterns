package main

import "fmt"

// interface incompatible class

type girl struct{}

func (g *girl) hobby(s string) {
	fmt.Println("My hobby is "+ s)
}

// an adapter let you fit the incompatible class into compatible

type girlAdapter struct{
	// In reality, the adapter contains an instance of the incompatible class
	girl girl
}

// it seems like now it is a interface compatible but in reality it is just a wrap of incompatible class

func(ga *girlAdapter) passion(s string){
	ga.girl.hobby(s)
}

// interface compatible class

type boy struct{}

func (b *boy) passion(s string){
	fmt.Println("Love to become a "+ s)
}
