package nameify

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Nameifier(t *testing.T) {
	Convey("When naming things", t, func() {
		nameifier := NewNameifier()
		nameifier.LoadJsonFiles()

		Convey("it returns a properly formatted string", func() {
			result, _ := nameifier.Nameify("chaucer")

			So(result, ShouldContainSubstring, "-")
		})

		Convey("it returns the right values", func() {
			result, _ := nameifier.Nameify(fmt.Sprintf("%s-%d", "shakespeare", 0))
			So(result, ShouldEqual, "broken-beyond")

			result, _ = nameifier.Nameify(fmt.Sprintf("%s-%d", "bocaccio", 0))
			So(result, ShouldEqual, "assistant-push")

			result, _ = nameifier.Nameify(fmt.Sprintf("%s-%d", "beowulf", 0))
			So(result, ShouldEqual, "gentle-concert")

			result, _ = nameifier.Nameify(fmt.Sprintf("%s-%d", "chaucer", 0))
			So(result, ShouldEqual, "spiritual-evening")
		})
	})
}
