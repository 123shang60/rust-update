package fs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNowPath(t *testing.T) {
	path, err := getNowPath()
	assert.NoError(t, err)
	assert.NotEqual(t, path, "")
	fmt.Println(path)
}
