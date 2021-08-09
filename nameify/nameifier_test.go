package nameify

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Nameifier(t *testing.T) {
	Convey("When naming things", t, func() {
		nameifier := NewNameifier()

		Convey("it returns a properly formatted string", func() {
			result, _ := nameifier.Nameify("chaucer")

			So(result, ShouldContainSubstring, "-")
		})

		Convey("it returns the right values", func() {
			result, _ := nameifier.Nameify("shakespeare")
			So(result, ShouldEqual, "mere-diamond")

			result, _ = nameifier.Nameify("bocaccio")
			So(result, ShouldEqual, "outside-clerk")

			result, _ = nameifier.Nameify("beowulf")
			So(result, ShouldEqual, "curly-bus")

			result, _ = nameifier.Nameify("chaucer")
			So(result, ShouldEqual, "lost-demand")
		})
	})
}
