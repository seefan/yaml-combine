# yaml-combine
Combine multiple yaml files for easy configuration management

## Merge Directories


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

