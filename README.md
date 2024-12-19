# Fleet YAML Converter (JavaScript Version)

This project contains a Node.js scripts that reads an SQL/Conf file, extracts SQL queries, and converts them into Fleet YAML format. The resulting YAML can be used with FleetDM for monitoring and management.

## Requirements

To run this project, you will need the following:

- **Node.js 16.x+**: Install Node.js from the official website: [https://nodejs.org/](https://nodejs.org/).
- **npm (Node Package Manager)**: npm is bundled with Node.js and used to install dependencies.

### Installing Node.js

1. **Windows/Mac/Linux**: Follow the instructions at [Node.js Installation Instructions](https://nodejs.org/en/download/) for your operating system.
2. Verify installation in your terminal:

   ```bash
   node -v
   npm -v
### Setup

1. **Clone the repository:**

   ```bash
   git clone [https://github.com/your-repository-name.git](https://github.com/your-repository-name.git)
   cd your-repository-name
   npm install
### File Structure

* `convertOsQueryToFleet.js`: Converts `.conf` queries to Fleet YAML format.
* `convertSqlFileToYaml.js`: Converts SQL queries to Fleet YAML format.
* `README.md`: This file (you're reading it!).

### Running the Script: e.g `convertSqlFileToYaml.js`

1. Open your terminal or command prompt.
2. Navigate to the directory where you cloned the repository.
3. Run the program using the following command, providing the input SQL file and the desired output YAML file:

   ```bash
   node convertSqlFileToYaml.js <input_file.sql> <output_file.yml>

### Output

- Converted 2 queries to FleetDM format. Saved to output_file.yml. 
- The output_file.yml will contain the YAML version of the SQL queries.

**Example YAML Output:**

```yaml
 apiVersion: v1
 kind: query
 metadata:
   name: query_1
 spec:
   name: query_1
   query: "SELECT * FROM users"
   description: "Query 1"
   platform: linux, darwin, windows
   interval: 3600
---
 apiVersion: v1
 kind: query
 metadata:
  name: query_2
 spec:
   name: query_2
   query: "SELECT * FROM orders"
   description: "Query 2"
   platform: linux, darwin, windows
   interval: 3600
```
**Applying the YAML to FleetDM:**

Once you have generated the YAML file (e.g., output_file.yml), you can apply it to FleetDM using fleetctl.

Steps to apply the YAML:
* Ensure FleetDM and fleetctl are installed: If you don’t have FleetDM installed, you can follow the installation instructions in the official FleetDM documentation: FleetDM Installation: https://fleetdm.com/docs/installation
* Run the following command to apply the YAML file:
   ```bash
   fleetctl apply -f <output_file.yml>

* Replace <output_file.yml> with the path to your generated YAML file.
Example:
   ```bash
   fleetctl apply -f <output_file.yml>
   ```
This command will apply the configuration defined in your YAML file to FleetDM, and the queries will be available for monitoring.

**Troubleshooting:**

* **Error: failed to read input file:** 
    * Make sure the input file path is correct.
    * Verify that the file exists in the specified location.

* **Error: failed to create output file:** 
    * Ensure you have write permissions in the directory where you're trying to save the output YAML file. 
    * Check if the directory path is correctly specified.

## License

This project is **not licensed**. All rights are reserved.