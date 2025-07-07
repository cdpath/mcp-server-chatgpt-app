package mcp

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ChatGPTClient interface for dependency injection
type ChatGPTClient interface {
	SendPrompt(ctx context.Context, prompt string, waitForOutput bool) (string, error)
}

// Server represents the MCP server with ChatGPT integration
type Server struct {
	mcpServer     *server.MCPServer
	chatgptClient ChatGPTClient
}

// NewServer creates a new MCP server with ChatGPT integration
func NewServer(chatgptClient ChatGPTClient) *Server {
	s := server.NewMCPServer(
		"ChatGPT",
		"0.1.5",
	)

	srv := &Server{
		mcpServer:     s,
		chatgptClient: chatgptClient,
	}

	srv.setupTools()
	return srv
}

// setupTools configures the available MCP tools
func (s *Server) setupTools() {
	askChatGPTTool := mcp.NewTool("ask_chatgpt",
		mcp.WithDescription("Send a prompt to ChatGPT macOS app using Shortcuts"),
		mcp.WithString("prompt",
			mcp.Required(),
			mcp.Description("The text to send to ChatGPT"),
		),
		mcp.WithBoolean("wait_for_output",
			mcp.Description("Whether to wait for ChatGPT to respond"),
		),
	)

	s.mcpServer.AddTool(askChatGPTTool, s.askChatGPTHandler)
}

// askChatGPTHandler handles the ask_chatgpt tool calls
func (s *Server) askChatGPTHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments

	prompt, ok := args["prompt"].(string)
	if !ok {
		return mcp.NewToolResultError("prompt is required and must be a string"), nil
	}

	waitForOutput, _ := args["wait_for_output"].(bool)

	result, err := s.chatgptClient.SendPrompt(ctx, prompt, waitForOutput)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("ChatGPT request failed: %v", err)), nil
	}

	return mcp.NewToolResultText(result), nil
}

// Serve starts the MCP server on stdio
func (s *Server) Serve() error {
	return server.ServeStdio(s.mcpServer)
}
