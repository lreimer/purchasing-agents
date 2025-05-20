package crm

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

// Customer repräsentiert die Datenstruktur eines Kunden im CRM System
type Customer struct {
	CustomerID     string         `json:"customerId"`
	CompanyName    string         `json:"companyName"`
	CompanyAddress CompanyAddress `json:"companyAddress"`
	ContactPerson  ContactPerson  `json:"contactPerson"`
}

// CompanyAddress repräsentiert die Adressdaten eines Unternehmens
type CompanyAddress struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

// ContactPerson repräsentiert die Kontaktperson eines Kunden
type ContactPerson struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// ToJSON serialisiert das Customer-Struct in ein formatiertes JSON-Format mit Einrückung
// und gibt ein Tool Result zurück für die direkte Verwendung in MCP Tools
func (c *Customer) ToJSON() (*mcp.CallToolResult, error) {
	jsonData, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return mcp.NewToolResultError("Fehler bei der JSON-Serialisierung: " + err.Error()), err
	}
	return mcp.NewToolResultText(string(jsonData)), nil
}
