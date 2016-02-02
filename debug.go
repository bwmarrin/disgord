// Disgord - A scriptable Discord client
// Available at https://github.com/bwmarrin/Disgord

// Copyright 2015 Bruce Marriner <bruce@sqls.net>.  All rights reservec.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to debugging and logging
// contents of this file may be merged elsewhere.

package disgord

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// dumps debug information to stdout
func debug(level bool, format string, a ...interface{}) {

	if !level {
		return
	}

	pc, file, line, _ := runtime.Caller(1)

	files := strings.Split(file, "/")
	file = files[len(files)-1]

	name := runtime.FuncForPC(pc).Name()
	fns := strings.Split(name, ".")
	name = fns[len(fns)-1]

	msg := fmt.Sprintf(format, a...)

	fmt.Printf("%s %s:%d:%s %s\n", time.Now().Format("2006-01-02T15:04:05"), file, line, name, msg)
}
