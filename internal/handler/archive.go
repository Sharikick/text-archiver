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
        text := string(util.ReadFile(args[0]))
        frequencies := calculateFrequencies(text)
        tree := buildHuffmanTree(frequencies)
        codes := make(map[rune]string)
        generateCodes(tree, "", codes)
        encodedText := encode(text, codes)
        fmt.Println(encodedText)
    },
}

func encode(text string, codes map[rune]string) string {
    encodedText := ""
    for _, symbol := range text {
        encodedText += codes[symbol]
    }
    return encodedText
}

func calculateFrequencies(text string) map[rune]int {
    frequencies := make(map[rune]int)
    for _, char := range text {
        frequencies[char]++
    }
    return frequencies
}

func buildHuffmanTree(frequencies map[rune]int) *model.Node {
    heap := model.PriorityQueue{}
    
    for char, freq := range frequencies {
        heap.Add(&model.Node{Char: char, Freq: freq})
    }

    for heap.Len() > 1 {
        left := heap.Pop()
        fmt.Println(left)
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

func generateCodes(node *model.Node, prefix string, codes map[rune]string) {
    if node == nil {
        return
    }
    if node.Char != 0 {
        codes[node.Char] = prefix
        return
    }
    generateCodes(node.Left, prefix+"0", codes)
    generateCodes(node.Right, prefix+"1", codes)
}

func init() {
    archiveCmd.Flags().StringVarP(&output, "output", "o", ".", "Output path for the created file")
    rootCmd.AddCommand(archiveCmd)
}
