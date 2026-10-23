// SPDX-FileCopyrightText: 2025-2026 caixw
//
// SPDX-License-Identifier: MIT

package serialize

import (
	"encoding/json/v2"
	"encoding/xml"
	"io/fs"
	"testing"

	"github.com/issue9/assert/v5"
	"golang.org/x/text/language"

	"github.com/issue9/messages"
)

func TestSaveFile(t *testing.T) {
	a := assert.New(t, false)

	f := &messages.File{
		Languages: []language.Tag{language.SimplifiedChinese},
		Messages: []messages.Message{
			{
				Key: "k1",
				Message: messages.Text{
					Msg: "m1",
				},
			},
			{
				Key: "k2",
				Message: messages.Text{
					Msg: "m2",
				},
			},
		},
	}

	a.NotError(SaveFile(f, "./testdata/json.out", func(v any) ([]byte, error) { return json.Marshal(v) }, fs.ModePerm))
	a.NotError(SaveFile(f, "./testdata/xml.out", xml.Marshal, fs.ModePerm))
}
