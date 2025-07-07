package chatgpt

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Client handles interactions with ChatGPT through macOS Shortcuts
type Client struct {
	shortcutName string
}

// NewClient creates a new ChatGPT client
func NewClient() *Client {
	return &Client{
		shortcutName: "Ask ChatGPT on Mac",
	}
}

// SendPrompt sends a prompt to ChatGPT
func (c *Client) SendPrompt(ctx context.Context, prompt string, waitForOutput bool) (string, error) {
	// Validate input
	if strings.TrimSpace(prompt) == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}

	if waitForOutput {
		return c.sendPromptWithOutput(ctx, prompt)
	}
	return c.sendPromptWithoutOutput(ctx, prompt)
}

// sendPromptWithOutput sends a prompt and waits for ChatGPT's response
func (c *Client) sendPromptWithOutput(ctx context.Context, prompt string) (string, error) {
	cmd := exec.CommandContext(ctx, "shortcuts", "run", c.shortcutName, "-i", prompt)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("shortcuts command failed: %v", err)
	}
	return string(output), nil
}

// sendPromptWithoutOutput sends a prompt without waiting for response
func (c *Client) sendPromptWithoutOutput(ctx context.Context, prompt string) (string, error) {
	cmd := exec.CommandContext(ctx, "shortcuts", "run", c.shortcutName, "-i", prompt)
	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("failed to start shortcuts command: %v", err)
	}
	return "Sent to ChatGPT (not waiting for response)", nil
}
