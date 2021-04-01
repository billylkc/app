package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "For development purpose.",
	Long:  `For development purpose.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// var s string
		// if len(args) > 0 {
		// 	s = strings.Join(args, " ")
		// } else {
		// 	return fmt.Errorf("not enough argumetns\n")
		// }
		// res, err := calc.FindMembers(s)
		// if err != nil {
		// 	return err
		// }

		// fmt.Println(PrettyPrint(res))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
