package testing_test

import (
	"seefan/yaml-combine/combine"
	"testing"
)

func TestCombineBytes(t *testing.T) {
	yc := new(combine.YamlCombine)
	err := yc.CombineFile("d1/1.yaml", "d1/2.yaml")
	if err != nil {
		t.Fatal(err)
	}
	b, err := yc.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func TestCombineFile(t *testing.T) {
	yc := new(combine.YamlCombine)
	err := yc.CombineFile("d1/1.yaml", "d1/2.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = yc.Save("all.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
func TestCombineDir(t *testing.T) {
	yc := new(combine.YamlCombine)
	err := yc.CombineDir("d1")
	if err != nil {
		t.Fatal(err)
	}
	err = yc.CombineDir("d2")
	if err != nil {
		t.Fatal(err)
	}
	err = yc.Save("all.yaml")
	if err != nil {
		t.Fatal(err)
	}
}
