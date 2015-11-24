package dump

import (
	"fmt"
	//"sync"

	"github.com/google/go-github/github"
	"github.com/knakayama/ghdump/credential"
	"github.com/knakayama/ghdump/utils"
)

// Dump github repository
func DumpRepository() {
	var allRepos []github.Repository

	ghClient, user := credential.GetGithubClient()
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		repos, response, err := ghClient.Repositories.List(user, opt)
		utils.Dieif(err)

		allRepos = append(allRepos, repos...)
		if response.NextPage == 0 {
			break
		} else {
			opt.ListOptions.Page = response.NextPage
		}
	}

	showResultRepository(allRepos)

	//fmt.Println(*repo.HTMLURL)
	//wg.Wait()
	//for _, repo := range allRepos {
	//	fmt.Println(*repo.HTMLURL)
	// TODO: goroutine
	//var wg sync.WaitGroup
	//count := 0
	//for {
	//	wg.Add(count)
	//	count += 1
	//	go func() {
	//		repos, response, err := ghClient.Repositories.List(user, opt)
	//		utils.Dieif(err)

	//		allRepos = append(allRepos, repos...)
	//		if response.NextPage == 0 {
	//			return
	//		} else {
	//			opt.ListOptions.Page = response.NextPage
	//		}
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//for _, repo := range allRepos {
	//	fmt.Println(*repo.HTMLURL)
	//}
}

// Dump github starred repository
func DumpStarredRepository() {
	var allRepos []github.StarredRepository
	ghClient, user := credential.GetGithubClient()
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	// TODO: goroutine
	for {
		repos, response, err := ghClient.Activity.ListStarred(user, opt)
		utils.Dieif(err)

		allRepos = append(allRepos, repos...)
		if response.NextPage == 0 {
			break
		} else {
			opt.ListOptions.Page = response.NextPage
		}
	}

	showResultStarredRepository(allRepos)
}

func showResultRepository(allRepos []github.Repository) {
	var urlWidth int
	for _, repo := range allRepos {
		if len(*repo.HTMLURL) > urlWidth {
			urlWidth = len(*repo.HTMLURL)
		}
	}
	urlFmt := fmt.Sprintf("%%-%ds | ", urlWidth)

	for _, repo := range allRepos {
		fmt.Printf(urlFmt, *repo.HTMLURL)
		if repo.Description != nil {
			fmt.Println(*repo.Description)
		} else {
			fmt.Println()
		}
	}
}

func showResultStarredRepository(allRepos []github.StarredRepository) {
	var urlWidth int

	for _, repo := range allRepos {
		if len(*repo.Repository.HTMLURL) > urlWidth {
			urlWidth = len(*repo.Repository.HTMLURL)
		}
	}
	urlFmt := fmt.Sprintf("%%-%ds | ", urlWidth)

	for _, repo := range allRepos {
		fmt.Printf(urlFmt, *repo.Repository.HTMLURL)
		if repo.Repository.Description != nil {
			fmt.Println(*repo.Repository.Description)
		} else {
			fmt.Println()
		}
	}
}
