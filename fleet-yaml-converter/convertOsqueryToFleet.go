package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// OsqueryPack represents the structure of the osquery pack JSON file.
type OsqueryPack struct {
	Queries map[string]QueryDetails `json:"queries"`
}

// QueryDetails represents the details of each query in the osquery pack.
type QueryDetails struct {
	Query       string `json:"query"`
	Description string `json:"description"`
	Platform    string `json:"platform"`
	Interval    string `json:"interval"` // Now we allow interval as a string
}

// FleetQuery represents the FleetDM-compatible YAML structure for a query.
type FleetQuery struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Name        string `yaml:"name"`
		Query       string `yaml:"query"`
		Description string `yaml:"description"`
		Platform    string `yaml:"platform"`
		Interval    int    `yaml:"interval"`
	} `yaml:"spec"`
}

// Default values for query platform and description
const defaultPlatform = "linux, darwin, windows"
const defaultDescription = "No description provided"

// convertOsqueryToFleet reads an osquery pack JSON (.conf) or SQL file and converts it to FleetDM YAML format.
func convertOsqueryToFleet(inputFile, outputFile string) error {
	// Read the input file
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	var fleetQueries []string

	// Determine if it's a .conf (JSON) file or .sql file
	if strings.HasSuffix(inputFile, ".conf") {
		// Process the .conf file (JSON format for osquery queries)
		var osqueryPack OsqueryPack
		if err := json.Unmarshal(inputData, &osqueryPack); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}

		// Convert each query to FleetDM format
		for name, details := range osqueryPack.Queries {
			query := FleetQuery{
				APIVersion: "v1",
				Kind:       "query",
			}
			query.Metadata.Name = name
			query.Spec.Name = name
			query.Spec.Query = details.Query
			query.Spec.Description = details.Description
			if query.Spec.Description == "" {
				query.Spec.Description = defaultDescription
			}
			query.Spec.Platform = details.Platform
			if query.Spec.Platform == "" {
				query.Spec.Platform = defaultPlatform
			}

			// Convert interval from string to int
			interval, err := strconv.Atoi(details.Interval)
			if err != nil {
				interval = 3600 // Default value in case of error
			}
			query.Spec.Interval = interval

			// Marshal to YAML
			yamlData, err := yaml.Marshal(&query)
			if err != nil {
				return fmt.Errorf("failed to marshal YAML: %w", err)
			}
			fleetQueries = append(fleetQueries, string(yamlData))
		}
	} else if strings.HasSuffix(inputFile, ".sql") {
		// Process SQL file
		queries, err := parseSQLFile(inputData)
		if err != nil {
			return fmt.Errorf("failed to parse SQL file: %w", err)
		}

		// Convert each SQL query to FleetDM format
		for i, queryStr := range queries {
			query := FleetQuery{
				APIVersion: "v1",
				Kind:       "query",
			}
			query.Metadata.Name = fmt.Sprintf("sql_query_%d", i+1) // Name based on index
			query.Spec.Name = fmt.Sprintf("sql_query_%d", i+1)
			query.Spec.Query = queryStr
			query.Spec.Description = "SQL query from .sql file"
			query.Spec.Platform = defaultPlatform
			query.Spec.Interval = 3600 // Default interval

			// Marshal to YAML
			yamlData, err := yaml.Marshal(&query)
			if err != nil {
				return fmt.Errorf("failed to marshal YAML: %w", err)
			}
			fleetQueries = append(fleetQueries, string(yamlData))
		}
	} else {
		return fmt.Errorf("unsupported file type: %s", inputFile)
	}

	// Write the output file with "---" to separate YAML documents
	outputData := []byte("---\n" + fmt.Sprintf("---\n%s", string([]byte(joinStrings(fleetQueries, "\n---\n")))))
	if err := ioutil.WriteFile(outputFile, outputData, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Converted %d queries to FleetDM format. Saved to %s.\n", len(fleetQueries), outputFile)
	return nil
}

// parseSQLFile extracts individual SQL queries from the provided SQL file content.
func parseSQLFile(fileContent []byte) ([]string, error) {
	// Remove any extra whitespace or comments (assuming SQL queries are separated by semicolons)
	content := string(fileContent)
	content = strings.TrimSpace(content)

	// Split the content into individual queries by semicolon
	queryRegex := regexp.MustCompile(`(?s)([^;]+);`)
	matches := queryRegex.FindAllStringSubmatch(content, -1)

	queries := make([]string, len(matches))
	for i, match := range matches {
		queries[i] = strings.TrimSpace(match[1]) // Trim any extra whitespace from each query
	}

	return queries, nil
}

// joinStrings joins a slice of strings with a separator
func joinStrings(strings []string, sep string) string {
	result := ""
	for i, s := range strings {
		result += s
		if i < len(strings)-1 {
			result += sep
		}
	}
	return result
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <input_file.conf/sql> <output_file.yml>\n", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if err := convertOsqueryToFleet(inputFile, outputFile); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
