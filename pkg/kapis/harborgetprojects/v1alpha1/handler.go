package harborgetprojects

import (
	"strings"

	"github.com/emicklei/go-restful"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) GetProjects(req *restful.Request, resp *restful.Response) {
	s := GetAllProjects()
	str := strings.ReplaceAll(s, "\n", "")
	str = strings.ReplaceAll(str, "    ", " ")
	str = strings.ReplaceAll(str, "   ", " ")

	resp.WriteAsJson(AllProjects{
		Message: str,
	})
	resp.Write([]byte("\n"))
}

type AllProjects struct {
	Message string `json:"harbor-projects"`
}
