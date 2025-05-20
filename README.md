# Purchasing Agents

A collection of purchasing agents that can help with customer enquiries.

## Project Setup

```bash
# initialize project with deps and tools
uv init
uv add google-adk
uv tool install google-adk

# run and open the demo console
# the agent only supports voice
adk web
open http://localhost:8000
```

## ERM and CRM MCP Tools

```bash
goreleaser build --snapshot --clean
goreleaser release --skip=publish --snapshot --clean
```

If you want to use the tool locally, e.g. with Claude Desktop, use the following
configuration for the MCP server.

```json
{
    "mcpServers": {
      "enterprise": {
        "command": "/Users/mario-leander.reimer/Applications/crm-erp-mcp-tools",
        "args": ["--transport", "stdio"],
        "env": {
        }
      }
    }
}
```

Alternatively, you can use the MCP introspector for easy local development:
```bash
# as stdio binary
npx @modelcontextprotocol/inspector go run main.go

# as SSE server using 
go run main.go --transport sse
npx @modelcontextprotocol/inspector npx mcp-remote@next http://localhost:8001/sse
npx @modelcontextprotocol/inspector
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the GPL open source license, read the `LICENSE` file for details.
