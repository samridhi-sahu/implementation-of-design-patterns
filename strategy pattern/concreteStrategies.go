package main

import "fmt"

// Concrete strategies implement the algorithm while following
// the base strategy interface. The interface makes them
// interchangeable in the context

type girl struct{}

func (g *girl) passion(s string) {
	fmt.Println("Want to become a "+ s)
}

type boy struct{}

func (b *boy) passion(s string){
	fmt.Println("Love to become a "+ s)
}
