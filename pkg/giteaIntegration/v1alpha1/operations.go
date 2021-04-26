package giteaIntegration

import (
	"encoding/json"
	"fmt"

	g "code.gitea.io/sdk/gitea"
)

var user *g.Client
var tokenID int32 = 0
var accessToken string
var repos []RepoDataForUI

type RepoDataForUI struct {
	Username   string `json:"username,omitempty"`
	RepoName   string `json:"repoName"`
	UpdateDate string `json:"lastUpdate"`
}

func initUser(url string, username string, password string) { //burada user authorization yapiliyor
	tokenID++

	user, _ = g.NewClient(url)

	accesstknopt := g.CreateAccessTokenOption{}
	accesstknopt.Name = username + string(tokenID)
	user.SetBasicAuth(username, password)

	lat := g.ListAccessTokensOptions{}
	a, _, _ := user.ListAccessTokens(lat)
	
	if a == nil {
		tkn, resp, err := user.CreateAccessToken(accesstknopt)
		if err != nil {
			fmt.Println(err,resp)
		}
		accessToken = tkn.Token
		fmt.Println(accessToken)
	}

}

func listAllRepos() string {

	options := g.ListReposOptions{}
	options.Page = 1
	options.PageSize = 1000

	repositories,resp, err := user.ListMyRepos(options)
	repos=nil
	if err != nil {
		fmt.Println(err.Error(), resp)

	}
	var dataUI RepoDataForUI

	for _, v := range repositories {

		dataUI.RepoName = v.Name
		dataUI.UpdateDate = v.Updated.String()
		dataUI.Username = v.Owner.UserName
		repos = append(repos, dataUI)
	}

	byteJs, err1 := json.MarshalIndent(repos, "", "   ")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return string(byteJs)

}

