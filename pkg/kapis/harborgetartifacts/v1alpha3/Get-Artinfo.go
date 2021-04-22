package harborgetartifacts

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRepoArtifacts() string {
	url := "http://116.203.78.99:30002/api/v2.0/projects/deneme/repositories/nginx/artifacts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	inputFmt := string(body)[1 : len(string(body))-2]

	fmt.Println("Got repos' artifact info.")

	return inputFmt

	//fmt.Println(string(body))

	/*FILE IO
	f, err2 := os.Create("project_repos.json")

	if err2 != nil {
		log.Fatal(err2)
	}

	defer f.Close()
	_, err3 := f.WriteString(string(body))

	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("Output in 'project_repos.json'.")*/

}
