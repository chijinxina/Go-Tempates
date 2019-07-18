package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	toggle bool
)

var testCmd = &cobra.Command{
	Use:   "testson",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testson called the var toggle is ", toggle)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	//testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	testCmd.Flags().BoolVarP(&toggle, "toggle", "t", false, "cin toggle")
}
