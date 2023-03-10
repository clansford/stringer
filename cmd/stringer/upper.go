package stringer

import (
    "fmt"

    "github.com/clansford/stringer/pkg/stringer"
    "github.com/spf13/cobra"
)

var upperCmd = &cobra.Command{
    Use: "upper",
    Aliases: []string{"up"},
    Short: "Converts string to uppercase",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        res := stringer.Upper(args[0])
        fmt.Println(res)
    },
}

func init()  {
    rootCmd.AddCommand(upperCmd)
}
