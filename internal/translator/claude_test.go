package translator

import (
	"strings"
	"testing"
)

func TestClaudeAudioToText_TranslateAudio(t *testing.T) {
	b64Audio := "UklGRiQAAABXQVZFZm10IBAAAAABAAEAQB8AAEAfAAABAAgAZGF0YQAAAAA="
	format := "wav"

	a := &AudioToText{}
	bytes, err := a.TranslateAudio(b64Audio, format)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	str := string(bytes)
	if !strings.Contains(str, `"document"`) {
		t.Errorf("expected 'document' type, got %s", str)
	}
	if !strings.Contains(str, `"audio/wav"`) {
		t.Errorf("expected 'audio/wav', got %s", str)
	}
	if !strings.Contains(str, b64Audio) {
		t.Errorf("expected base64 data, got %s", str)
	}
}

func TestClaudeAudioToText_Parse(t *testing.T) {
	payload := []byte(`{
		"type": "document",
		"source": {
			"type": "base64",
			"media_type": "audio/mp3",
			"data": "base64audio"
		}
	}`)

	a := &AudioToText{}
	data, format, ok := a.Parse(payload)
	if !ok {
		t.Fatalf("expected successful parse")
	}
	if data != "base64audio" {
		t.Errorf("expected 'base64audio', got %s", data)
	}
	if format != "mp3" {
		t.Errorf("expected 'mp3', got %s", format)
	}
}
