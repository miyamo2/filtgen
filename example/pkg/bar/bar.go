//go:generate filtgen generate --source=$GOFILE
package bar

type Bar struct {
	String string `filtgen:"*"`
}
