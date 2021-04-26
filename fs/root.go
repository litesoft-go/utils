package fs

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const PathSeparator = string(os.PathSeparator)

type FS struct {
	root string
}

func NewFS(root string) (*FS, error) {
	err := RequireDirectory(root)
	if err != nil {
		return nil, err
	}
	return &FS{root: root}, nil
}

func (in *FS) RootDirectory() Directory {
	return Directory{fs: in, path: &Path{}}
}

func (in *FS) String() string {
	return "FS: " + in.root
}

type Directory struct {
	fs   *FS
	path *Path
}

func (in Directory) String() string {
	return "(" + in.fs.root + "):" + in.path.String()
}

func (in Directory) AsPathString() string {
	return in.fs.root + in.path.String()
}

func (in Directory) Path() string {
	return in.path.String()
}

func (in Directory) IsEquivalent(ni Directory) bool {
	return in.Path() == ni.Path()
}

func (in Directory) IsLess(ni Directory) bool {
	return in.Path() < ni.Path()
}

func (in *Directory) AsOrderable() *string {
	if in == nil {
		return nil
	}
	path := in.Path()
	return &path
}

func FirstDirectory(dirs []Directory) *Directory {
	if len(dirs) == 0 {
		return nil
	}
	return &dirs[0]
}

type File struct {
	fs   *FS
	path *Path
	name string
}

func (in File) String() string {
	return "(" + in.fs.root + "):" + in.path.String() + PathSeparator + in.name
}

func (in File) AsPathString() string {
	return in.AsParentPathString() + "/" + in.name
}

func (in File) Path() string {
	return in.path.String() + "/" + in.name
}

func (in File) AsParentPathString() string {
	return in.fs.root + in.path.String()
}

func (in File) Name() string {
	return in.name
}

func (in File) IsEquivalent(ni File) bool {
	return in.Path() == ni.Path()
}

func (in File) IsLess(ni File) bool {
	return in.Path() < ni.Path()
}

func (in *File) AsOrderable() *string {
	if in == nil {
		return nil
	}
	path := in.Path()
	return &path
}

func FirstFile(files []File) *File {
	if len(files) == 0 {
		return nil
	}
	return &files[0]
}

type Other struct {
	fs   *FS
	path *Path
	info os.FileInfo
}

func (in Other) String() string {
	return "(" + in.fs.root + "):" + in.path.String() + PathSeparator + in.info.Name() + " (" + in.info.Mode().String() + ")"
}

func (in Other) AsPathString() string {
	return in.fs.root + in.path.String() + PathSeparator + in.info.Name()
}

func (in Other) Info() os.FileInfo {
	return in.info
}

func (in Directory) Children() (dirs []Directory, files []File, others []Other, err error) {
	var infos []os.FileInfo
	infos, err = ioutil.ReadDir(in.AsPathString())
	if err == nil {
		for _, info := range infos {
			if info.IsDir() {
				dirs = append(dirs, in.AsDir(info.Name()))
			} else if info.Mode().IsRegular() { // File
				files = append(files, in.AsFile(info.Name()))
			} else { // Other
				others = append(others, in.AsOther(info))
			}
		}
	}
	//fmt.Println("From: ", in)
	//fmt.Println("    dirs: ", dirs)
	//fmt.Println("    files: ", files)
	//fmt.Println("    others: ", others)
	return
}

func (in Directory) AsDir(name string) Directory {
	return Directory{fs: in.fs, path: in.path.Plus(name)}
}

func (in Directory) AsFile(name string) File {
	return File{fs: in.fs, path: in.path, name: name}
}

func (in Directory) AsOther(info os.FileInfo) Other {
	return Other{fs: in.fs, path: in.path, info: info}
}

func (in File) RelativeTo(fs *FS) File {
	return File{fs: fs, path: in.path, name: in.name}
}

func (in File) WithName(name string) File {
	return File{fs: in.fs, path: in.path, name: name}
}

func (in File) Exists() (exists bool, err error) {
	_, _, err = in.check() // path always returned
	if err == nil {
		exists = true
	} else if os.IsNotExist(err) {
		err = nil
	}
	return
}

func (in File) CopyTo(out File) error {
	srcPath, _, err := in.check() // path always returned
	if err != nil {
		return err
	}
	dstPath, _, err := out.check() // path always returned
	if (err != nil) && !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(out.AsParentPathString(), 0755)
	if err != nil {
		return fmt.Errorf("unable to create '%s': %w", dstPath, err)
	}
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer CloseQuietly(src) // Don;t care about error!

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, src)
	return EnsureClosed(dst, err)
}

func (in File) Rename(out File) error {
	srcPath, _, err := in.check() // path always returned
	if err != nil {
		return err
	}
	dstPath, _, err := out.check() // path always returned
	if err == nil {
		err = os.Remove(dstPath)
	} else if !os.IsNotExist(err) {
		return err
	}
	return os.Rename(srcPath, dstPath)
}

func (in File) Delete() error {
	return os.Remove(in.AsPathString())
}

func (in File) IsEmpty() (empty bool, err error) {
	var size int64
	size, err = in.Size()
	if err == nil {
		empty = size == 0
	}
	return
}

func (in File) Size() (size int64, err error) {
	var info os.FileInfo
	_, info, err = in.check()
	if err == nil {
		size = info.Size()
	}
	return
}

func (in File) check() (path string, info os.FileInfo, err error) {
	path = in.AsPathString()
	info, err = os.Stat(path)
	if err == nil {
		if !info.Mode().IsRegular() {
			err = fmt.Errorf("not a Regular file: %s :: %s", path, info.Mode().String())
		}
	}
	return
}
