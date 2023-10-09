package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input       events.APIGatewayProxyRequest
		expected    string
		shouldError bool
	}{
		// Test Case 1: Valid input
		{
			input: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{
					"name": "John",
				},
			},
			expected:    "Hello John!",
			shouldError: false,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run("Test Handler", func(t *testing.T) {
			// Call the handler function with the test input
			response, err := handler(context.Background(), tc.input)

			if tc.shouldError {
				// Ensure an error occurred
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				// Ensure no error occurred
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Ensure the response matches the expected result
				if response.Body != tc.expected {
					t.Errorf("Expected %s, but got %s", tc.expected, response.Body)
				}
			}
		})
	}
}
