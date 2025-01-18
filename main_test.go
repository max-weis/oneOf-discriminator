package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostJobs(t *testing.T) {
	e := echo.New()
	server := &Server{}
	RegisterHandlers(e, server)

	tests := []struct {
		name       string
		payload    string
		statusCode int
		expected   string
	}{
		{
			name: "Valid DataProcessingJob",
			payload: `{
				"id": "123",
				"parameters": {
					"jobType": "dataProcessing",
					"dataset": "dataset-001",
					"algorithm": "algorithm-x"
				}
			}`,
			statusCode: http.StatusOK,
			expected:   `{"message":"Job accepted","jobId":"123"}`,
		},
		{
			name: "Valid NotificationJob",
			payload: `{
				"id": "456",
				"parameters": {
					"jobType": "notification",
					"recipient": "user@example.com",
					"message": "Your job is complete."
				}
			}`,
			statusCode: http.StatusOK,
			expected:   `{"message":"Job accepted","jobId":"456"}`,
		},
		{
			name: "Invalid Job Type",
			payload: `{
				"id": "789",
				"parameters": {
					"jobType": "unknownType",
					"data": "invalid"
				}
			}`,
			statusCode: http.StatusBadRequest,
			expected:   `{"error":"Invalid job parameters: unknown discriminator value: unknownType"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewReader([]byte(test.payload)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, server.PostJobs(c)) {
				assert.Equal(t, test.statusCode, rec.Code)
				assert.JSONEq(t, test.expected, rec.Body.String())
			}
		})
	}
}
