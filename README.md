# GitHub User Activity CLI

A simple command-line interface (CLI) application built in Go to fetch and display the recent public activity of any GitHub user. 

This project is a solution to the [GitHub User Activity challenge](https://roadmap.sh/projects/github-user-activity) from roadmap.sh.

## Features

- Fetches recent public activity for a specified GitHub user using the GitHub REST API.
- Parses and formats various GitHub events into readable text, including:
  - 🚀 Push events (with commit counts)
  - 🐛 Issue events (opened, closed, etc.)
  - ⭐ Watch/Star events
  - 🍴 Fork events
  - 📝 Pull Request events
  - 💬 Issue Comment events
- Gracefully handles errors (e.g., invalid usernames, API limits, network issues).
- Clean code architecture with separated data-fetching and display logic.
- **Zero dependencies:** Built entirely using Go's standard library (`net/http`, `encoding/json`).

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or higher recommended)

## Installation

1. Clone the repository or copy the source code into a new directory.
2. Navigate to the project directory:
   ```bash
   cd github-activity
   ```
3. Build the executable:
   ```bash
   go build -o github-activity
   ```
   *(Note: On Windows, this will generate `github-activity.exe`)*

## Usage

Run the compiled binary and pass a GitHub username as an argument:

```bash
./github-activity <username>
```

### Examples

Fetching activity for a specific user:
```bash
./github-activity defunkt
```

**Example Output:**
```text
Output:
- Starred roadmapsh/roadmapsh
- Opened a pull request in user/repo
- Pushed 3 commit(s) to user/repo
- Created a comment on an issue in user/repo
```

If the user does not exist:
```bash
./github-activity non-existent-user-12345
```
**Example Output:**
```text
Error: user 'non-existent-user-12345' not found
```

## Acknowledgments

- Project idea provided by [roadmap.sh](https://roadmap.sh/projects/github-user-activity).
- Data provided by the [GitHub REST API](https://docs.github.com/en/rest).