package kubeappsIntegration

import "github.com/emicklei/go-restful"

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) getAppRepos(req *restful.Request, resp *restful.Response) {
	setToken("eyJhbGciOiJSUzI1NiIsImtpZCI6ImJtRDI3T1RUVnljZEJmRmF6bEtZZkRUbXkyZTlnZGJnTkxZbXcwOGdHLVUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yLXRva2VuLXdyOHRnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZGI2NGJjMzQtMjhkYi00MTZjLWFjZjEtZDczZTMzMGU2MWNjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6a3ViZWFwcHMtb3BlcmF0b3IifQ.s6Z634t_EHZ7o_v1fdpxSjcW2R79dkXV1aku0bsYTN36xFV2voykVez6NQqkOxzj8T3f4pz-8L4vnQ4pEOZ7ZBsRuYhCUkXBv1H9f0k0lwzewoh05Z_5qluX2fDsS2gXAOybUb4hPcZnJclHhE4xndcaosktvkM_reF3xhaAm5_OM48g-43oDDUuPD3Mj5vkaMRI7CUGhdkAt9xtSbQnRGtxNdR6jrpGnT6RmVxinAxHxqxtA9GoTdMBcKjRVBOPnlzuXYhdPtX_N_Yt4eLeuxpbS99aCya3XjyGIc68gPc0h4MUZItha3JldabAi1xZtQsvbvF90FmNVjyrvMsx8Q")	
	clustername := req.PathParameter("cluster")
	resp.Write([]byte(getApprepositories(clustername)))
}

func (h handler) getReleased(req *restful.Request, resp *restful.Response) {
	setToken("eyJhbGciOiJSUzI1NiIsImtpZCI6ImJtRDI3T1RUVnljZEJmRmF6bEtZZkRUbXkyZTlnZGJnTkxZbXcwOGdHLVUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yLXRva2VuLXdyOHRnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZGI2NGJjMzQtMjhkYi00MTZjLWFjZjEtZDczZTMzMGU2MWNjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6a3ViZWFwcHMtb3BlcmF0b3IifQ.s6Z634t_EHZ7o_v1fdpxSjcW2R79dkXV1aku0bsYTN36xFV2voykVez6NQqkOxzj8T3f4pz-8L4vnQ4pEOZ7ZBsRuYhCUkXBv1H9f0k0lwzewoh05Z_5qluX2fDsS2gXAOybUb4hPcZnJclHhE4xndcaosktvkM_reF3xhaAm5_OM48g-43oDDUuPD3Mj5vkaMRI7CUGhdkAt9xtSbQnRGtxNdR6jrpGnT6RmVxinAxHxqxtA9GoTdMBcKjRVBOPnlzuXYhdPtX_N_Yt4eLeuxpbS99aCya3XjyGIc68gPc0h4MUZItha3JldabAi1xZtQsvbvF90FmNVjyrvMsx8Q")
	clustername := req.PathParameter("cluster")
	resp.Write([]byte(releasedApps(clustername)))
}

func (h handler) getChartsInfo(req *restful.Request, resp *restful.Response) {
	setToken("eyJhbGciOiJSUzI1NiIsImtpZCI6ImJtRDI3T1RUVnljZEJmRmF6bEtZZkRUbXkyZTlnZGJnTkxZbXcwOGdHLVUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yLXRva2VuLXdyOHRnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZGI2NGJjMzQtMjhkYi00MTZjLWFjZjEtZDczZTMzMGU2MWNjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6a3ViZWFwcHMtb3BlcmF0b3IifQ.s6Z634t_EHZ7o_v1fdpxSjcW2R79dkXV1aku0bsYTN36xFV2voykVez6NQqkOxzj8T3f4pz-8L4vnQ4pEOZ7ZBsRuYhCUkXBv1H9f0k0lwzewoh05Z_5qluX2fDsS2gXAOybUb4hPcZnJclHhE4xndcaosktvkM_reF3xhaAm5_OM48g-43oDDUuPD3Mj5vkaMRI7CUGhdkAt9xtSbQnRGtxNdR6jrpGnT6RmVxinAxHxqxtA9GoTdMBcKjRVBOPnlzuXYhdPtX_N_Yt4eLeuxpbS99aCya3XjyGIc68gPc0h4MUZItha3JldabAi1xZtQsvbvF90FmNVjyrvMsx8Q")
	clustername := req.PathParameter("cluster")
	namespace := req.PathParameter("namespace")
	resp.Write([]byte(getCharts(clustername, namespace)))
}

func (h handler) getSpecifiedReleased(req *restful.Request, resp *restful.Response) {
	setToken("eyJhbGciOiJSUzI1NiIsImtpZCI6ImJtRDI3T1RUVnljZEJmRmF6bEtZZkRUbXkyZTlnZGJnTkxZbXcwOGdHLVUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yLXRva2VuLXdyOHRnIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Imt1YmVhcHBzLW9wZXJhdG9yIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZGI2NGJjMzQtMjhkYi00MTZjLWFjZjEtZDczZTMzMGU2MWNjIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6a3ViZWFwcHMtb3BlcmF0b3IifQ.s6Z634t_EHZ7o_v1fdpxSjcW2R79dkXV1aku0bsYTN36xFV2voykVez6NQqkOxzj8T3f4pz-8L4vnQ4pEOZ7ZBsRuYhCUkXBv1H9f0k0lwzewoh05Z_5qluX2fDsS2gXAOybUb4hPcZnJclHhE4xndcaosktvkM_reF3xhaAm5_OM48g-43oDDUuPD3Mj5vkaMRI7CUGhdkAt9xtSbQnRGtxNdR6jrpGnT6RmVxinAxHxqxtA9GoTdMBcKjRVBOPnlzuXYhdPtX_N_Yt4eLeuxpbS99aCya3XjyGIc68gPc0h4MUZItha3JldabAi1xZtQsvbvF90FmNVjyrvMsx8Q")
	clustername := req.PathParameter("cluster")
	namespace := req.PathParameter("namespace")
	appname := req.PathParameter("app")
	resp.Write([]byte(getReleasedApp(clustername, namespace, appname)))
}

type response struct {
	Message string `json:"response: "`
}
