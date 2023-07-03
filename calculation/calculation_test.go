package calculation_test

import (
	"bdd_example/calculation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculation", func() {
	Describe("Add", func() {
		It("Add two number", func() {
			result := calculation.Add(5, 3)
			Expect(result).To(Equal(8))
		})

		It("Add two negative number", func() {
			result := calculation.Add(-5, -3)
			Expect(result).To(Equal(-8))
		})

		It("Add with negative number", func() {
			result := calculation.Add(5, -3)
			Expect(result).To(Equal(2))
		})
	})
})
