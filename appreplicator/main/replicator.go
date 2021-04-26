package main

import (
	"fmt"
	"github.com/litesoft-go/utils/appreplicator/copier"
	"github.com/litesoft-go/utils/fs"
	"github.com/litesoft-go/utils/iso8601"
	"github.com/litesoft-go/utils/simpleconf"
	"github.com/litesoft-go/utils/strs"
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
	srcDirs := config.ValuesFor("sources")
	if len(srcDirs) == 0 {
		panic("no 'sources' specified")
	}
	dstDirs := config.ValuesFor("destinations")
	if len(dstDirs) == 0 {
		panic("no 'destinations' specified")
	}
	ignoreExtensions := config.ValuesFor("ignoreExtensions")
	for i, ext := range ignoreExtensions {
		if !strings.HasPrefix(ext, ".") {
			ignoreExtensions[i] = "." + ext
		} else if ext == "." {
			panic("'ignoreExtensions' contained just a dot '.'")
		}
	}
	if -1 == strs.FindIn(".tmp", ignoreExtensions) {
		fmt.Println("*** Warning ***: '.tmp' not in 'ignoreExtensions'; this may cause problems")
	}
	fmt.Println("  config:")
	fmt.Println("         sources (dirs): ", srcDirs)
	fmt.Println("    destinations (dirs): ", dstDirs)
	fmt.Println("       ignoreExtensions: ", ignoreExtensions)

	srcFSs, err := fileSystemsFor("sources", srcDirs)
	if err != nil {
		panic(err)
	}
	dstFSs, err := fileSystemsFor("destinations", dstDirs)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("  Note: files beginning with '_try_' are temporary destination Dir files!")
	fmt.Println()

	zMover := copier.NewCopier(srcFSs, dstFSs, ignoreExtensions)

	err = zMover.CopyFiles()

	fmt.Println()
	fmt.Println()

	if err != nil {
		panic(err)
	}
}

func fileSystemsFor(what string, dirs []string) (fileSystems []*fs.FS, err error) {
	var zFS *fs.FS
	for _, dir := range dirs {
		zFS, err = fs.NewFS(dir)
		if err != nil {
			err = fmt.Errorf("%s: %w", what, err)
			return
		}
		fileSystems = append(fileSystems, zFS)
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
