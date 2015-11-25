package director_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/voelzmo/bosh-director-info/director"
)

var dir director.Director

var _ = Describe("Director", func() {
	Context("When initialized with a target", func() {

		BeforeEach(func() {
			dir = director.NewDirector("https://52.2.165.66:25555")
		})

		It("Contacts the target on #status", func() {
			status := dir.Status()
			Expect(status.Name).To(Equal("my-bosh"))
		})
	})
})
