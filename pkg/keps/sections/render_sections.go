package sections

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/calebamiles/keps/pkg/keps/sections/internal/rendering"
	"github.com/calebamiles/keps/pkg/keps/states"
)

func RenderMissingForDraftlState(provider renderingInfoProvider) ([]Entry, error) {
	requiredSections := []string{
		Summary,
		Motivation,
	}

	sectionsInclude := map[string]bool{}
	for _, sectionFilename := range provider.SectionLocations() {
		sectionsInclude[NameForFilename(sectionFilename)] = true
	}

	missingEntries := []Entry{}

	// TODO this isn't really thread safe, better think about how this should work more
	// let's try granting a file lock on the metadata when we open a KEP
	// https://github.com/gofrs/flock
	var errs *multierror.Error
	for _, requiredSection := range requiredSections {
		if !sectionsInclude[requiredSection] {
			newEntry, err := Render(provider, requiredSection)
			errs = multierror.Append(errs, err)
			missingEntries = append(missingEntries, newEntry)
		}
	}

	if errs.ErrorOrNil() != nil {
		return nil, errs
	}

	return missingEntries, nil
}

func RenderMissingForProvisionalState(provider renderingInfoProvider) ([]Entry, error) {
	requiredSections := []string{
		Summary,
		Motivation,
	}

	sectionsInclude := map[string]bool{}
	for _, sectionFilename := range provider.SectionLocations() {
		sectionsInclude[NameForFilename(sectionFilename)] = true
	}

	missingEntries := []Entry{}

	// TODO this isn't really thread safe, better think about how this should work more
	// let's try granting a file lock on the metadata when we open a KEP
	// https://github.com/gofrs/flock
	var errs *multierror.Error
	for _, requiredSection := range requiredSections {
		if !sectionsInclude[requiredSection] {
			newEntry, err := Render(provider, requiredSection)
			errs = multierror.Append(errs, err)
			missingEntries = append(missingEntries, newEntry)
		}
	}

	if errs.ErrorOrNil() != nil {
		return nil, errs
	}

	return missingEntries, nil
}

func RenderMissingForImplementableState(provider renderingInfoProvider) ([]Entry, error) {
	requiredSections := []string{
		Summary,
		Motivation,
		DeveloperGuide,
		OperatorGuide,
		TeacherGuide,
		GraduationCriteria,
	}

	sectionsInclude := map[string]bool{}
	for _, sectionFilename := range provider.SectionLocations() {
		sectionsInclude[NameForFilename(sectionFilename)] = true
	}

	missingEntries := []Entry{}

	// TODO this isn't really thread safe, better think about how this should work more
	// let's try granting a file lock on the metadata when we open a KEP
	// https://github.com/gofrs/flock
	var errs *multierror.Error
	for _, requiredSection := range requiredSections {
		if !sectionsInclude[requiredSection] {
			newEntry, err := Render(provider, requiredSection)
			errs = multierror.Append(errs, err)
			missingEntries = append(missingEntries, newEntry)
		}
	}

	if errs.ErrorOrNil() != nil {
		return nil, errs
	}

	return missingEntries, nil
}

func Render(info renderingInfoProvider, name string) (Entry, error) {
	render := rendererFor[name]
	if render == nil {
		switch IsAutogenerated(name) {
		case true:
			return nil, fmt.Errorf("the %s section is autogenerated. Render should not be called directly for autogenerated sections.", name)
		case false:
			return nil, fmt.Errorf("no rendering information for section: %s. Ensure that it is not autogenerated", name)
		}
	}

	sectionContent, err := render(info)
	if err != nil {
		return nil, err
	}

	sec := &persistableSection{
		commonSectionInfo: &commonSectionInfo{
			filename:   Filename(name),
			name:       name,
			contentDir: info.ContentDir(),
			content:    sectionContent,
		},
	}

	return sec, nil
}

//TODO clean up this interface (e.g. whether it should be exported or not)
type renderingInfoProvider interface {
	Title() string
	Authors() []string
	OwningSIG() string
	State() states.Name
	ContentDir() string
	LastUpdated() time.Time
	SectionLocations() []string
}

type renderer func(rendering.InfoProvider) ([]byte, error)

var rendererFor = map[string]renderer{
	Summary:            rendering.NewSummary,
	Motivation:         rendering.NewMotivation,
	DeveloperGuide:     rendering.NewDeveloperGuide,
	OperatorGuide:      rendering.NewOperatorGuide,
	TeacherGuide:       rendering.NewTeacherGuide,
	GraduationCriteria: rendering.NewGraduationCriteria,
}
