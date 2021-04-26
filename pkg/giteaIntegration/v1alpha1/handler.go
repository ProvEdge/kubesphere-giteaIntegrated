package giteaIntegration

import (
	"github.com/emicklei/go-restful"
)

type handler struct{}

type requestData struct {
	url      string `json:"url"`
	userName string `json:"username"`
	passWord string `json:"password"`
}

func newHandler() handler {
	return handler{}
}


func (h handler) GetRepos(req *restful.Request, resp *restful.Response) {
	var newData requestData
	
	newData.userName=req.HeaderParameter("usr")
	newData.passWord=req.HeaderParameter("pass")
	newData.url = req.HeaderParameter("url")

	initUser(newData.url, newData.userName, newData.passWord)
	resp.WriteAsJson(giteaResponse{Message: listAllRepos()})
}


type giteaResponse struct {
	Message string `json:"gitea-projects"`
}

