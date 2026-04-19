package translator

import (
	"encoding/json"
	"strings"
)

// OpenAIInputAudio represents an OpenAI audio payload structure.
type OpenAIInputAudio struct {
	Type       string          `json:"type"`        // "input_audio"
	InputAudio OpenAIAudioData `json:"input_audio"` // nested audio data
}

// OpenAIAudioData represents the nested data for OpenAI input_audio.
type OpenAIAudioData struct {
	Data   string `json:"data"`   // base64 encoded audio
	Format string `json:"format"` // e.g. "mp3", "wav"
}

// ConvertAudioToOpenAIFormat takes raw base64 audio and a format and returns an OpenAI input_audio part.
func ConvertAudioToOpenAIFormat(base64Audio, format string) ([]byte, error) {
	audioFormat := "wav"
	if format != "" {
		audioFormat = strings.TrimPrefix(format, "audio/")
	}

	audioPart := OpenAIInputAudio{
		Type: "input_audio",
		InputAudio: OpenAIAudioData{
			Data:   base64Audio,
			Format: audioFormat,
		},
	}
	return json.Marshal(audioPart)
}

// ParseOpenAIAudioFormat extracts base64 audio and format from an OpenAI input_audio part.
func ParseOpenAIAudioFormat(payload []byte) (string, string, bool) {
	var audioPart OpenAIInputAudio
	if err := json.Unmarshal(payload, &audioPart); err != nil {
		return "", "", false
	}

	if audioPart.Type != "input_audio" || audioPart.InputAudio.Data == "" || audioPart.InputAudio.Format == "" {
		return "", "", false
	}

	return audioPart.InputAudio.Data, audioPart.InputAudio.Format, true
}
