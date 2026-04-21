package translator

import (
	"encoding/json"
	"strings"
)

// AudioToText handles multi-modal audio-to-text translations
// for Claude endpoints.
type AudioToText struct{}

// ClaudeAudioData represents an audio payload for Claude.
type ClaudeAudioData struct {
	Type      string `json:"type"`       // e.g. "base64"
	MediaType string `json:"media_type"` // e.g. "audio/mp3"
	Data      string `json:"data"`       // base64 encoded audio
}

// ClaudeDocument represents a Claude document part containing audio or other media.
type ClaudeDocument struct {
	Type   string          `json:"type"`   // "document"
	Source ClaudeAudioData `json:"source"`
}

// TranslateAudio converts raw base64 audio and format into a Claude document part.
func (a *AudioToText) TranslateAudio(base64Audio, format string) ([]byte, error) {
	mediaType := "audio/wav"
	if format != "" {
		if strings.HasPrefix(format, "audio/") {
			mediaType = format
		} else {
			mediaType = "audio/" + format
		}
	}

	doc := ClaudeDocument{
		Type: "document",
		Source: ClaudeAudioData{
			Type:      "base64",
			MediaType: mediaType,
			Data:      base64Audio,
		},
	}
	return json.Marshal(doc)
}

// Parse extracts base64 audio and format from a Claude document part.
func (a *AudioToText) Parse(payload []byte) (string, string, bool) {
	var doc ClaudeDocument
	if err := json.Unmarshal(payload, &doc); err != nil {
		return "", "", false
	}

	if doc.Type != "document" || doc.Source.Type != "base64" {
		return "", "", false
	}

	format := strings.TrimPrefix(doc.Source.MediaType, "audio/")
	if format == doc.Source.MediaType {
		// Not an audio type
		return "", "", false
	}

	return doc.Source.Data, format, true
}
