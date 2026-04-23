package main

import (
	"context"
	"fmt"

	"github.com/router-for-me/CLIProxyAPI/v6/sdk/translator"
)

// dummyPlugin implements the translator.Plugin interface.
type dummyPlugin struct{}

func (p *dummyPlugin) Name() string {
	return "DummyPlugin"
}

func (p *dummyPlugin) Register(registry *translator.Registry) error {
	// Register a simple translation from "dummy" to "openai"
	registry.Register(
		translator.FromString("dummy"),
		translator.FromString("openai"),
		func(model string, rawJSON []byte, stream bool) []byte {
			// A simple request transformer: just print and return
			fmt.Printf("[DummyPlugin] Translating request for model: %s\n", model)
			return rawJSON
		},
		translator.ResponseTransform{
			NonStream: func(ctx context.Context, model string, originalRequestRawJSON, requestRawJSON, rawJSON []byte, param *any) []byte {
				fmt.Printf("[DummyPlugin] Translating non-stream response for model: %s\n", model)
				return rawJSON
			},
			Stream: func(ctx context.Context, model string, originalRequestRawJSON, requestRawJSON, rawJSON []byte, param *any) [][]byte {
				fmt.Printf("[DummyPlugin] Translating stream response for model: %s\n", model)
				return [][]byte{rawJSON}
			},
		},
	)
	return nil
}

// TranslatorPlugin is the exported symbol that CLIProxyAPI Plus will look for.
var TranslatorPlugin translator.Plugin = &dummyPlugin{}
