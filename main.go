package main

import (
	"rust-update/analyze"
	"rust-update/fs"
)

func main() {
	toml, err := fs.GetCargoToml()
	if err != nil {
		panic(err)
	}
	s, err := analyze.Analyze(toml)
	if err != nil {
		panic(err)
	}
	err = fs.WriteUpdateSh(s)
	if err != nil {
		panic(err)
	}
}
