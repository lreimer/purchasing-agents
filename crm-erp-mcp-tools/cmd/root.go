package cmd

import (
	"log"
	"os"

	"github.com/lreimer/purchasing-agents/crm"
	"github.com/lreimer/purchasing-agents/erp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

var version string
var transport string
var baseURL string
var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crm-erp-mcp-tools",
	Short: "A collection of MCP tools to access a CRM and ERP system",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new MCP server
		s := server.NewMCPServer(
			"Prcuhasing Tools",
			version,
			server.WithRecovery(),
			server.WithLogging(),
		)

		// Add tools to the server
		crm.AddCrmTools(s)
		erp.AddErpTools(s)

		// Only check for "sse" since stdio is the default
		if transport == "sse" {
			sseServer := server.NewSSEServer(s, server.WithBaseURL(baseURL))
			ssePort := "0.0.0.0:" + port
			log.Printf("MCP Server (SSE) listening on %s", ssePort)
			if err := sseServer.Start(ssePort); err != nil {
				log.Fatalf("MCP Server (SSE) error: %v", err)
			}
		} else {
			if err := server.ServeStdio(s); err != nil {
				log.Fatalf("MCP Server (stdio) error: %v", err)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// SetVersion set the application version to be used in the MCP server
func SetVersion(v string) {
	version = v
}

func init() {
	rootCmd.Flags().StringVarP(&transport, "transport", "t", "stdio", "Transport to use. Valid options: stdio, sse")
	rootCmd.Flags().StringVarP(&baseURL, "url", "u", "http://localhost:8001", "The public SSE base URL to use.")
	rootCmd.Flags().StringVarP(&port, "port", "p", "8001", "The local SSE server port to use.")

	// override the default port with ENV if specified
	// use port parameter as default
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort
	}
	// override the default baseURL with ENV if specified
	// use baseURL parameter as default
	if envBaseURL, ok := os.LookupEnv("BASE_URL"); ok {
		baseURL = envBaseURL
	}
}
