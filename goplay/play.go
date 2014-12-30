// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goplay

import (
	"io"
	"net/http"

	"code.google.com/p/go.tools/godoc/static"
)

func init() {
	http.HandleFunc("/playground.js", play)
}

func play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/javascript")
	io.WriteString(w, static.Files["playground.js"])
}
