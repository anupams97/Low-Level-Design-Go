package filesystem

import (
	"fmt"
	"strings"
)

type File struct {
	Name     string
	FileType string
	Contents string
}

func (f *File) PrintName(indent int) {
	fmt.Printf("%sFile: %s.%s\n", strings.Repeat("  ", indent), f.Name, f.FileType)
}

func (f *File) SetContent(content string) {
	f.Contents = content
}

func (f *File) GetContent() string {
	return f.Contents
}

func NewFile(name, fileType, contents string) *File {
	return &File{Name: name, FileType: fileType, Contents: contents}
}
