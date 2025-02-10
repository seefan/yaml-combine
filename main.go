package main

import (
	"flag"
	"seefan/yaml-combine/combine"
)

var (
	dir     = flag.String("dir", "", "yaml文件目录")
	outFile = flag.String("out", "combine.yaml", "输出文件")
)

func main() {
	flag.Parse()
	if *dir == "" {
		panic("dir is empty")
	}
	yc := new(combine.YamlCombine)
	err := yc.CombineDir(*dir)
	if err != nil {
		panic(err)
	}
	err = yc.Save(*outFile)
	if err != nil {
		panic(err)
	}
}
