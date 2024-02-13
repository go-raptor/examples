package main

import (
	"github.com/go-raptor/examples/api/config"
	"github.com/go-raptor/examples/api/config/initializers"
	"github.com/go-raptor/raptor"
)

func main() {
	r := raptor.NewRaptor()

	r.Init(initializers.App())
	r.Routes(config.Routes())

	r.Listen()
}
