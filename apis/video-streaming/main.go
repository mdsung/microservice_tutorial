package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// fastergoding.Run()

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
	videoPath := "videos/file_example_MP4_480_1_5MG.mp4"

	file, err := os.Open(videoPath)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Video not found")
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()

	c.Status(fiber.StatusPartialContent)
	c.Set(fiber.HeaderContentType, "video/mp4")
	c.Set("Content-Length", strconv.Itoa(int(fileSize)))

	return c.SendFile(videoPath)
}
