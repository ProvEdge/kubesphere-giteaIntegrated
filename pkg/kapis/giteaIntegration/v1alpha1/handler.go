package giteaIntegration

import (
	"github.com/emicklei/go-restful"
)

type handler struct{}

type requestData struct {
	url      string 
	userName string 
	passWord string 
}

func newHandler() handler {
	return handler{}
}


func (h handler) GetRepos(req *restful.Request, resp *restful.Response) {
	var newData requestData
	
	//newData.userName=req.HeaderParameter("usr")
	//newData.passWord=req.HeaderParameter("pass")
	//newData.url = req.HeaderParameter("url") these should be changed

	//initUser(newData.url, newData.userName, newData.passWord)
	newData.userName="adminDev"
	newData.passWord="provedge2021"
	newData.url="http://116.203.78.99:31024"
	initUser(newData.url,newData.userName,newData.passWord)
	resp.Write([]byte(listAllRepos()))
}


type giteaResponse struct {
	Message string `json:"gitea-projects"`
}

