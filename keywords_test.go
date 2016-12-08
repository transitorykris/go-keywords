package keywords

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMatchedUsers(t *testing.T) {
	Convey("When creating a new Keywords", t, func() {
		Convey("And it is successful", func() {
			kw := New()
			So(*kw, ShouldNotBeNil)
		})
	})

	Convey("When adding a user to a keyword", t, func() {
		Convey("And it's a new keywords", func() {
			kw := New()
			kw.Add("hello", 1)
			So(kw.kw["hello"], ShouldResemble, []int64{1})
		})
		Convey("And there is already another user", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Add("HELLo", 2)
			So(kw.kw["hello"], ShouldResemble, []int64{1, 2})
		})
		Convey("And the user already has that keyword", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Add("hELlO", 1)
			So(kw.kw["hello"], ShouldResemble, []int64{1})
			So(kw.kw["hELlO"], ShouldBeNil)
		})
	})

	Convey("When removing a user from a keyword", t, func() {
		Convey("And the keyword does not exist", func() {
			kw := New()
			kw.Remove("nothere", 1)
			So(kw.kw["nothere"], ShouldBeNil)
		})
		Convey("And the keyword exists but the user is not in it", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Remove("hello", 2)
			So(kw.kw["hello"], ShouldResemble, []int64{1})
		})
		Convey("And they are the only ones in the list", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Remove("hello", 1)
			So(kw.kw["hello"], ShouldBeNil)
		})
		Convey("And there are others in the list too", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Add("hello", 2)
			kw.Remove("hello", 1)
			So(kw.kw["hello"], ShouldResemble, []int64{2})
		})
	})

	Convey("Given a line of text", t, func() {
		Convey("And there is no matching user", func() {
			kw := New()
			users := kw.MatchedUsers("This line does not match anything")
			So(users, ShouldBeNil)
		})
		Convey("And the text is empty", func() {

		})
		Convey("And there is a single matching user", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Add("Keywords", 1)
			users := kw.MatchedUsers("Hello, Keywords!")
			So(users, ShouldResemble, []int64{1})
		})
		Convey("And there are multiple matching users", func() {
			kw := New()
			kw.Add("hello", 1)
			kw.Add("keywords", 2)
			kw.Add("keywords", 3)
			users := kw.MatchedUsers("Hello, Keywords!")
			So(users, ShouldResemble, []int64{1, 2, 3})
		})
	})
}
