package format

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestTemplateFixture(t *testing.T) {
	gunit.Run(new(TemplateFixture), t)
}

type TemplateFixture struct {
	*gunit.Fixture
}

func (this *TemplateFixture) TestNamedPlaceholder() {
	template := "Hi, {you}! My name is {me}."
	result := Template(template, Values{"you": "Mickey", "me": "Donald"})
	this.So(result, should.Equal, "Hi, Mickey! My name is Donald.")
}

func (this *TemplateFixture) SkipTestAdvancedFormatting() {
	// python: "| {a:<10} | {b:_^+20} | {c:05} |".format(a=42, b=234234, c=3)
	// result: "| 42         | ______+234234_______ | 00003 |"
	result := Template("| {a:<10} | {b:_^+20} | {c:05} |", Values{"a": 42, "b": 234234, "c": 3})
	this.So(result, should.Equal, "| 42         | ______+234234_______ | 00003 |")
}
