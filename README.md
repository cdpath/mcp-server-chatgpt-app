# mcp-server-chatgpt-app

## prerequisite

- [ChatGPT macOS app](https://openai.com/chatgpt/download/)
- Install the ["Ask ChatGPT on Mac" shortcut](https://www.icloud.com/shortcuts/6ae86a39a31e4ec5938abad953ecfd64)

## usage

### cursor

update `.mcp.json` to add the following:

```
{
    "mcpServers": {
      "chatgpt": {
        "command": "/path/to/mcp-server-chatgpt-app-darwin-arm64",
        "args": [],
        "env": {},
        "disabled": false,
        "autoApprove": []
      }
    }
}
```

### claude desktop

1. Download `mcp-server-chatgpt-app.dxt` from releases
2. Double click the `.dxt` file to install in Claude Desktop
3. `chmod +x ~/Library/Application\ Support/Claude/Claude\ Extensions/local.dxt.cdpath.mcp-server-chatgpt-app/mcp-server-chatgpt-app-darwin-arm64`


### chatwise

Go to Settings -> Tools -> Add and use the following config:

```
Type: stdio
ID: ChatGPT
Command: /path/to/mcp-server-chatgpt-app-darwin-arm64
```


## development

### build

```
make build
```

### run from source

```
make run
```

### available commands

```
make help
```

## Available Tools

- **`ask_chatgpt`**: Send a prompt to ChatGPT macOS app