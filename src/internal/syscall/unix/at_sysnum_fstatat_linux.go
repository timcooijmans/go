// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64 || loong64 || riscv64

package unix

import "syscall"

const fstatatTrap uintptr = syscall.SYS_FSTATAT
