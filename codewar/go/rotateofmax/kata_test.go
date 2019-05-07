package kata_test

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n int64, exp int64) {
	var ans = MaxRot(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests MaxRot", func() {

	It("should handle basic cases", func() {
		dotest(38458215, 85821534)
		dotest(195881031, 988103115)
		dotest(896219342, 962193428)

	})
})
