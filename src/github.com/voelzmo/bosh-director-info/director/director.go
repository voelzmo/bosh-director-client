package director

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/voelzmo/bosh-director-info/api"
)

type Director interface {
	Status() api.Status
}

type director struct {
	target     string
	rootCAPath string
}

func NewDirector(target string, rootCAPath string) Director {
	return &director{target, rootCAPath}
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

func (d *director) Login() bool {
	directorStatus := d.Status()
	authURL := directorStatus.UserAuthentication.Options["url"]

	client := &http.Client{}

	req, _ := http.NewRequest("POST", authURL, nil)
	// ...
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("accept", "application/json;charset=utf-8")
	resp, _ := client.Do(req)

	return true
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
