package main

import (
	"fmt"
	"os"

	"github.com/cdpath/mcp-server-chatgpt-app/internal/chatgpt"
	"github.com/cdpath/mcp-server-chatgpt-app/internal/mcp"
)

func main() {
	// Create ChatGPT client
	chatgptClient := chatgpt.NewClient()

	// Create MCP server with ChatGPT client
	server := mcp.NewServer(chatgptClient)

	// Start the server
	if err := server.Serve(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
