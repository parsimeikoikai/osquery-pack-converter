package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
	Interval    int    `json:"interval"`
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

// convertOsqueryToFleet reads an osquery pack JSON file and converts it to FleetDM YAML format.
func convertOsqueryToFleet(inputFile, outputFile string) error {
	// Read the input file
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Parse the osquery pack JSON
	var osqueryPack OsqueryPack
	if err := json.Unmarshal(inputData, &osqueryPack); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Convert each query to FleetDM format
	var fleetQueries []string
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
		query.Spec.Interval = details.Interval

		// Marshal the query to YAML
		yamlData, err := yaml.Marshal(&query)
		if err != nil {
			return fmt.Errorf("failed to marshal YAML: %w", err)
		}
		fleetQueries = append(fleetQueries, string(yamlData))
	}

	// Write the output file with "---" to separate YAML documents
	outputData := []byte("---\n" + fmt.Sprintf("---\n%s", string([]byte(joinStrings(fleetQueries, "\n---\n")))))
	if err := ioutil.WriteFile(outputFile, outputData, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	fmt.Printf("Converted %d queries to FleetDM format. Saved to %s.\n", len(fleetQueries), outputFile)
	return nil
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
	// Check for input and output file arguments
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <input_file.json> <output_file.yml>\n", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Perform the conversion
	if err := convertOsqueryToFleet(inputFile, outputFile); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
