package crm

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// add all CRM system tools
func AddCrmTools(s *server.MCPServer) {
	searchCustomer(s)
}

func searchCustomer(s *server.MCPServer) {
	tool := mcp.NewTool("search_customer",
		mcp.WithDescription("Sucht nach einem Kunden im CRM System"),
		mcp.WithOpenWorldHintAnnotation(true),
		mcp.WithString("kundenNummer",
			mcp.Description("Die Kundennummer des Kunden im CRM System"),
			mcp.Required(),
		),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		kundenNummer, err := request.RequireString("kundenNummer")
		if err != nil {
			return mcp.NewToolResultError("kundenNummer is required"), nil
		}

		return mcp.NewToolResultText(kundenNummer), nil
	})
}
