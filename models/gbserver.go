package models

type OwnerInfo struct {
	Login    string `json:"login"`
	ID       int    `json:"id"`
	NodeID   string `json:"nodeId"`
	UserType string `json:"Type"`
}

type RepoResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Node_ID     string `json:"node_id"`
	Description string `json:"description"`
	OwnerInfo   OwnerInfo
}

type CreateRepoRequest struct {
	Name        string
	Description string

	//{"Name":"Hello-World","Description":"This is your first Repository",
	//"homepage":"https://github.com","private":false,"has_issues":true,"has_projects":true,"has_wiki":true}'
}

type CommitDetails struct {
	SHA string
	URL string
}

type ListBranchresponse struct {
	Name      string
	Commit    CommitDetails
	Protected bool
	// {
	// 	"Name": "master",
	// 	"Commit": {
	// 	  "SHA": "c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc",
	// 	  "URL": "https://api.github.com/Repos/admin/Hello-World/commits/c5b97d5ae6c19d5c5df71a34c7fbeeda2479ccbc"
	// 	},
	// 	"Protected": true,
}

type CreateBranchRequest struct {
	Ref string
	SHA string
	//{"Ref":"refs/heads/featureA","SHA":"aa218f56b14c9653891f9e74264a383fa43fefbd"}'
}

type CreateBranchObjectResponse struct {
	Type string
	SHA  string
	URL  string
}

type CreateBranchResponse struct {
	Ref    string
	NodeID string
	URL    string
	Object CreateBranchObjectResponse
	//	{Ref: "refs/heads/gbbranch", NodeID: "MDM6UmVmcmVmcy9oZWFkcy9mZWF0dXJlQQ==", URL: "https://api.github.com/Repos/gbUser/gbRepo/git/refs/heads/gbbranch",
	//
	//	"Object": {
	//			  "type": "Commit",
	//			  "SHA": "aa218f56b14c9653891f9e74264a383fa43fefbd",
	//			  "URL": "https://api.github.com/Repos/octocat/Hello-World/git/commits/aa218f56b14c9653891f9e74264a383fa43fefbd"
	//			},
}

type baseHeadPRResponse struct {
	Label string
	Ref   string
	SHA   string
	User  OwnerInfo
	Repo  RepoResponse
}

type PRResponse struct {
	URL          string
	ID           int
	NodeID       string
	Title        string
	Body         string
	State        string
	User         OwnerInfo
	Commits      int
	Additions    int
	Deletions    int
	ChangedFiles int
	Head         baseHeadPRResponse
	Base         baseHeadPRResponse
}

type PRRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Head  string `json:"head"`
	Base  string `json:"base"`
	State string `json:"state"`
	// '{"Title":"Amazing new feature",
	// "Body":"Please pull these awesome changes in!","head":"admin:new-feature","base":"master"}'

	//close PR
	//'{"Title":"new Title","Body":"updated Body","State":"open","base":"master"}'
}
