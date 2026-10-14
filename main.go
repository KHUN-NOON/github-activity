package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Action string `json:"action"`
		RefType string `json:"ref_type"`
		Commits []interface{} `json:"commits"`
	} `json:"payload`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	// 2. Fetch the data using our extracted function
	events, err := fetchGitHubActivity(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	// 3. Display the formatted activity
	fmt.Println("Output:")
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			commitCount := len(event.Payload.Commits)
			fmt.Printf("- Pushed %d commit(s) to %s\n", commitCount, event.Repo.Name)
		case "IssuesEvent":
			action := capitalize(event.Payload.Action)
			fmt.Printf("- %s an issue in %s\n", action, event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		case "ForkEvent":
			fmt.Printf("- Forked %s\n", event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("- Created %s in %s\n", event.Payload.RefType, event.Repo.Name)
		case "PullRequestEvent":
			action := capitalize(event.Payload.Action)
			fmt.Printf("- %s a pull request in %s\n", action, event.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("- Created a comment on an issue in %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s on %s\n", event.Type, event.Repo.Name)
		}
	}
}

func fetchGitHubActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	// GitHub API requires a User-Agent header
	req.Header.Set("User-Agent", "github-activity-cli")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user '%s' not found", username)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body %w", err)
	}

	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %w", err)
	}

	return events, nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}
