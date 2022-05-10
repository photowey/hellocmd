package main

import (
	"fmt"
	"log"

	"github.com/douyu/jupiter"

	"github.com/photowey/hellocmd/app/helloapps/internal/app/engine"
	"github.com/photowey/hellocmd/app/helloapps/internal/app/model"
	"github.com/photowey/hellocmd/app/helloapps/internal/app/service"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
		fmt.Println("exit jupiter app ...")
		return nil
	})

	model.Init()
	service.Init()
	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
