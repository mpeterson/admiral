package e2e

import (
	"testing"

	"github.com/submariner-io/admiral/test/e2e/framework"
)

func init() {
	framework.ParseFlags()
}

func TestE2E(t *testing.T) {

	RunE2ETests(t)
}
