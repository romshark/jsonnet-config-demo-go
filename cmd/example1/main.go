package main

import (
	"fmt"

	"github.com/romshark/jsonnet-config-demo-go/config"
)

func main() {
	c, parsing := config.MustParse[Config]("config.jsonnet")
	fmt.Println("PARSING:", parsing)

	fmt.Println("CONFIG:")
	fmt.Printf("%#v\n", c)
}

type Config struct {
	Host   string   `json:"host" validate:"required,hostname_port"`
	Admins []string `json:"admins" validate:"required"`
}
