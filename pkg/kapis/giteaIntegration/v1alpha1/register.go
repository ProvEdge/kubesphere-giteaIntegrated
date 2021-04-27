package giteaIntegration

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
)

const (
	GroupName = "git.kubesphere.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func AddToContainer(container *restful.Container) error {

	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler()

	webservice.Route(webservice.GET("/listRepos").
		Reads("").
		To(handler.GetRepos).
		Returns(http.StatusOK, api.StatusOK, giteaResponse{})).
		Doc("Api for listing repositories")

	container.Add(webservice)


	return nil

}

