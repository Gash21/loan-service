package main

import (
	"github.com/Gash21/amartha-test/pkg/deps"
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func BootstrapApp(dep *deps.Dependency) *deps.Instance {

	e := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           gojson.Marshal,
		JSONDecoder:           gojson.Unmarshal,
		BodyLimit:             11 * 1024 * 1024,
	})

	e.Use(cors.New())
	e.Use(recover.New())

	instance := &deps.Instance{
		Fiber: e,
		DB:    dep.DB,
	}

	return instance
}
