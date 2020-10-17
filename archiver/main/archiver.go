package main

import (
	"fmt"
	"github.com/litesoft-go/utils/archiver/mover"
	"github.com/litesoft-go/utils/fs"
	"github.com/litesoft-go/utils/iso8601"
	"github.com/litesoft-go/utils/simpleconf"
	"github.com/litesoft-go/utils/strs"
	"os"
	"strings"
	"time"
)

// Main function
func main() {
	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		fmt.Println("No CLI parameters supported, but got: ", argsWithProg)
		os.Exit(1)
	}
	name := cleanAppName(os.Args[0])
	fmt.Println(name, " vs 1.2 ", iso8601.ZmillisNow())

	configFile, err := os.Open(name + ".conf")
	if err != nil {
		panic(err)
	}
	config, err := simpleconf.Load(configFile)
	if err != nil {
		panic(err)
	}
	loopAfter, err := config.ValueOf("loopAfter")
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
	fmt.Println("              loopAfter: ", loopAfter)
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
	loopAfterDuration, err := durationFor("loopAfter", loopAfter)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("  Note: files beginning with '_try__' and more '-'s, are temporary")
	fmt.Println("        destination Dir files.  As such the shortest file name that")
	fmt.Println("         can be transferred is 6 characters long!")
	fmt.Println()

	zMover := mover.NewMover(srcFSs, dstFSs, ignoreExtensions)

	err = zMover.MoveFiles()
	for err == nil {
		if loopAfterDuration == nil {
			os.Exit(0)
		}
		time.Sleep(*loopAfterDuration)
		err = zMover.MoveFiles()
	}

	fmt.Println()
	fmt.Println()
	panic(err)
}

func durationFor(what, durationStr string) (duration *time.Duration, err error) {
	if durationStr == "" {
		return
	}
	var d time.Duration
	d, err = time.ParseDuration(durationStr)
	if err != nil {
		err = fmt.Errorf("'%s': %w", what, err)
	} else {
		duration = &d
	}
	return
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
