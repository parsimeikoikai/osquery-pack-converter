const fs = require("fs");
const yaml = require("js-yaml");

// Function to extract SQL queries from a file and convert them to YAML
function convertSqlFileToYaml(inputFile, outputFile) {
  try {
    // Read the SQL file content
    const fileContent = fs.readFileSync(inputFile, "utf8");
    console.log("File content loaded successfully.");

    // Remove SQL comments (both single-line and multi-line comments)
    const cleanContent = fileContent
      .replace(/--.*$/gm, "") // Remove single-line comments (lines starting with --)
      .replace(/\/\*[\s\S]*?\*\//g, ""); // Remove multi-line comments
    // Split queries by semicolon and trim
    const queries = cleanContent
      .split(";") // Split by semicolon
      .map((query) => query.trim()) // Trim spaces from each query
      .filter((query) => query); // Remove empty queries

    // Check if we have valid queries
    if (queries.length === 0) {
      console.log("No valid queries found in the SQL file.");
      return;
    }

    // Map queries to FleetDM format with descriptions and other metadata
    const fleetQueries = queries.map((query, index) => {
      const description = `Query ${index + 1}`; // Default description, can be customized
      return {
        apiVersion: "v1",
        kind: "query",
        metadata: {
          name: `query_${index + 1}`, // Default name if not provided
        },
        spec: {
          name: `query_${index + 1}`,
          query,
          description,
          platform: "linux, darwin, windows", // Default platform
          interval: 3600, // Default interval: 1 hour
        },
      };
    });

    // Convert to YAML and write to output file
    const yamlData = fleetQueries
      .map((query) => yaml.dump(query))
      .join("\n---\n");
    fs.writeFileSync(outputFile, yamlData, "utf8");
    console.log(
      `Converted ${fleetQueries.length} queries to FleetDM format. Saved to ${outputFile}.`
    );
  } catch (error) {
    console.error(`Error: ${error.message}`);
    process.exit(1);
  }
}

const inputFile = process.argv[2]; // Input SQL file
const outputFile = process.argv[3]; // Output YAML file

// Check if both input and output files are provided
if (!inputFile || !outputFile) {
  console.error(
    "Usage: node convertSqlFileToYaml.js <input_file.sql> <output_file.yml>"
  );
  process.exit(1);
}

// Run the conversion
convertSqlFileToYaml(inputFile, outputFile);
