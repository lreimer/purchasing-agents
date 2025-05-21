package crm

import (
	"context"
	"log"

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
		kundenNummer := request.Params.Arguments["kundenNummer"].(string)
		if kundenNummer == "" {
			return mcp.NewToolResultError("kundenNummer is required"), nil
		}

		customer := &Customer{
			CustomerID:  kundenNummer,
			CompanyName: "QAware GmbH",
			CompanyAddress: CompanyAddress{
				Street:   "Aschauer Str. 20",
				Postcode: "81549",
				City:     "München",
				Country:  "Germany",
			},
			ContactPerson: ContactPerson{
				Firstname: "Mario-Leander",
				Lastname:  "Reimer",
				Email:     "mlr@qaware.de",
			},
		}

		log.Printf("Found customer: %s", kundenNummer)
		return customer.ToJSON()
	})
}
