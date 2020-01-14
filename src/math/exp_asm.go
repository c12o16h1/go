// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64

package math

import "github.com/c12o16h1/go/src/internal/cpu"

var useFMA = cpu.X86.HasAVX && cpu.X86.HasFMA
