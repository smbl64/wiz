package patmgr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	mgr := New("./testdata")
	list, err := mgr.List()

	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"pat1", "pat2", "pat3"}, list)
}

func TestLoadOk(t *testing.T) {
	mgr := New("./testdata")
	content, err := mgr.Load("pat1")
	assert.NoError(t, err)
	assert.Equal(t, "I am pattern 1\n", content)
}

func TestLoadNotExist(t *testing.T) {
	mgr := New("./testdata")
	_, err := mgr.Load("pat-non-existing")
	assert.Error(t, err)
}
