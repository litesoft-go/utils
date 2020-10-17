package mover

import (
	"fmt"
	"github.com/litesoft-go/utils/fs"
	"strings"
)

type Mover struct {
	skippedFiles     []string
	ignoreExtensions []string
	srcFSs           []*fs.FS
	dstFSs           []*fs.FS
}

func NewMover(srcFSs, dstFSs []*fs.FS, ignoreExtensions []string) *Mover {
	return &Mover{srcFSs: srcFSs, dstFSs: dstFSs, ignoreExtensions: ignoreExtensions}
}

func (in *Mover) moveSrcFile(srcFile fs.File) error {
	path := srcFile.AsPathString()
	for _, ext := range in.ignoreExtensions {
		if strings.HasSuffix(path, ext) {
			return nil
		}
	}
	fmt.Print("    ", path)
	for _, skip := range in.skippedFiles {
		if path == skip {
			fmt.Println("  :: Skipping")
			return nil
		}
	}
	name := srcFile.Name()
	if len(name) < 6 {
		in.skippedFiles = append(in.skippedFiles, path)
		fmt.Println("  :: Too Short")
		return nil
	}
	tmpName := convertToTempName(name)
	for _, dstFS := range in.dstFSs {
		err := srcFile.CopyTo(srcFile.RelativeTo(dstFS).WithName(tmpName))
		if err != nil {
			return err
		}
	}
	for _, dstFS := range in.dstFSs {
		dstFile := srcFile.RelativeTo(dstFS)
		err := dstFile.WithName(tmpName).Rename(dstFile)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Print("        ", dstFile.AsPathString())
	}
	err := srcFile.Delete()
	fmt.Println()
	return err
}

func convertToTempName(name string) string {
	return "_try" + strings.Repeat("_", len(name)-4)
}

func (in *Mover) MoveFiles() (err error) {
	fmt.Println()
	for _, srcFS := range in.srcFSs {
		fmt.Println(srcFS)
		err = in.moveSrcDir(srcFS.RootDirectory())
		if err != nil {
			return
		}
	}
	return
}

func (in *Mover) moveSrcDir(srcDir fs.Directory) error {
	dirs, files, others, err := srcDir.Children()
	if err != nil {
		return err
	}
	for _, other := range others {
		fmt.Println("    ", other.AsPathString()+"  :: ("+other.Info().Mode().String()+")")
	}
	others = nil
	for _, file := range files {
		err = in.moveSrcFile(file)
		if err != nil {
			return err
		}
	}
	files = nil
	for _, dir := range dirs {
		err = in.moveSrcDir(dir)
		if err != nil {
			return err
		}
	}
	return nil
}
