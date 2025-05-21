[![Build and Test MCP Tools](https://github.com/lreimer/purchasing-agents/actions/workflows/build-mcp.yml/badge.svg)](https://github.com/lreimer/purchasing-agents/actions/workflows/build-mcp.yml)
[![Docker Publish MCP Tools](https://github.com/lreimer/purchasing-agents/actions/workflows/docker-publish-mcp.yml/badge.svg)](https://github.com/lreimer/purchasing-agents/actions/workflows/docker-publish-mcp.yml)

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

## Cloud Run Deployment

```bash
# make sure you enabled these Google APIs
gcloud services enable cloudbuild.googleapis.com artifactregistry.googleapis.com
gcloud services enable run.googleapis.com

# deploy the MCP tool server first
gcloud run deploy crm-erp-mcp-tools --source=crm-erp-mcp-tools/ \
  --region=europe-north1 \
  --min=1 \
  --port=8001 --allow-unauthenticated \
  --set-env-vars=BASE_URL=https://crm-erp-mcp-tools-343509396461.europe-north1.run.app

# now we deploy the agent with UI
export GOOGLE_API_KEY=<insert here>

# we use plain gcloud CLI to perform the deployment
gcloud run deploy purchasing-agents --source=. \
  --region=europe-north1 \
  --min=1 \
  --port=8000 --allow-unauthenticated \
  --set-env-vars=MCP_SERVER_URL=https://crm-erp-mcp-tools-343509396461.europe-north1.run.app/sse,GOOGLE_API_KEY=$GOOGLE_API_KEY,GOOGLE_GENAI_USE_VERTEXAI=FALSE

# there is a `adk deploy cloud_run command`
# however, it lacks the possibility to provide ENV variables
# also, the --with_ui option does not seem to have effect

# if you need to debug have a look at the logs
gcloud run services logs read crm-erp-mcp-tools --region=europe-north1
gcloud run services logs read purchasing-agents --region=europe-north1

# use these commands to delete the workloads
gcloud run services list
gcloud run services delete crm-erp-mcp-tools --async --region=europe-north1
gcloud run services delete purchasing-agents --async --region=europe-north1
```

## Maintainer

M.-Leander Reimer (@lreimer), <mario-leander.reimer@qaware.de>

## License

This software is provided under the GPL open source license, read the `LICENSE` file for details.
