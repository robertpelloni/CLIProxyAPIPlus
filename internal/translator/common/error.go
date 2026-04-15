package common

import (
	"encoding/json"
	"fmt"
)

// ExtractUpstreamError parses an upstream JSON response and extracts a human-readable error message.
// If the response is not valid JSON or doesn't match known schemas, it returns a generic error wrapping
// the first 200 bytes of the response for context, preventing HTML blobs from corrupting standard JSON parsers.
func ExtractUpstreamError(body []byte) string {
	if len(body) == 0 {
		return "upstream returned empty response"
	}

	var generic map[string]interface{}
	if err := json.Unmarshal(body, &generic); err != nil {
		previewLen := 200
		if len(body) < previewLen {
			previewLen = len(body)
		}
		return fmt.Sprintf("upstream returned non-JSON response: %s", string(body[:previewLen]))
	}

	// Try standard OpenAI error format
	if errMap, ok := generic["error"].(map[string]interface{}); ok {
		if msg, ok := errMap["message"].(string); ok && msg != "" {
			return msg
		}
	}

	// Try Gemini/Google error format
	if errMap, ok := generic["error"].(map[string]interface{}); ok {
		if msg, ok := errMap["message"].(string); ok && msg != "" {
			return msg
		}
		if details, ok := errMap["details"].([]interface{}); ok && len(details) > 0 {
			if detailMap, ok := details[0].(map[string]interface{}); ok {
				if reason, ok := detailMap["reason"].(string); ok && reason != "" {
					return reason
				}
			}
		}
	}

	// Try Anthropic/Claude error format
	if errMap, ok := generic["error"].(map[string]interface{}); ok {
		if msg, ok := errMap["message"].(string); ok && msg != "" {
			return msg
		}
	}

	return "upstream returned an unknown JSON error schema"
}
