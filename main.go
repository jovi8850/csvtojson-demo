// csvtojson converts a comma-delimited CSV file of housing data into a JSON Lines (.jl) file.
// Usage: go run main.go housesInput.csv housesOutput.jl

// Detailing that this is an executable program not just a library
package main

// Importing standard libraries needed for conversion
import (
	"encoding/csv"  // for reading CSV files
	"encoding/json" // for converting structs into JSON
	"fmt"           // for printing / writing text
	"log"           // for error messages and logging
	"os"            // for working with files and command-line arguments
	"strconv"       // for converting strings to numbers
)

// define structure of json file with approriate value formating
type House struct {
	Value    int     `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

// main function and review for Go program to review and convert csv to json

func main() {
	// construction of command-line arguments
	// required files are needed to be detailed in the command-line input
	// 2 files i.e. input and output file to be detailed
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <input.csv> <output.jl>")
	}

	// file names based on index of input in command-line saved in variables
	inputFile := os.Args[1]  // housesInput.csv
	outputFile := os.Args[2] // housesOutput.jl

	// open the input CSV file
	file, err := os.Open(inputFile) // variable for input file eval
	if err != nil {
		log.Fatal(err) // quit if file not found
	}
	defer file.Close() // make sure file closes when main() ends

	//CSV reader variable for open file
	reader := csv.NewReader(file)

	// Read all rows from the CSV into a slice of slices meaning a 2D array of strings
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err) // quit if issue with file
	}

	// Create (or overwrite) the output JSON Lines file
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err) // quit if issue with JSON file creation
	}
	defer out.Close() // close file when done

	// Loop over all rows exclude header
	for _, row := range records[1:] {
		// Convert the CSV row (all strings) into a House struct
		house := House{
			Value:    atoi(row[0]),
			Income:   atof(row[1]),
			Age:      atoi(row[2]),
			Rooms:    atoi(row[3]),
			Bedrooms: atoi(row[4]),
			Pop:      atoi(row[5]),
			HH:       atoi(row[6]),
		}

		// Turn the struct into JSON text
		jsonLine, err := json.Marshal(house)
		if err != nil {
			log.Fatal(err)
		}

		// Write the JSON to the output file
		fmt.Fprintln(out, string(jsonLine))
	}

}

// Helper function: convert string to int
func atoi(s string) int {
	i, _ := strconv.Atoi(s) // ignore errors
	return i
}

// Helper function: convert string to float64
func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64) // ignore errors
	return f
}
