package main

import "github.com/spf13/cobra"

func main() {
	cmd, err := Command()
	cobra.CheckErr(err)
	cobra.CheckErr(cmd.Execute())
}
