package handler

import (
	"fmt"

	"github.com/spf13/cobra"
	"tsyden.com/archive/internal/model"
	"tsyden.com/archive/internal/util"
)

var output string
var archiveCmd = &cobra.Command{
    Use: "archive",
    Short: "",
    Long: "",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        data := util.ReadFile(args[0])
        println(string(data)) 

        heap := model.PriorityQueue{}

        heap.Add(&model.Node{
            Freq: 10,
            Char: 's',
        })

        fmt.Println(heap)
    },
}



func init() {
    archiveCmd.Flags().StringVarP(&output, "output", "o", ".", "Output path for the created file")
    rootCmd.AddCommand(archiveCmd)
}
