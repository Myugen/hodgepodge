package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gmailCmd = &cobra.Command{
	Use:   "gmail",
	Short: "Command to display mails on your gmail account",
	RunE:  runGmail,
}

func init() {
	gmailCmd.Flags().StringP("filter", "f", "", "Filter mails")
	rootCmd.AddCommand(gmailCmd)
}

func runGmail(cmd *cobra.Command, _ []string) error {
	filter, err := cmd.Flags().GetString("filter")
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Filter flag: %s", filter))
	return nil
}
