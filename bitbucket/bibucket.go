package bitbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func authSet(u string, p string, req *http.Request) {
	if len(u) > 0 && len(p) > 0 {
		req.SetBasicAuth(strings.TrimSpace(u), strings.TrimSpace(p))
	} else {
		log.Fatal("Error: username/password cannot be empty")
	}
}

func GetMergeRequests(apiurl string, workspace string, reponame string, username string, password string) map[string]interface{} {
	url := apiurl + "/repositories/" + workspace + "/" + reponame + "/pullrequests"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Cannot create http request: %v", err)
	}
	authSet(username, password, req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
