# Patrol Install

This project is a Go-based utility for managing the installation and versioning of the Patrol CLI. It automates the process of checking if the Patrol CLI is installed, installing it if necessary, and verifying its version.

## Features
- Automatically checks if the Patrol CLI is installed.
- Installs the Patrol CLI if it is not present.
- Retrieves and parses the Patrol CLI version using semantic versioning.
- Provides compatibility checks for Patrol CLI, Flutter, and Patrol package versions.

## Prerequisites
- Go 1.24 or higher
- Access to the `dart` and `flutter` commands in your terminal
- Internet connection for downloading dependencies

## Installation
1. Install dependencies:

```bash
go mod tidy
```

2export CUSTOM_PATROL_VERSION=3.5.1. Build the project:

```bash
go build -o patrol-install
```

## Usage

Run the project using the following command:

```bash
go run main.go
```

Alternatively, if you built the project, execute the binary:

```bash
./patrol-install
```

## Environment Variables


You can set the following environment variable to customize the installation:

`CUSTOM_PATROL_VERSION`: Specify a custom version of the Patrol CLI to install. If not set, the latest version will be installed.

example:
```bash
export CUSTOM_PATROL_VERSION=3.5.1
```

## Project Structure

* `commands/`: Defines terminal commands used in the project.
* `steps/install/`: Contains logic for installing and managing the Patrol CLI.
* `utils/`: Utility functions for printing, executing commands, and managing environment variables.
* `constants/`: Contains regex patterns and compatibility tables.
