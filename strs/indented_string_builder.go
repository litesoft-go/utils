package strs

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	// standard libs only above!
)

type IndentStringer interface {
	IndentString(prefix string, ib *IndentedBuilder)
}

//noinspection GoUnusedExportedFunction
func NewIndentedBuilder(defaultIndent uint) *IndentedBuilder {
	zDI := int(defaultIndent)
	return &IndentedBuilder{defaultIndent: zDI, sb: &strings.Builder{}}
}

type IndentedBuilder struct {
	defaultIndent int
	sb            *strings.Builder
	curIndent     *indention
}

func (in *IndentedBuilder) String() string {
	return in.sb.String()
}

func (in *IndentedBuilder) DropIndent() {
	in.curIndent = in.curIndent.dropIndent()
}

func (in *IndentedBuilder) AddIndent() {
	in.curIndent = in.curIndent.addIndent(in, nil)
}

func (in *IndentedBuilder) AddIndentOf(newBy uint) {
	in.curIndent = in.curIndent.addIndent(in, &newBy)
}

func (in *IndentedBuilder) AddJSON(v interface{}) error {
	bytes, err := json.Marshal(v)
	if err == nil {
		in.lines(string(bytes))
	}
	return err
}

func (in *IndentedBuilder) AddLine(a ...interface{}) {
	in.lines(fmt.Sprint(a...))
}

func (in *IndentedBuilder) lines(data string) {
	if data == "" {
		in.line(data)
	} else {
		parts := strings.Split(data, "\n")
		for _, part := range parts {
			in.line(part)
		}
	}
}

func (in *IndentedBuilder) line(data string) {
	if data != "" {
		in.curIndent.addIndention(in.sb)
		_, _ = in.sb.WriteString(data)
	}
	_ = in.sb.WriteByte('\n')
}

type indention struct {
	parent *indention
	by     int
}

func (in *indention) addIndent(ib *IndentedBuilder, newBy *uint) *indention {
	by := ib.defaultIndent
	if newBy != nil {
		by = int(*newBy)
	}
	return &indention{parent: in, by: by}
}

func (in *indention) dropIndent() (prev *indention) {
	if in != nil {
		prev = in.parent
	}
	return
}

func (in *indention) getIndent() (indent int) {
	if in != nil {
		indent = in.by + in.parent.getIndent()
	}
	return
}

func (in *indention) addIndention(writer io.ByteWriter) {
	for indent := in.getIndent(); indent > 0; indent-- {
		_ = writer.WriteByte(' ')
	}
}
