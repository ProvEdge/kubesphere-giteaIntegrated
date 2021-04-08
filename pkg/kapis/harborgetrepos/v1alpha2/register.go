package harborgetrepos

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

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

func AddToContainer(container *restful.Container) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler()

	webservice.Route(webservice.GET("/get-repos").
		Reads("").
		To(handler.GetRepos).
		Returns(http.StatusOK, api.StatusOK, AllRepos{})).
		Doc("Api for get-repos")

	container.Add(webservice)

	return nil
}
