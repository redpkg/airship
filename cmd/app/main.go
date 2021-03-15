package main

import "github.com/redpkg/airship/internal/config"

func main() {
	if err := config.Init("./config.yaml"); err != nil {
		panic(err)
	}

}
