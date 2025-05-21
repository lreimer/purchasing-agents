package erp

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

// Order repräsentiert eine Bestellung im ERP System
type Order struct {
	OrderID     string `json:"orderId"`
	Description string `json:"description"`
	Destination string `json:"destination"`
	Status      string `json:"status"`
}

// OrderList repräsentiert eine Liste von Bestellungen
type OrderList []*Order

// ToJSON serialisiert eine OrderList in ein formatiertes JSON-Format mit Einrückung
// und gibt ein Tool Result zurück für die direkte Verwendung in MCP Tools
func (ol OrderList) ToJSON() (*mcp.CallToolResult, error) {
	jsonData, err := json.MarshalIndent(ol, "", "    ")
	if err != nil {
		return mcp.NewToolResultError("Fehler bei der JSON-Serialisierung: " + err.Error()), err
	}
	return mcp.NewToolResultText(string(jsonData)), nil
}
