package copier

import (
	"fmt"
	"github.com/litesoft-go/utils/fs"
	"strings"
)

type Copier struct {
	ignoreExtensions []string
	srcFSs           []*fs.FS
	dstFSs           []*fs.FS
}

func NewCopier(srcFSs, dstFSs []*fs.FS, ignoreExtensions []string) *Copier {
	return &Copier{srcFSs: srcFSs, dstFSs: dstFSs, ignoreExtensions: ignoreExtensions}
}

func (in *Copier) CopyFiles() (err error) {
	fmt.Println()
	for _, srcFS := range in.srcFSs {
		fmt.Println(srcFS)
		err = in.copySrcDir(srcFS.RootDirectory())
		if err != nil {
			return
		}
	}
	return
}

func (in *Copier) copySrcDir(srcDir fs.Directory) error {
	dirs, files, others, err := srcDir.Children()
	if err != nil {
		return err
	}
	for _, other := range others {
		fmt.Println("    ", other.AsPathString()+"  :: ("+other.Info().Mode().String()+")")
	}
	others = nil
	for _, file := range files {
		err = in.copySrcFile(file)
		if err != nil {
			fmt.Println("\n        ", err)
		}
	}
	files = nil
	for _, dir := range dirs {
		err = in.copySrcDir(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

func (in *Copier) copySrcFile(srcFile fs.File) error {
	path := srcFile.AsPathString()
	for _, ext := range in.ignoreExtensions {
		if strings.HasSuffix(path, ext) {
			return nil
		}
	}
	printed := false
	name := srcFile.Name()

	tmpName := convertToTempName(name)
	for i, dstFS := range in.dstFSs {
		dstFile := srcFile.RelativeTo(dstFS)
		exists, err := dstFile.Exists()
		if err != nil {
			return err
		}
		if !exists {
			if !printed {
				fmt.Print("    ", path)
				printed = true
			}
			err = srcFile.CopyTo(dstFile.WithName(tmpName))
			if err != nil {
				return err
			}
			fmt.Print(" || cd", i)
			err := dstFile.WithName(tmpName).Rename(dstFile)
			if err != nil {
				return err
			}
			fmt.Print(" || td", i)
		}
	}
	if printed {
		fmt.Println()
	}
	return nil
}

func convertToTempName(name string) string {
	return "_try_" + name
}
