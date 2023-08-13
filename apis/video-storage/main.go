package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
)

const (
	bucketName = "microservice-tutorial"
	awsRegion  = "ap-northeast-2" // Change to your region
)

func main() {
	app := fiber.New()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "4000"
	}

	app.Get("/video", videoHandler)

	log.Fatal(app.Listen(":" + port))
}

func videoHandler(c *fiber.Ctx) error {

	filePath := c.Query("path")
	if filePath == "" {
		return c.Status(fiber.StatusBadRequest).SendString("URL Param 'path' is missing")
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to initiate AWS session")
	}

	svc := s3.New(sess)
	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Unable to fetch video: %s", err.Error()))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Unable to read video: %s", err.Error()))
	}

	c.Set("Content-Type", "video/mp4")
	c.Set("Content-Length", fmt.Sprintf("%d", resp.ContentLength))
	fmt.Println(body)
	return c.Send(body)
}
