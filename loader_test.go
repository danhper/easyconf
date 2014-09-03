package easyconf_test

import (
	"github.com/tuvistavie/easyconf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"
)

func checkConfig(rawConfig string, parser easyconf.Parser, expected map[string]interface{}) {
	reader := strings.NewReader(rawConfig)
	config, err := parser.ParseConfig(reader)
	Expect(err).To(BeNil())
	for key, val := range expected {
		Expect(config[key]).To(Equal(val))
	}
}

func checkError(rawConfig string, parser easyconf.Parser) {
	reader := strings.NewReader(rawConfig)
	_, err := parser.ParseConfig(reader)
	Expect(err).NotTo(BeNil())
}

var _ = Describe("Loader", func() {
	Describe("Parser", func() {
		var parser easyconf.Parser

		Describe("JsonParser", func() {
			BeforeEach(func() { parser = &easyconf.JsonParser{} })

			It("should parse json", func() {
				checkConfig(`{"foo": "bar", "bar":2}`, parser, map[string]interface{}{
					"foo": "bar", "bar": float64(2),
				})
			})

			It("should fail on bad json", func() {
				checkError(`{"foo": "bar", "bar":2`, parser)
			})
		})

		Describe("YamlParser", func() {
			BeforeEach(func() { parser = &easyconf.YamlParser{} })

			It("should parse yaml", func() {
				checkConfig("foo: bar\nbar: 2", parser, map[string]interface{}{
					"foo": "bar", "bar": 2,
				})
			})

			It("should fail on bad yaml", func() {
				checkError("foo: bar\n\tbar:2", parser)
			})
		})
	})
})
