package hermetic

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

const (
	arbitraryBasicAuthUsername = "ossKEPtool" // Username can be anything except empty
)

// Fork is responsible for "forking" an upstream Git repository that is hosted on GitHub into the account of the authenticated GitHub user
// with the goal of proposing an isolated change to upstream. Initial upstream targets are imagined to be an "enhancements tracking"
// repository, used for managing release over release changes to Kubernetes, as well as an "API review tracking" repository used by SIG
// Architecture to manage API changes with a Kubernetes wide scope
//
// Fork:
//   - Issues a fork request to the GitHub API
//   - Clones the default branch from upstream
//   - Creates a new branch to add changes
//   - Sets the Git remote name "origin" to the forked repository
//   - Sets the Git remote name "upstream" to the upstream repository (for advanced uesrs who may desire to work with the Git repo directly)
//   - Creates callbacks to push changes to the forked repository; create a PR against the upstream repository; and to delete the forked
//     repository
func Fork(githubHandle string, token tokenProvider, owner string, repo string, toLocation string, withBranchName string) (Repo, error) {
	if _, err := os.Stat(toLocation); !os.IsNotExist(err) {
		log.Errorf("location: %s may exist already, refusing to overwrite", toLocation)
		return nil, fmt.Errorf("location: %s may exist already, refusing to overwrite", toLocation)
	}

	// call fork API
	forkUpstream, err := newCreateForkFunc(packageHttpClient, token, githubForkUrl(owner, repo))
	if err != nil {
		log.Errorf("creating fork function: %s", err)
		return nil, err
	}

	err = forkUpstream()
	if err != nil {
		log.Errorf("forking upstream %s: %s", githubForkUrl(owner, repo), err)
		return nil, err
	}

	authToken, err := token.Value()
	if err != nil {
		return nil, err
	}

	gitRepo, err := git.PlainClone(toLocation, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: arbitraryBasicAuthUsername,
			Password: authToken,
		},
		URL: githubGitUrl(owner, repo),
	})

	if err != nil {
		log.Errorf("cloning `upstream` remote: %s", err)
		return nil, err
	}

	err = gitRepo.DeleteRemote(OriginRemoteName)
	if err != nil {
		log.Errorf("deleting remote `origin`: %s", err)
		return nil, err
	}

	_, err = gitRepo.CreateRemote(&config.RemoteConfig{Name: UpstreamRemoteName, URLs: []string{githubGitUrl(owner, repo)}})
	if err != nil {
		log.Errorf("creating `upstream` remote: %s", err)
		return nil, err
	}

	// create new branch
	head, err := gitRepo.Head()
	if err != nil {
		log.Errorf("getting HEAD of repository: %s", err)
		return nil, err
	}

	refName := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", withBranchName))
	ref := plumbing.NewHashReference(refName, head.Hash())
	err = gitRepo.Storer.SetReference(ref)

	// checkout branch
	worktree, err := gitRepo.Worktree()
	if err != nil {
		log.Errorf("opening working tree of repository: %s", err)
		return nil, err
	}

	err = worktree.Checkout(&git.CheckoutOptions{Branch: refName})
	if err != nil {
		log.Errorf("checking out hermetic branch: %s", err)
		return nil, err
	}

	// construct source location
	sourceLocation := fmt.Sprintf("%s:%s", githubHandle, withBranchName)

	// create pull request callback
	createPullRequest, err := newCreatePRFunc(packageHttpClient, token, githubPrUrl(owner, repo), sourceLocation)
	if err != nil {
		log.Errorf("creating pull request creator callback: %s", err)
		return nil, err
	}

	// set origin to forked repo
	_, err = gitRepo.CreateRemote(&config.RemoteConfig{Name: OriginRemoteName, URLs: []string{githubGitUrl(githubHandle, repo)}})
	if err != nil {
		log.Errorf("setting origin URL to forked location: %s", err)
		return nil, err
	}

	// create delete repo callback
	deleteGithubRepo, err := newDeleteGithubUserRepoFunc(packageHttpClient, token, githubRepoApiUrl(githubHandle, repo))
	if err != nil {
		log.Errorf("creating delete forked repo callback: %s", err)
		return nil, err
	}

	// create push local changes callback
	pushLocalChanges, err := newPushRepoChangesFunc(gitRepo, token, withBranchName)
	if err != nil {
		log.Errorf("creating push local changes callback: %s", err)
		return nil, err
	}

	r := &repository{
		gitRepo:           gitRepo,
		createPullRequest: createPullRequest,
		deleteGithubRepo:  deleteGithubRepo,
		pushLocalChanges:  pushLocalChanges,
		branchName:        withBranchName,
		localPath:         toLocation,
		token:             token,
		locker:            &sync.Mutex{},
	}

	return r, nil
}
