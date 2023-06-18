package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/farhanriziq01/websocket-heroku/module"
	"github.com/farhanriziq01/websocket-heroku/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
