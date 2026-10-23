module github.com/issue9/messages/extract/testdata

require github.com/issue9/localeutil v0.34.0

// NOTE: 需要保证与根目录中 go.mod 的 text 具有相同的版本，否则测试会失败！
require golang.org/x/text v0.43.0 // indirect

go 1.27.0
