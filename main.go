package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/video", func(c *fiber.Ctx) error {
		return playVideoHandler(c)
	})

	log.Fatal(app.Listen(":3000"))
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
