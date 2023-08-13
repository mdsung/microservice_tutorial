package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	fmt.Println("port:", port)
	if !ok {
		port = "3000"
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/video", func(c *fiber.Ctx) error {
		return playVideoHandler(c)
	})

	log.Fatal(app.Listen(":" + port))
}

func playVideoHandler(c *fiber.Ctx) error {
	videoPath := "file_example_MP4_480_1_5MG.mp4"

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI("http://localhost:4000/video?path=" + videoPath)

	if err := a.Parse(); err != nil {
		panic(err)
	}

	code, body, errs := a.Bytes()
	if len(errs) > 0 {
		panic(errs[0])
	}
	if code != 200 {
		panic(code)
	}

	c.Set(fiber.HeaderContentType, "video/mp4")
	c.Set(fiber.HeaderContentLength, fmt.Sprintf("%d", len(body)))
	return c.Send(body)
}
