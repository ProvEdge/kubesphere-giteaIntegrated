package giteaintegration

import (
	"github.com/emicklei/go-restful"
)

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) GetRepos(req *restful.Request, resp *restful.Response) {
	resp.WriteAsJson(GiteaResponse{
		Message: ListAllRepos(),
	})

}
func (h handler) GiteaAuth(req *restful.Request, resp *restful.Response) {
	InitUser("http://116.203.78.99:31024", "adminDev", "provedge2021")
	resp.WriteAsJson(GiteaResponse{Message: ListAllRepos()})
}

type GiteaResponse struct {
	Message string `json:"gitea-projects"`
}
