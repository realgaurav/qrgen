// Copyright (c) 2022 Gaurav Khatri. All rights reserved.
//
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package qr

import (
	"bytes"
	"testing"
)

func TestEscape(t *testing.T) {
	for _, sa := range [][]string{
		{`tested`, `tested`},
		{`test\ed`, `test\\ed`},
		{`test;ed`, `test\;ed`},
		{`test,ed`, `test\,ed`},
		{`test"ed`, `test\"ed`},
		{`test:ed`, `test\:ed`},
	} {
		tbi := []byte(sa[0])
		tbo := []byte(sa[1])

		if !bytes.Equal(tbo, escape(tbi)) {
			t.Fatalf(`escape(%q) want %s got %q`, tbi, tbo, escape(tbi))
		}
	}
}
