package common

import (
	"testing"
)

func TestExtractUpstreamError(t *testing.T) {
	tests := []struct {
		name     string
		body     []byte
		expected string
	}{
		{
			name:     "empty body",
			body:     []byte(""),
			expected: "upstream returned empty response",
		},
		{
			name:     "standard openai error",
			body:     []byte(`{"error": {"message": "You exceeded your current quota"}}`),
			expected: "You exceeded your current quota",
		},
		{
			name:     "html error blob",
			body:     []byte(`<html><body><h1>502 Bad Gateway</h1></body></html>`),
			expected: "upstream returned non-JSON response: <html><body><h1>502 Bad Gateway</h1></body></html>",
		},
		{
			name:     "unknown json schema",
			body:     []byte(`{"weird": "schema"}`),
			expected: "upstream returned an unknown JSON error schema",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractUpstreamError(tt.body)
			if got != tt.expected {
				t.Errorf("ExtractUpstreamError() = %v, want %v", got, tt.expected)
			}
		})
	}
}
