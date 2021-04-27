package main

import (
	"fmt"
	"github.com/litesoft-go/utils/apparchiver/mover"
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
	name := cleanAppName(os.Args[0])
	fmt.Println(name, " vs 1.3 ", iso8601.ZmillisNow())

	args := simpleconf.WithArgs(os.Args)
	config := simpleconf.Adapter(args.Load(simpleconf.LoadFile(name + ".conf")))
	loopAfter := config.ValueOfOptional("loopAfter")
	srcDirs := config.ValuesForRequired("sources")
	dstDirs := config.ValuesForRequired("destinations")
	ignoreExtensions := checkExtensions(config.ValuesForOptional("ignoreExtensions"))
	args.CheckNoMoreArgs()
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

func checkExtensions(extensions []string) []string {
	for i, ext := range extensions {
		if !strings.HasPrefix(ext, ".") {
			extensions[i] = "." + ext
		} else if ext == "." {
			panic("'ignoreExtensions' contained just a dot '.'")
		}
	}
	if -1 == strs.FindIn(".tmp", extensions) {
		fmt.Println("*** Warning ***: '.tmp' not in 'ignoreExtensions'; this may cause problems")
	}
	return extensions
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
