package chatgpt

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	
	if client.shortcutName == "" {
		t.Error("shortcutName should not be empty")
	}
	
	expectedName := "Ask ChatGPT on Mac"
	if client.shortcutName != expectedName {
		t.Errorf("Expected shortcutName to be %q, got %q", expectedName, client.shortcutName)
	}
}

func TestSendPromptValidation(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	
	// Test with empty prompt
	_, err := client.SendPrompt(ctx, "", false)
	if err == nil {
		t.Error("Expected error for empty prompt, got nil")
	}
} 