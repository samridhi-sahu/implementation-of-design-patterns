package main

// The client code picks a concrete strategy and passes it to
// the context. The client should be aware of the differences
// between strategies in order to make the right choice.

func main() {
	context := &context{strategy: &girl{}}
	context.passion("CTO")
	context.set(&boy{})
	context.passion("Politician")
}
