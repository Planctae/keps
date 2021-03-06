// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package keps_test

import (
	"time"
)

type mockMetadataProvider struct {
	AddSectionsCalled chan bool
	AddSectionsInput  struct {
		Arg0 chan []string
	}
	PersistCalled chan bool
	PersistOutput struct {
		Ret0 chan error
	}
	SectionsCalled chan bool
	SectionsOutput struct {
		Ret0 chan []string
	}
	UniqueIDCalled chan bool
	UniqueIDOutput struct {
		Ret0 chan string
	}
	TitleCalled chan bool
	TitleOutput struct {
		Ret0 chan string
	}
	OwningSIGCalled chan bool
	OwningSIGOutput struct {
		Ret0 chan string
	}
	AuthorsCalled chan bool
	AuthorsOutput struct {
		Ret0 chan []string
	}
	ContentDirCalled chan bool
	ContentDirOutput struct {
		Ret0 chan string
	}
	CreatedCalled chan bool
	CreatedOutput struct {
		Ret0 chan time.Time
	}
	LastUpdatedCalled chan bool
	LastUpdatedOutput struct {
		Ret0 chan time.Time
	}
}

func newMockMetadataProvider() *mockMetadataProvider {
	m := &mockMetadataProvider{}
	m.AddSectionsCalled = make(chan bool, 100)
	m.AddSectionsInput.Arg0 = make(chan []string, 100)
	m.PersistCalled = make(chan bool, 100)
	m.PersistOutput.Ret0 = make(chan error, 100)
	m.SectionsCalled = make(chan bool, 100)
	m.SectionsOutput.Ret0 = make(chan []string, 100)
	m.UniqueIDCalled = make(chan bool, 100)
	m.UniqueIDOutput.Ret0 = make(chan string, 100)
	m.TitleCalled = make(chan bool, 100)
	m.TitleOutput.Ret0 = make(chan string, 100)
	m.OwningSIGCalled = make(chan bool, 100)
	m.OwningSIGOutput.Ret0 = make(chan string, 100)
	m.AuthorsCalled = make(chan bool, 100)
	m.AuthorsOutput.Ret0 = make(chan []string, 100)
	m.ContentDirCalled = make(chan bool, 100)
	m.ContentDirOutput.Ret0 = make(chan string, 100)
	m.CreatedCalled = make(chan bool, 100)
	m.CreatedOutput.Ret0 = make(chan time.Time, 100)
	m.LastUpdatedCalled = make(chan bool, 100)
	m.LastUpdatedOutput.Ret0 = make(chan time.Time, 100)
	return m
}
func (m *mockMetadataProvider) AddSections(arg0 []string) {
	m.AddSectionsCalled <- true
	m.AddSectionsInput.Arg0 <- arg0
}
func (m *mockMetadataProvider) Persist() error {
	m.PersistCalled <- true
	return <-m.PersistOutput.Ret0
}
func (m *mockMetadataProvider) Sections() []string {
	m.SectionsCalled <- true
	return <-m.SectionsOutput.Ret0
}
func (m *mockMetadataProvider) UniqueID() string {
	m.UniqueIDCalled <- true
	return <-m.UniqueIDOutput.Ret0
}
func (m *mockMetadataProvider) Title() string {
	m.TitleCalled <- true
	return <-m.TitleOutput.Ret0
}
func (m *mockMetadataProvider) OwningSIG() string {
	m.OwningSIGCalled <- true
	return <-m.OwningSIGOutput.Ret0
}
func (m *mockMetadataProvider) Authors() []string {
	m.AuthorsCalled <- true
	return <-m.AuthorsOutput.Ret0
}
func (m *mockMetadataProvider) ContentDir() string {
	m.ContentDirCalled <- true
	return <-m.ContentDirOutput.Ret0
}
func (m *mockMetadataProvider) Created() time.Time {
	m.CreatedCalled <- true
	return <-m.CreatedOutput.Ret0
}
func (m *mockMetadataProvider) LastUpdated() time.Time {
	m.LastUpdatedCalled <- true
	return <-m.LastUpdatedOutput.Ret0
}

type mockContentProvider struct {
	SectionsCalled chan bool
	SectionsOutput struct {
		Ret0 chan []string
	}
	EraseCalled chan bool
	EraseOutput struct {
		Ret0 chan error
	}
	PersistCalled chan bool
	PersistOutput struct {
		Ret0 chan error
	}
	ContentDirCalled chan bool
	ContentDirOutput struct {
		Ret0 chan string
	}
}

func newMockContentProvider() *mockContentProvider {
	m := &mockContentProvider{}
	m.SectionsCalled = make(chan bool, 100)
	m.SectionsOutput.Ret0 = make(chan []string, 100)
	m.EraseCalled = make(chan bool, 100)
	m.EraseOutput.Ret0 = make(chan error, 100)
	m.PersistCalled = make(chan bool, 100)
	m.PersistOutput.Ret0 = make(chan error, 100)
	m.ContentDirCalled = make(chan bool, 100)
	m.ContentDirOutput.Ret0 = make(chan string, 100)
	return m
}
func (m *mockContentProvider) Sections() []string {
	m.SectionsCalled <- true
	return <-m.SectionsOutput.Ret0
}
func (m *mockContentProvider) Erase() error {
	m.EraseCalled <- true
	return <-m.EraseOutput.Ret0
}
func (m *mockContentProvider) Persist() error {
	m.PersistCalled <- true
	return <-m.PersistOutput.Ret0
}
func (m *mockContentProvider) ContentDir() string {
	m.ContentDirCalled <- true
	return <-m.ContentDirOutput.Ret0
}
