package main

import(
	"time"
)

type todo struct{
	name string
	priority string	
	deadline time.Time
	completed bool
}


var todos = make(map[string]todo)
