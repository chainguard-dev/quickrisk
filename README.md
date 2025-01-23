# QuickRisk: Dead Simple Threat Modeling Tool

This Go-based program processes YAML configuration files to generate risk assessments and dependency graphs for components.

Need to write a threat model quickly but don't want to learn complicated tools? Have you been asked to start recording a risk register? Just start a brain dump of your components in YAML format, with a list of risks for each component:


```
components:
  github:
    risks:
      - Account takeover:
      - Improperly configured GitHub action:
  google:
    risks:
      - Account takeover:
      - Improperly configured GitHub action:
  CloudSQL:
    risks:
      - Access misconfiguration:
  frontend:
    risks:
      - XSS
      - Denial of Service
```

That's all you need. You can even go further if you like with risk scoring, dependencies, and mitigations if you like:

```
  frontend:
    risks:
      - XSS:
        risk: 2
        likelihood: 100
      - Denial of Service:
        risk: 2
        likelihood: 1
        - mitigations:
          Cloudflare: -2
    deps:
      - CloudSQL
      - google
```

Now you'll be able to generate pretty GraphViz diagrams.

## Features

- Merges configurations from multiple YAML files or directories.
- Calculates risk scores
- Outputs dependencies and risks in various formats.
- Identifies high-risk components and marks them in visualizations (e.g., red nodes in DOT format).
It supports multiple output formats, including:

- Text-based risk summaries
- CSV for tabular risk data
- Graphviz DOT for dependency visualization
- Open Threat Modeling (OTM) JSON format
- Threagile JSON format


## Usage

### Basic Command Structure

```bash
./quickrisk [--csv | --dot | --otm | --threagile] <yaml_file_or_directory>...
```

### Command-line Flags

- `--csv`: Output risk data in CSV format.
- `--dot`: Output dependency graph in Graphviz DOT format.
- `--otm`: Output risk data in Open Threat Modeling (OTM) JSON format.
- `--threagile`: Output risk data in Threagile JSON format.

If no flag is specified, the program outputs text-based risk summaries to the console.

### Examples

#### Text Output

To process a directory of YAML files and display risk summaries:

```bash
./quickrisk ./configs
```

#### CSV Output

To generate a CSV file with risk scores:

```bash
./quickrisk --csv ./configs > risks.csv
```

#### Graphviz DOT Output

To visualize dependencies in Graphviz format:

```bash
./yaml-parser --dot ./configs > dependencies.dot
```

You can render the DOT file using `dot`:

```bash
dot -Tpng dependencies.dot -o dependencies.png
```

#### OTM JSON Output

To generate an OTM JSON file:

```bash
./yaml-parser --otm ./configs > output.otm.json
```

#### Threagile JSON Output

To generate a Threagile JSON file:

```bash
./yaml-parser --threagile ./configs > output.threagile.json
```

### Input Handling

- The program accepts one or more file or directory paths.
- Directories are recursively scanned for `.yaml` and `.yml` files.
- Non-YAML files are ignored.

## Risk Scoring

Each risk is scored using the formula:

```
risk_score = (impact - 2) + (likelihood - 2)
```

High-risk components (score >= 3) are highlighted in visual outputs (e.g., red nodes in DOT).

## Requirements

- Go 1.18 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/chainguard-dev/quickrisk.git
   cd quickrisk
   ```

2. Build the binary:

   ```bash
   go build .
   ```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.
