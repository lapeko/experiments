package main

import (
	"fmt"
	"log"
	"profiling/internal/api"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	a := api.NewApi()
	log.Fatalf("%v\n", a.Start())
}
