package harborgetartifacts

import (
	"strings"

	"github.com/emicklei/go-restful"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) GetArts(req *restful.Request, resp *restful.Response) {
	s := GetRepoArtifacts()
	str := strings.ReplaceAll(s, "\n", "")
	str = strings.ReplaceAll(str, "    ", " ")
	str = strings.ReplaceAll(str, "   ", " ")

	resp.WriteAsJson(AllArts{
		Message: str,
	})
	resp.Write([]byte("\n"))
}

type AllArts struct {
	Message string `json:"harbor-repos-artinfo"`
}
