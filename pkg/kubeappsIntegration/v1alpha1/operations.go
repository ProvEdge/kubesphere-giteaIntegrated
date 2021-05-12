package kubeappsIntegration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var kubeappsToken string

func getApprepositories(clusterName string) string {
	if clusterName == "" {
		clusterName = "default"
	}

	url := fmt.Sprintf("http://116.203.74.185:32280/api/v1/clusters/%s/apprepositories", clusterName)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("accept", "application/json")

	req.Header.Add("Authorization", "Bearer "+kubeappsToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var dst bytes.Buffer
	json.Indent(&dst, body, "", "    ")

	return dst.String()
}

func releasedApps(clusterName string) string {
	if clusterName == "" {
		clusterName = "default"
	}

	url := fmt.Sprintf("http://116.203.74.185:32280/api/kubeops/v1/clusters/%s/releases", clusterName)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("accept", "application/json")

	req.Header.Add("Authorization", "Bearer "+kubeappsToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var dst bytes.Buffer
	json.Indent(&dst, body, "", "    ")

	return dst.String()
}

func getCharts(clusterName, namespace string) string {
	if clusterName == "" {
		clusterName = "default"
	}

	url := fmt.Sprintf("http://116.203.74.185:32280/api/assetsvc/v1/clusters/%s/namespaces/%s/charts", clusterName, namespace)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("accept", "application/json")

	req.Header.Add("Authorization", "Bearer "+kubeappsToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var dst bytes.Buffer
	json.Indent(&dst, body, "", "    ")

	return dst.String()
}

func getReleasedApp(clusterName, namespace, appName string) string {
	if clusterName == "" {
		clusterName = "default"
	}

	url := fmt.Sprintf("http://116.203.74.185:32280/api/kubeops/v1/clusters/%s/namespaces/%s/releases/%s", clusterName, namespace, appName)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("accept", "application/json")

	req.Header.Add("Authorization", "Bearer "+kubeappsToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var dst bytes.Buffer
	json.Indent(&dst, body, "", "    ")

	return dst.String()
}

func setToken(tkn string) {
	kubeappsToken = tkn
}

