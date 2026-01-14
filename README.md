# CSV to JSON Lines Converter in Go

## Overview

This repository is a clone of my orginial CSV to JSON project. AWS was used for this demonstration

This command-line Go application converts comma-delimited text files (CSV) into JSON Lines (`.jl`) format. Each line of the output file represents a valid JSON object that corresponds to one record in the input dataset. This enables analysts and engineers to easily transform raw CSV data—such as California housing prices—into JSON Lines files ready for ingestion into analytics environments or NoSQL databases.

- Command-line input and validation  
- Data transformation from CSV to JSON Lines format  
- Use of Go’s standard library for file I/O, encoding, and error handling  
- Testing and documentation practices for reproducible results  


## Program Functionality
The Go program performs the following steps:

1. Accepts command-line input for both the input CSV file and output JSON Lines file.  
2. Validates that two arguments were provided (input and output file names).  
3. Opens and reads all rows from the CSV file into memory.  
4. Converts each row into a structured Go type (`House`).  
5. Serializes each record into JSON and writes it as a single line to the output file.  
6. Closes all files safely once the process completes.  

---

## Data Description
The input dataset (`housesInput.csv`) contains California housing data derived from Miller (2015).  
Each row includes the following fields:

| Column     | Description                             | Type     |
|-------------|------------------------------------------|----------|
| value       | Median house value                      | Integer  |
| income      | Median income                           | Float    |
| age         | Median house age                        | Integer  |
| rooms       | Average number of rooms per household    | Integer  |
| bedrooms    | Average number of bedrooms per household | Integer  |
| pop         | Population count                        | Integer  |
| hh          | Number of households                    | Integer  |

---

## Command-Line Usage

### Example Command (Windows PowerShell)
powershell
go run main.go housesInput.csv housesOutput.jl

## AI Assistance Disclosure
AI tools (ChatGPT / GPT-5) were used as a learning and writing assistant throughout this project. Specifically, AI assistance included: Code Explanation, Guided Development, Debugging Help, and Documentation Support

All final code implementation, debugging validation, and testing were performed manually. 
The AI acted as an educational assistant and documentation aid.


