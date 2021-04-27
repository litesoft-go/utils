package comparator

import (
	"fmt"
	"github.com/litesoft-go/utils/fs"
	picker "github.com/litesoft-go/utils/twoorderedslicepicker"
	"io"
	"os"
	"sort"
	"strings"
)

type Comparator struct {
	ignoreExtensions []string
	ignoreDirs       []string
	dir1FS           *fs.FS
	dir2FS           *fs.FS
}

func NewComparator(dir1FS, dir2FS *fs.FS, ignoreDirs []string, ignoreExtensions []string) *Comparator {
	return &Comparator{dir1FS: dir1FS, dir2FS: dir2FS, ignoreDirs: ignoreDirs, ignoreExtensions: ignoreExtensions}
}

func (in *Comparator) CompareFiles() {
	fmt.Println()
	in.compareDirs(in.dir1FS.RootDirectory(), in.dir2FS.RootDirectory())
}

func (in *Comparator) compareDirs(dir1, dir2 fs.Directory) {
	if in.ignoreDir(dir1) || in.ignoreDir(dir2) {
		return
	}
	dirs1, files1, _, err := dir1.Children()
	if err != nil {
		showDirectoryError(err, dir1)
		return
	}
	dirs2, files2, _, err := dir2.Children()
	if err != nil {
		showDirectoryError(err, dir2)
		return
	}

	files1, files2 = in.compareFiles(files1, files2)

	sort.Slice(dirs1, func(i, j int) bool { return dirs1[i].IsLess(dirs1[j]) })
	sort.Slice(dirs2, func(i, j int) bool { return dirs2[i].IsLess(dirs2[j]) })

	for {
		dir1 := fs.FirstDirectory(dirs1)
		dir2 := fs.FirstDirectory(dirs2)
		pick := picker.Which(dir1, dir2)
		// fmt.Println("***  dirs ", pick.Name(), " 1:", dir1, " | ", dir2, ":2")
		switch pick {
		case picker.Done:
			return
		case picker.Left:
			in.showDirectorySingle(*dir1)
			dirs1 = dirs1[1:]
		case picker.Right:
			in.showDirectorySingle(*dir2)
			dirs2 = dirs2[1:]
		case picker.Equivalent:
			in.compareDirs(*dir1, *dir2)
			dirs1 = dirs1[1:]
			dirs2 = dirs2[1:]
		}
	}
}

func (in *Comparator) compareFiles(files1, files2 []fs.File) (f1, f2 []fs.File) {
	sort.Slice(files1, func(i, j int) bool { return files1[i].IsLess(files1[j]) })
	sort.Slice(files2, func(i, j int) bool { return files2[i].IsLess(files2[j]) })
	for {
		file1 := fs.FirstFile(files1)
		file2 := fs.FirstFile(files2)
		pick := picker.Which(file1, file2)
		// fmt.Println("*** files ", pick.Name(), " 1:", file1, " | ", file2, ":2")
		switch pick {
		case picker.Done:
			return
		case picker.Left:
			in.showFileSingle(*file1)
			files1 = files1[1:]
		case picker.Right:
			in.showFileSingle(*file2)
			files2 = files2[1:]
		case picker.Equivalent:
			in.compareFile(*file1, *file2)
			files1 = files1[1:]
			files2 = files2[1:]
		}
	}
}

const LEADER = "\n        "
const _ERR__ = " (err)"
const _ONLY_ = "(only)"
const _DIFF_ = "(DIFF)"

func showFileError(err error, file fs.File) {
	fmt.Println(LEADER, "file", _ERR__, ": ", file, LEADER, "  of: ", err)
}

func showDirectoryError(err error, dir fs.Directory) {
	fmt.Println(LEADER, " dir", _ERR__, ": ", dir, LEADER, "  of: ", err)
}

func showFileOnly(file fs.File) {
	fmt.Println(LEADER, "file", _ONLY_, ": ", file)
}

func showDirectoryOnly(dir fs.Directory) {
	fmt.Println(LEADER, " dir", _ONLY_, ": ", dir)
}

func showFileDiff(file1, file2 fs.File) {
	fmt.Println(LEADER, "file", _DIFF_, ":",
		LEADER, "   1:", file1,
		LEADER, "   2:", file2)
}

func (in *Comparator) showDirectorySingle(dir fs.Directory) {
	if !in.ignoreDir(dir) {
		showDirectoryOnly(dir)
	}
}

func (in *Comparator) showFileSingle(file fs.File) {
	if !in.ignoreFile(file) {
		empty, err := file.IsEmpty()
		if err != nil {
			showFileError(err, file)
		} else if !empty {
			showFileOnly(file)
		}
	}
}

func (in *Comparator) ignoreDir(directory fs.Directory) bool {
	path := directory.AsPathString()
	for _, dir := range in.ignoreDirs {
		if strings.HasSuffix(path, "/"+dir) {
			return true
		}
	}
	return false
}

func (in *Comparator) ignoreFile(file fs.File) bool {
	path := file.AsPathString()
	for _, ext := range in.ignoreExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func (in *Comparator) compareFile(file1, file2 fs.File) {
	if in.ignoreFile(file1) {
		return
	}
	size1, err := file1.Size()
	if err != nil {
		showFileError(err, file1)
		return
	}
	size2, err := file2.Size()
	if err != nil {
		showFileError(err, file2)
		return
	}
	if size1 == size2 {
		if sameContents(file1, file2) {
			return
		}
	}
	showFileDiff(file1, file2)
}

func sameContents(file1, file2 fs.File) bool {
	f1 := NewFileReadBlockStream(file1)
	f2 := NewFileReadBlockStream(file2)
	defer func() {
		fs.CloseQuietly(f1)
		fs.CloseQuietly(f2)
	}()

	count1, bytes1 := f1.read()
	count2, bytes2 := f2.read()
	for (0 <= count1) && (0 <= count2) { // -1 == Error
		// fmt.Println("*** same ", f1.fsFile, " 1:", count1, " | ", count2, ":2 ", f2.fsFile)

		if count2 < count1 {
			count1 = count2
		}
		for i := 0; i < count1; i++ {
			if bytes1[i] != bytes2[i] {
				return false
			}
		}
		f1.consumed(count1)
		f2.consumed(count1)

		count1, bytes1 = f1.read()
		count2, bytes2 = f2.read()
	}
	f1.Check()
	f2.Check()
	return true // Don't show diff
}

type FileReadBlockStream struct {
	fsFile   fs.File
	buf      []byte
	offset   int
	bufBytes int
	osFile   *os.File
	err      error
}

func NewFileReadBlockStream(file fs.File) *FileReadBlockStream {
	f, err := os.Open(file.AsPathString())
	size := 32 * 1024 // See io.copyBuffer
	return &FileReadBlockStream{fsFile: file, buf: make([]byte, size),
		offset: 0, bufBytes: 0, osFile: f, err: err}
}

func (in *FileReadBlockStream) Check() {
	if (in.err != nil) && (in.err != io.EOF) { // expect EOF
		showFileError(in.err, in.fsFile)
	}
}

func (in *FileReadBlockStream) Close() error {
	return fs.EnsureClosed(in.osFile, in.err)
}

// count == -1 -> Error
func (in *FileReadBlockStream) read() (count int, available []byte) {
	for count = in.bufBytes - in.offset; 0 <= count; count = in.bufBytes - in.offset {
		if count != 0 {
			available = in.buf[in.offset:]
			return
		}
		in.offset = 0
		in.bufBytes = -1 // Stop the loop if err != nil
		if in.err == nil {
			in.bufBytes, in.err = in.osFile.Read(in.buf)
		}
	}
	return -1, nil
}

func (in *FileReadBlockStream) consumed(count int) {
	in.offset += count
}
