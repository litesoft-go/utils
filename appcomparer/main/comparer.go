package main

import (
	"fmt"
	"github.com/litesoft-go/utils/appcomparer/comparator"
	"github.com/litesoft-go/utils/fs"
	"github.com/litesoft-go/utils/iso8601"
	"github.com/litesoft-go/utils/simpleconf"
	"os"
	"strings"
)

// Main function
func main() {
	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		fmt.Println("No CLI parameters supported, but got: ", argsWithProg)
		os.Exit(1)
	}
	name := cleanAppName(os.Args[0])
	fmt.Println(name, " vs 1.0 ", iso8601.ZmillisNow())

	configFile, err := os.Open(name + ".conf")
	if err != nil {
		panic(err)
	}
	config, err := simpleconf.Load(configFile)
	if err != nil {
		panic(err)
	}
	dir1, err := config.ValueOf("dir1")
	if err != nil {
		panic(err)
	}
	dir2, err := config.ValueOf("dir2")
	if err != nil {
		panic(err)
	}
	ignoreDirs := config.ValuesFor("ignoreDirs")
	ignoreExtensions := config.ValuesFor("ignoreExtensions")
	for i, ext := range ignoreExtensions {
		if !strings.HasPrefix(ext, ".") {
			ignoreExtensions[i] = "." + ext
		} else if ext == "." {
			panic("'extensions' contained just a dot '.'")
		}
	}
	fmt.Println("  config:")
	fmt.Println("         dir1: ", dir1)
	fmt.Println("         dir2: ", dir2)
	fmt.Println("         ignoreDirs: ", ignoreDirs)
	fmt.Println("         ignoreExtensions: ", ignoreExtensions)

	dir1FS, err := fileSystemFor("dir1", dir1)
	if err != nil {
		panic(err)
	}
	dir2FS, err := fileSystemFor("dir2", dir2)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	zComparator := comparator.NewComparator(dir1FS, dir2FS, ignoreDirs, ignoreExtensions)

	zComparator.CompareFiles()

	fmt.Println()
	fmt.Println("Done...")
}

func fileSystemFor(what string, dir string) (fileSystem *fs.FS, err error) {
	fileSystem, err = fs.NewFS(dir)
	if err != nil {
		err = fmt.Errorf("%s: %w", what, err)
	}
	return
}

func cleanAppName(name string) string {
	name = "/" + name + " "
	name = name[strings.LastIndex(name, "/")+1:]
	for name[0] == '_' {
		name = name[1:]
	}
	return strings.TrimSpace(name)
}
