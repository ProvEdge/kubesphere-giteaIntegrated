package harborgetrepos

import (
	"strings"

	"github.com/emicklei/go-restful"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) GetRepos(req *restful.Request, resp *restful.Response) {
	s := GetProjectRepositories()
	str := strings.ReplaceAll(s, "\n", "")
	str = strings.ReplaceAll(str, "    ", " ")
	str = strings.ReplaceAll(str, "   ", " ")

	resp.WriteAsJson(AllRepos{
		Message: str,
	})
	resp.Write([]byte("\n"))
}

type AllRepos struct {
	Message string `json:"harbor-repos"`
}
