package granularity

import (
	"testing"

	"github.com/grafadruid/go-druid/builder/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewSimple(t *testing.T) {
	expected := []byte(`"all"`)
	s := NewSimple()
	s.SetGranularity("all")

	built, err := Load(expected)
	assert.Nil(t, err)

	testutil.Compare(t, expected, s, built)
}
