package workflow_test

import (
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
)

var _ = Describe("Proposing a KEP", func() {
	const (
		authorOne         = "handleOne"
		title             = "A Great but Complicated Idea"
		kubernetesWideDir = "kubernetes-wide"
		metadataFilename  = "metadata.yaml"
	)

	Context("with a valid KEP", func() {
		It("proposes a KEP and returns a link to the GitHub Pull Request", func() {
			By("reading in the runtime settings")

			By("returning the URL that a human would care about")
		})
	})

	Context("when the KEP has already been proposed", func() {
		It("returns an existing link to the GitHub Pull Request", func() {
			By("reading in the runtime settings")

			By("returning the URL that a human would care about")
		})
	})
})
