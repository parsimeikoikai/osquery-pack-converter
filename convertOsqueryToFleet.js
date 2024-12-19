const fs = require('fs');
const yaml = require('js-yaml');

// Function to convert osquery pack JSON to FleetDM YAML format
function convertOsqueryToFleet(inputFile, outputFile) {
  try {
    // Read and parse the osquery pack
    const osqueryPack = JSON.parse(fs.readFileSync(inputFile, 'utf8'));

    // Validate the osquery pack
    if (!osqueryPack.queries || typeof osqueryPack.queries !== 'object') {
      throw new Error('Invalid osquery pack: Missing or invalid "queries" field.');
    }

    // Convert each query into the FleetDM format
    const fleetQueries = Object.entries(osqueryPack.queries).map(([queryName, details]) => {
      if (!details.query || !details.interval) {
        throw new Error(`Query "${queryName}" is missing required fields "query" or "interval".`);
      }

      const queryPlatform = details.platform || 'linux, darwin, windows'; // Default platform
      const queryDescription = details.description || 'No description provided'; // Default description
      const queryInterval = parseInt(details.interval, 10); // Ensure interval is an integer

      return {
        apiVersion: 'v1',
        kind: 'query',
        metadata: {
          name: queryName,
        },
        spec: {
          name: queryName,
          query: details.query,
          description: queryDescription,
          platform: queryPlatform,
          interval: queryInterval,
        },
      };
    });

    // Write the FleetDM-compatible queries to the output YAML file
    const yamlData = fleetQueries.map(query => yaml.dump(query)).join('\n---\n');
    fs.writeFileSync(outputFile, yamlData, 'utf8');
    console.log(`Converted ${fleetQueries.length} queries to FleetDM format. Saved to ${outputFile}.`);
  } catch (error) {
    console.error(`Error: ${error.message}`);
    process.exit(1);
  }
}

// Input/Output file paths
const inputFile = process.argv[2]; 
const outputFile = process.argv[3]; 

if (!inputFile || !outputFile) {
  console.error('Usage: node convertOsqueryToFleet.js <osquery_pack.json> <output_fleet.yml>');
  process.exit(1);
}

convertOsqueryToFleet(inputFile, outputFile);