package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func (s *Server) PostJobs(ctx echo.Context) error {
	var job Job
	if err := ctx.Bind(&job); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	parameters, err := job.Parameters.ValueByDiscriminator()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid job parameters: " + err.Error()})
	}

	switch p := parameters.(type) {
	case DataProcessingJob:
		fmt.Printf("Processing data job with dataset %s and algorithm %s\n", p.Dataset, p.Algorithm)
	case NotificationJob:
		fmt.Printf("Sending notification to %s with message: %s\n", p.Recipient, p.Message)
	default:
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Unknown job type"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Job accepted", "jobId": job.Id})
}

func main() {
	e := echo.New()
	server := &Server{}
	RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":8080"))
}
