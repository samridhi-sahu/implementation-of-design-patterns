package main

// The context defines the interface of interest to clients.

type context struct{
	// The context maintains a reference to one of the strategy
    // objects. The context doesn't know the concrete class of a
    // strategy. It should work with all strategies via the
    // strategy interface.
	strategy strategy
}

// Usually the context accepts a strategy through the
// constructor, and also provides a setter so that the
// strategy can be switched at runtime.

func(c *context) set(s strategy){
    c.strategy = s
}

// The context delegates some work to the strategy object
// instead of implementing multiple versions of the
// algorithm on its own.

func(c *context) passion(s string){
	c.strategy.passion(s)
}