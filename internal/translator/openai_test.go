package translator

import (
	"strings"
	"testing"
)

func TestConvertAudioToOpenAIFormat(t *testing.T) {
	b64Audio := "base64encodedaudio=="
	format := "mp3"

	bytes, err := ConvertAudioToOpenAIFormat(b64Audio, format)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	str := string(bytes)
	if !strings.Contains(str, `"input_audio"`) {
		t.Errorf("expected 'input_audio' type, got %s", str)
	}
	if !strings.Contains(str, `"mp3"`) {
		t.Errorf("expected 'mp3' format, got %s", str)
	}
	if !strings.Contains(str, b64Audio) {
		t.Errorf("expected base64 data, got %s", str)
	}
}

func TestParseOpenAIAudioFormat(t *testing.T) {
	payload := []byte(`{
		"type": "input_audio",
		"input_audio": {
			"data": "base64audio==",
			"format": "wav"
		}
	}`)

	data, format, ok := ParseOpenAIAudioFormat(payload)
	if !ok {
		t.Fatalf("expected successful parse")
	}
	if data != "base64audio==" {
		t.Errorf("expected 'base64audio==', got %s", data)
	}
	if format != "wav" {
		t.Errorf("expected 'wav', got %s", format)
	}
}
