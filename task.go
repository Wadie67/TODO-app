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

func taskmain() {

	todos := make(map[string]todo)

}