package main

import (
	"fmt"
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
	srcDirs := config.ValuesForRequired("sources")
	extensions := checkExtensions(config.ValuesForOptional("extensions"))
	args.CheckNoMoreArgs()
	fmt.Println("  config:")
	fmt.Println("         sources (dirs): ", srcDirs)
	fmt.Println("             extensions: ", extensions)

	// srcFSs, err := fileSystemsFor("sources", srcDirs)
	_, err := fileSystemsFor("sources", srcDirs)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	//zMover := mover.NewMover(srcFSs, dstFSs, ignoreExtensions)
	//
	//err = zMover.MoveFiles()
	//for err == nil {
	//	if loopAfterDuration == nil {
	//		os.Exit(0)
	//	}
	//	time.Sleep(*loopAfterDuration)
	//	err = zMover.MoveFiles()
	//}

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
	return extensions
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
