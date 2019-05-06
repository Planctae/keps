package orgs

import (
	"errors"

	"github.com/calebamiles/keps/pkg/settings"
)

const (
	sandboxEnhancementsRepository = "enhancements"
	sandboxTrackingRepository     = "enhancements-tracking"
	sandboxApiReivewRepository    = "api-review"
	sandboxDefaultBranch          = "sandbox"
)

func MewSandbox(r settings.Runtime) (Instance, error) {
	owner := r.PrincipalGithubHandle()

	switch owner {
	case "":
		return nil, errors.New("no GitHub handle given")
	default:
		return &sandbox{owner: owner}, nil
	}
}

type sandbox struct {
	owner string
}

func (s *sandbox) EnhancementsRepository() string                      { return sandboxEnhancementsRepository }
func (s *sandbox) EnhancementsTrackingRepository() string              { return sandboxTrackingRepository }
func (s *sandbox) ApiReviewRepository() string                         { return sandboxApiReivewRepository }
func (s *sandbox) EnhancementsRepositoryOwner() string                 { return s.owner }
func (s *sandbox) EnhancementsTrackingRepositoryOwner() string         { return s.owner }
func (s *sandbox) ApiReviewRepositoryOwner() string                    { return s.owner }
func (s *sandbox) ApiReviewDefaultBranch() string                      { return sandboxDefaultBranch }
func (s *sandbox) EnhancementsRepositoryDefaultBranch() string         { return sandboxDefaultBranch }
func (s *sandbox) EnhancementsTrackingRepositoryDefaultBranch() string { return sandboxDefaultBranch }

func (s *sandbox) IsAuthorized(_ settings.Runtime) (bool, error) {

	// the user is always allowed to use the sandbox

	return true, nil
}
