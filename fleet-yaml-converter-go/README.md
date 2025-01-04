# Fleet YAML Converter (Go Version)

The Fleet YAML Converter is a tool written in Go that helps you convert queries from various formats (such as `.conf` for osquery packs and `.sql` files) into the FleetDM-compatible YAML format. This tool is ideal for integrating your queries with FleetDM, a modern platform for managing osquery agents at scale.

## Features

- **Supports `.conf` (osquery pack JSON) files**: Converts JSON-based osquery query packs into FleetDM YAML format.
- **Supports `.sql` files**: Converts SQL queries into FleetDM YAML format, automatically assigning metadata and default configurations.
- **Customizable**: Easily extendable to handle additional formats or modify the output as needed.
- **Simple and fast**: Written in Go for performance and ease of use.

## Prerequisites

Before you can use the Fleet YAML Converter, ensure you have the following:

- **Go**: You need to have Go installed on your system. If it's not installed, you can follow the installation guide here: [Install Go](https://golang.org/doc/install).
- **Input files**:
  - `.conf` file (osquery pack in JSON format)
  - `.sql` file (SQL queries)
  
## Installation

To install the Fleet YAML Converter on your local machine, follow these steps:

1. **Clone the repository** (or create a directory to store the code):
   ```bash
   git clone https://github.com/your-repo/fleet-yaml-converter.git
   cd fleet-yaml-converter
2. **Build the Go program**
   ```bash
   go build -o fleet-yaml-converter
This will create an executable named fleet-yaml-converter in the current directory.

**Usage**

The Fleet YAML Converter supports two types of input files:
	•	.conf files (osquery pack in JSON format)
	•	.sql files (SQL queries)

**Command Format**

To use the tool, run the following command in your terminal:

```bash
./fleet-yaml-converter <input_file.conf/sql> <output_file.yml>
```

### Key Details:
- **`<input_file.conf/sql>`** is where you specify your input file (either `.conf` or `.sql`).
- **`<output_file.yml>`** is the output YAML file you want to generate.


### Output Format:

For both .conf and .sql files, the output YAML format follows the FleetDM structure:

```bash
---
apiVersion: v1
kind: query
metadata:
  name: query_name
spec:
  name: query_name
  query: |
    SELECT * FROM some_table;
  description: "Query description"
  platform: linux, darwin, windows
  interval: 3600
---
```
Each query is converted into a FleetDM query with the following attributes:
- **apiVersion**: The version of the FleetDM API (always v1)
- **kind**: The type of object (always query)
- **metadata.name**: The name of the query
- **spec.name**: The name of the query
- **spec.query**: The SQL query or osquery query
- **spec.description**: The description of the query (default value if not provided)
- **spec.platform**: The platforms the query should run on (linux, darwin, windows by default)
- **spec.interval**: The interval at which the query runs (default is 3600 seconds)

### Customization:

You can extend the Fleet YAML Converter to handle other formats or make adjustments to the FleetDM query output. The tool is written in Go, so it’s simple to modify if you need to:
- Change default values for query intervals or descriptions.
- Add custom metadata or labels to queries.
- Add support for more input formats (e.g., CSV, XML).

Feel free to fork the repository and make adjustments based on your specific use case.


### Contributing

We welcome contributions to the Fleet YAML Converter! If you’d like to improve the tool, please:

1. Fork the repository.
2. Create a new branch for your changes.
3. Open a pull request with a description of what you’ve changed.


## License  
This project is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html).  

You are free to use, modify, and distribute this software under the terms of the GPL v3.0.