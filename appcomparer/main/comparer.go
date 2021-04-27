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
	name := cleanAppName(os.Args[0])
	fmt.Println(name, " vs 1.1 ", iso8601.ZmillisNow())

	args := simpleconf.WithArgs(os.Args)
	config := simpleconf.Adapter(args.Load(simpleconf.LoadFile(name + ".conf")))
	dir1 := config.ValueOfRequired("dir1")
	dir2 := config.ValueOfRequired("dir2")
	ignoreDirs := config.ValuesForOptional("ignoreDirs")
	ignoreExtensions := checkExtensions(config.ValuesForOptional("ignoreExtensions"))
	args.CheckNoMoreArgs()
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

func checkExtensions(extensions []string) []string {
	for i, ext := range extensions {
		if !strings.HasPrefix(ext, ".") {
			extensions[i] = "." + ext
		} else if ext == "." {
			panic("'ignoreExtensions' contained just a dot '.'")
		}
	}
	return extensions
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
