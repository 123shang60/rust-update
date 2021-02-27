package analyze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var input = `
[lib]
name = "test-update"
path = "src/lib.rs"

[dependencies]
test-update        = { git = "ssh://git@github.com/test.git" , branch = "main" }
`

func TestAnalyze(t *testing.T) {
	regex = `(?m)(?P<package>[a-z|0-9|\-|_]*).*git@github.com.*$`
	analyze, err := Analyze(input)
	assert.NoError(t, err)
	assert.Equal(t, analyze, "#!/bin/sh\necho \"update test-update\"\ncargo update -p test-update\n")
}
