package stringer

import (
    "fmt"

    "github.com/clansford/stringer/pkg/stringer"
    "github.com/spf13/cobra"
)

var lowerCmd = &cobra.Command{
    Use: "lower",
    Aliases: []string{"low"},
    Short: "Converts string to lowercase",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        res := stringer.Lower(args[0])
        fmt.Println(res)
    },
}

func init()  {
    rootCmd.AddCommand(lowerCmd)
}
