package stringer

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
    Use: "stringer",
    Version: version,
    Short: "stringer - simple CLI to transform and inspect strings",
    Long: `stringer CLI 

    Modify or inspect strings straight from the terminal`,
    
    Run: func (cmd *cobra.Command, args []string)  {
        
    },
}

func Execute()  {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error while executing stringer: '%s'", err)
        os.Exit(1)
    }
}
