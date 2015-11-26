package director

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/voelzmo/bosh-director-info/api"
)

type Director interface {
	Status() api.Status
	Login() api.Login
}

type director struct {
	target       string
	rootCAPath   string
	clientName   string
	clientSecret string
}

func NewDirector(target string, rootCAPath string, clientName string, clientSecret string) Director {
	return &director{target, rootCAPath, clientName, clientSecret}
}

func (d *director) Status() api.Status {
	var status api.Status

	directorClient := NewClient(d.rootCAPath)

	resp, err := directorClient.Get(fmt.Sprintf("%s/info", d.target))
	if err != nil {
		log.Fatal("Error getting director status: %s", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &status)

	return status
}

func (d *director) Login() api.Login {
	var auth api.Login

	directorStatus := d.Status()
	authURL := directorStatus.UserAuthentication.Options["url"]

	client := NewClient(d.rootCAPath)

	postBody := bytes.NewReader([]byte(`grant_type=client_credentials`))
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", authURL), postBody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("accept", "application/json;charset=utf-8")
	req.SetBasicAuth(d.clientName, d.clientSecret)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error logging in: %s", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &auth)

	return auth
}

//
// func (d *director) Tasks() []api.Task {
// 	var tasks []api.Task
//
// 	directorClient := NewClient(d.rootCAPath)
//
// 	resp, err := directorClient.Get(fmt.Sprintf("%s/tasks", d.target))
// 	if err != nil {
// 		log.Fatal("Error getting director task: %s", err)
// 	}
// 	defer resp.Body.Close()
//
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	json.Unmarshal(body, &tasks)
//
// 	return tasks
// }
