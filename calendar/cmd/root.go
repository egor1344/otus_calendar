package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd - cobra command
var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar micorservice",
}

func init() {
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(GrpcClientCmd)
}
