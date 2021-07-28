package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Title     string
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues </h1>
<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
	{{end}}
</table>
`))

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%d issues:\n", result.TotalCount)
	//for _, item := range result.Items {
	//	fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	//}

}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
