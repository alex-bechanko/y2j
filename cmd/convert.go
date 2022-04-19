/*
Copyright © 2022 Alex Bechanko alexbechanko@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var input string
var output string
var prettify bool

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a yaml file to a json file",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		inputData, err := os.ReadFile(input)
		if err != nil {
			fmt.Printf("Error reading from input file %s: %s\n", input, err)
			os.Exit(1)
		}

		outputData, err := yaml.YAMLToJSON(inputData)
		if err != nil {
			fmt.Printf("Error converting yaml in input file %s: %s\n", input, err)
			os.Exit(1)
		}

		if prettify {
			var prettyOutputData bytes.Buffer
			if err := json.Indent(&prettyOutputData, outputData, "", "  "); err != nil {
				fmt.Printf("Error prettifying json from input file %s: %s\n", input, err)
				os.Exit(1)
			}

			outputData = prettyOutputData.Bytes()
		}

		err = os.WriteFile(output, outputData, 0666)
		if err != nil {
			fmt.Printf("Error writing to output file %s: %s\n", output, err)
			os.Exit(1)
		}

		fmt.Printf("Converted yaml file %s to json file %s\n", input, output)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVar(&input, "input", "", "Input YAML file")
	convertCmd.Flags().StringVar(&output, "output", "", "Output JSON file")
	convertCmd.Flags().BoolVar(&prettify, "prettify", false, "Prettify JSON")

	convertCmd.MarkFlagRequired("input")
	convertCmd.MarkFlagRequired("output")
}
