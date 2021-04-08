package harborgetrepos

import (
	"github.com/emicklei/go-restful"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) GetRepos(req *restful.Request, resp *restful.Response) {
	/*	s := GetAllProjects()
		var m []interface{}

		if err := json.Unmarshal([]byte(s), &m); err != nil {
			log.Fatal(err)
		}

		str := fmt.Sprintf("%v", m)
	*/
	resp.WriteAsJson(AllRepos{
		Message: GetAllRepos(),
	})
}

type AllRepos struct {
	Message string `json:"harbor-repos"`
}
