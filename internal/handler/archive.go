package handler

import (
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
        frequencies := calculateFrequencies(string(data))        
        tree := buildHuffmanTree(frequencies)
    },
}

func calculateFrequencies(text string) map[rune]int {
    frequencies := make(map[rune]int)
    for _, char := range text {
        frequencies[char]++
    }
    return frequencies
}

func buildHuffmanTree(frequencies map[rune]int) *model.Node {
    heap := make(model.PriorityQueue, 0, len(frequencies))
    
    for char, freq := range frequencies {
        heap.Add(&model.Node{Char: char, Freq: freq})
    }

    for heap.Len() > 1 {
        left := heap.Pop()
        right := heap.Pop()
        node := &model.Node{
            Freq:  left.Freq + right.Freq,
            Right: right,
            Left:  left,
        }
        heap.Add(node)
    }

    return heap.Pop()
}

func init() {
    archiveCmd.Flags().StringVarP(&output, "output", "o", ".", "Output path for the created file")
    rootCmd.AddCommand(archiveCmd)
}
