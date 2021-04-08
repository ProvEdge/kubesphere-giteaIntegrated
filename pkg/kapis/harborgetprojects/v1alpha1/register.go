package harborgetprojects

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
)

const (
	GroupName = "image.kubesphere.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func AddToContainer(container *restful.Container) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler()

	webservice.Route(webservice.GET("/get-projects").
		Reads("").
		To(handler.GetProjects).
		Returns(http.StatusOK, api.StatusOK, AllProjects{})).
		Doc("Api for get-projects")

	container.Add(webservice)

	return nil
}
