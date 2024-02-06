package main

import (
	"fmt"
	"time"
)

type Task interface {
	Process()
}

type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	time.Sleep(3 * time.Second)
}

type ImageProcessingTask struct {
	ImageUrl string
}

func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImageUrl)
	time.Sleep(3 * time.Second)
}
