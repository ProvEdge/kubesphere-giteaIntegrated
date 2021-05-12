package kubeappsIntegration

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
)

const (
	GroupName = "apps.kubesphere.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func AddToContainer(container *restful.Container) error {

	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler()

	webservice.Route(webservice.GET("/getapprepos/cluster/{cluster}").
		Reads("").
		To(handler.getAppRepos).
		Returns(http.StatusOK, api.StatusOK, response{})).
		Doc("Api for listing all namespces' apprepositories")

	webservice.Route(webservice.GET("/getreleases/cluster/{cluster}").
		Reads("").
		To(handler.getReleased).
		Returns(http.StatusOK, api.StatusOK, response{})).
		Doc("listing released(deployed) apps")

	webservice.Route(webservice.GET("/getreleases/cluster/{cluster}/namespaces/{namespace}/app/{app}").
		Reads("").
		To(handler.getSpecifiedReleased).
		Returns(http.StatusOK, api.StatusOK, response{})).
		Doc("get specified released app")

	webservice.Route(webservice.GET("/getcharts/cluster/{cluster}/namespaces/{namespace}").
		Reads("").
		To(handler.getChartsInfo).
		Returns(http.StatusOK, api.StatusOK, response{})).
		Doc("get charts specified namespace")

	container.Add(webservice)

	return nil

}
