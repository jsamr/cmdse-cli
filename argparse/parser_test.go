package argparse

import (
	. "github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("parseArgument function", func() {
	DescribeTable("output",
		func(arg string, ttype *ContextFreeTokenType) {
			Expect(parseArgument(arg).Name()).To(Equal(ttype.Name()))
		},
		Entry("should match typical gnu explicit assignment",
			"--option=value", CfGnuExplicitAssignment),
		Entry("should match gnu explicit assignment with integer value",
			"--option=12", CfGnuExplicitAssignment),
		Entry("should match gnu long explicit assignment with hyphen",
			"--long-option=value", CfGnuExplicitAssignment),
		Entry("should match gnu long explicit assignment with underscore",
			"--long_option=value", CfGnuExplicitAssignment),
		Entry("should match gnu long explicit assignment with underscore an integer value",
			"--long_option=12", CfGnuExplicitAssignment),
		Entry("should match gnu explicit assignment with uppercase value",
			"--po=TOTO_to", CfGnuExplicitAssignment),
		Entry("should match typical one dash word",
			"-opt", CfOneDashWordAlphaNum),
		Entry("should match typical x-toolkit explicit assignment",
			"-option=value", CfX2lktExplicitAssignment),
		Entry("should match x-toolkit assignment with integer value",
			"-option=12", CfX2lktExplicitAssignment),
		Entry("should match long x-toolkit assignment with hyphen",
			"-long-option=value", CfX2lktExplicitAssignment),
		Entry("should match long x-toolkit assignment with underscore",
			"-long_option=value", CfX2lktExplicitAssignment),
		Entry("should match typical x-toolkit reverse switch",
			"+option", CfX2lktReverseSwitch),
		Entry("should match long x-toolkit reverse switch with hyphen",
			"+long-option", CfX2lktReverseSwitch),
		Entry("should match long x-toolkit reverse switch with underscore",
			"+long_option", CfX2lktReverseSwitch),
		Entry("should match end of option",
			"--", CfEndOfOptions),
		Entry("should match typical one dash letter",
			"-o", CfOneDashLetter),
		Entry("should match one dash letter number",
			"-1", CfOneDashLetter),
		Entry("should match go cli style namespaced flags",
			"-ns.flag", CfOneDashWord),
		Entry("should match typical stick value assignment",
			"-n3", CfPosixShortStickyValue),
		Entry("should match typical stick value assignment with multiple digits",
			"-n12", CfPosixShortStickyValue),
		Entry("should not be matched by sticky value",
			"-n12p", CfOneDashWordAlphaNum),
		Entry("should match typical long one dash word",
			"-long-option", CfOneDashWord),
		Entry("should match typical two dash word",
			"--option", CfTwoDashWord),
		Entry("should match typical long two dash word",
			"--long-option", CfTwoDashWord),
		Entry("should match weird word starting with two dashes that should not be matched as an option",
			"-_not_an_option", CfWord),
		Entry("should match weird word starting with one dash that should not be matched as an option",
			"--_not_an_option", CfWord),
		Entry("should match typical old-style option word",
			"word", CfOptWord),
		Entry("should match typical sentence",
			"this is a sentence", CfWord),
		Entry("should match typical path",
			"/path/to/resource", CfWord),
		Entry("should match typical url",
			"http://foo.com/bar", CfWord),
	)
})