package easyconf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEasyconf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Easyconf Suite")
}
