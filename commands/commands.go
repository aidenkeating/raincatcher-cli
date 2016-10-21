package commands

import "fmt"

func getRepoLists() (appRepoList, moduleRepoList []string) {
	return []string{"raincatcher-demo-auth", "raincatcher-demo-cloud",
			"raincatcher-demo-mobile", "raincatcher-demo-portal"},
		[]string{"raincatcher-appform", "raincatcher-signature",
			"raincatcher-mediator", "raincatcher-risk-assessment",
			"raincatcher-sync", "raincatcher-template-build",
			"raincatcher-user", "raincatcher-vehicle-inspection",
			"raincatcher-workflow", "raincatcher-workorder",
			"raincatcher-result", "raincatcher-message",
			"raincatcher-map", "raincatcher-schedule",
			"raincatcher-file", "raincatcher-camera",
			"raincatcher-analytics"}
}

func getRepoInfo(repoPathTemplate, appRepo string) (string, string) {
	repoURLTemplate := "https://github.com/feedhenry-raincatcher/%s"

	repoURL := fmt.Sprintf(repoURLTemplate, appRepo)
	repoPath := fmt.Sprintf(repoPathTemplate, appRepo)

	return repoURL, repoPath
}
