# mcp-server-chatgpt-app

## Prerequisite

- [ChatGPT macOS app](https://openai.com/chatgpt/download/)
- [uv](https://docs.astral.sh/uv/getting-started/installation/)
- Install the ["Ask ChatGPT on Mac" shortcuts](https://www.icloud.com/shortcuts/6ae86a39a31e4ec5938abad953ecfd64)

## Usage

### cursor

update `.mcp.json` to add the following:

```
{
    "mcpServers": {
      "chatgpt": {
        "command": "uvx",
        "args": ["mcp-server-chatgpt-app"],
        "env": {},
        "disabled": false,
        "autoApprove": []
      }
    }
}
```

### chatwise

Go to Settings -> Tools -> Add and use the following config:

```
Type: stdio
ID: ChatGPT
Command: uvx mcp-server-chatgpt-app
```

> [!CAUTION]
> Chatwise did not close mcp server even when itself is closed, which may lead to multiple mcp servers running at the same time.
> DO remember to clean them all up: `pkill -f 'mcp-server-chatgpt-app'`

### [MCP Inspector](https://github.com/modelcontextprotocol/inspector)

```
Transport type: stdio
Command: uv
Arguments: --directory /Users/<your-username>/Developer/mcp-server-chatgpt-app/src/mcp_server_chatgpt run server.py
Configuration/Request Timeout: 100000
```


## local development

```
uv --directory $HOME/Developer/mcp-server-chatgpt-app/src/mcp_server_chatgpt run server.py
```

