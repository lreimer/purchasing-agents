package erp

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// add all ERP system tools
func AddErpTools(s *server.MCPServer) {
	getOrders(s)
}

func getOrders(s *server.MCPServer) {
	tool := mcp.NewTool("get_orders",
		mcp.WithDescription("Ruft Bestellungen aus dem ERP-System ab"),
		mcp.WithOpenWorldHintAnnotation(true),
		mcp.WithString("kundenNummer",
			mcp.Description("Die Kundennummer des Kunden im CRM System"),
			mcp.Required(),
		),
		mcp.WithString("status",
			mcp.Description("Filtert Bestellungen nach Status (z.B. waiting, offen, verschickt). Default: offen"),
			mcp.Enum("waiting", "offen", "verschickt"),
			mcp.DefaultString("offen"),
		),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		orders := OrderList{
			{
				OrderID:     "12345",
				Destination: "Rosenheim",
				Description: "Bestellung Rohre",
				Status:      "waiting",
			},
			{
				OrderID:     "98765",
				Destination: "Der Mond",
				Description: "Bestellung Raketen",
				Status:      "verschickt",
			},
			{
				OrderID:     "24680",
				Destination: "Berlin",
				Description: "Bestellung Gummi Muffen",
				Status:      "offen",
			},
		}

		kundenNummer := request.GetString("kundenNummer", "")
		if kundenNummer == "" {
			return mcp.NewToolResultError("kundenNummer is required"), nil
		}
		statusString := request.GetString("status", "offen")
		if statusString == "" {
			return mcp.NewToolResultError("status is required"), nil
		}

		filteredOrders := OrderList{}
		for _, order := range orders {
			if order.Status == statusString {
				filteredOrders = append(filteredOrders, order)
			}
		}
		orders = filteredOrders

		return orders.ToJSON()
	})
}
