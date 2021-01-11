package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Nao funciona em arquitetura amd64")
	}
	///...
	t.Fail()
}
