//go:generate filtgen generate --source=$GOFILE
package main

import (
	s "database/sql"
	"github.com/miyamo2/filtgen-example/pkg/bar"
	"time"
)

type Foo struct {
	String     string        `filtgen:"*"`
	Int        int           `filtgen:"*"`
	Bool       bool          `filtgen:"*"`
	Uint       uint          `filtgen:"*"`
	Uint8      uint8         `filtgen:"*"`
	Uint16     uint16        `filtgen:"*"`
	Uint32     uint32        `filtgen:"*"`
	Uint64     uint64        `filtgen:"*"`
	Uintptr    uintptr       `filtgen:"*"`
	Float32    float32       `filtgen:"*"`
	Float64    float64       `filtgen:"*"`
	Rune       rune          `filtgen:"*"`
	Byte       byte          `filtgen:"*"`
	Int8       int8          `filtgen:"*"`
	Int16      int16         `filtgen:"*"`
	Int32      int32         `filtgen:"*"`
	Int64      int64         `filtgen:"*"`
	Complex64  complex64     `filtgen:"*"`
	Complex128 complex128    `filtgen:"*"`
	Error      error         `filtgen:"*"`
	Bar        bar.Bar       `filtgen:"*"`
	Time       time.Time     `filtgen:"*"`
	NullString s.NullString  `filtgen:"*"`
	Embed      `filtgen:"*"` // nothing is being generated.
}

type Baz struct {
	String     string       `filtgen:"is,isnt,as,asnt"` // nothing is being generated.
	Int        int          // nothing is being generated.
	Bool       bool         `filtgen:"gt,lt,ge,le"`       // nothing is being generated.
	Complex64  complex64    `filtgen:"eq,ne,gt,lt,ge,le"` // nothing is being generated.
	Complex128 complex128   `filtgen:"eq,ne,gt,lt,ge,le"` // nothing is being generated.
	Error      error        `filtgen:"eq,ne,gt,lt,ge,le"` // nothing is being generated.
	Time       time.Time    `filtgen:"eq,ne,gt,lt,ge,le"` // nothing is being generated.
	NullString s.NullString `filtgen:"eq,ne,gt,lt,ge,le"` // nothing is being generated.
}
