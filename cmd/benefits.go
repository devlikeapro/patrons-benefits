package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shouldGiveBenefits bool
var shouldTakeAwayBenefits bool

// benefitsCmd represents the benefits command
var benefitsCmd = &cobra.Command{
	Use:   "benefits",
	Short: "Benefits control for patrons in database",
	Long:  `Give or take away benefits based on the current database state.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("benefits called")
		fmt.Println(shouldGiveBenefits)
		fmt.Println(shouldTakeAwayBenefits)
	},
}

func init() {
	rootCmd.AddCommand(benefitsCmd)
	benefitsCmd.Flags().BoolVar(
		&shouldGiveBenefits,
		"give",
		false,
		"Give benefits for these who deserve them",
	)
	benefitsCmd.Flags().BoolVar(
		&shouldTakeAwayBenefits,
		"take-away",
		false,
		"Take away benefits from people who are not patrons anymore",
	)
}
