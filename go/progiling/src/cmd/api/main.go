package main

import (
	"log"
	"profiling/internal/api"
)

func main() {
	a := api.NewApi()
	log.Fatalf("%v\n", a.Start())
}
