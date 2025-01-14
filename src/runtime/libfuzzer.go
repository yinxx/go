// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build libfuzzer

package runtime

import "unsafe"

func libfuzzerCallWithTwoByteBuffers(fn, start, end *byte)
func libfuzzerCall4(fn *byte, fakePC uintptr, s1, s2 unsafe.Pointer, result uintptr)
func libfuzzerCall(fn *byte, arg0, arg1 uintptr)

func libfuzzerTraceCmp1(arg0, arg1 uint8) {
	libfuzzerCall(&__sanitizer_cov_trace_cmp1, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceCmp2(arg0, arg1 uint16) {
	libfuzzerCall(&__sanitizer_cov_trace_cmp2, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceCmp4(arg0, arg1 uint32) {
	libfuzzerCall(&__sanitizer_cov_trace_cmp4, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceCmp8(arg0, arg1 uint64) {
	libfuzzerCall(&__sanitizer_cov_trace_cmp8, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceConstCmp1(arg0, arg1 uint8) {
	libfuzzerCall(&__sanitizer_cov_trace_const_cmp1, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceConstCmp2(arg0, arg1 uint16) {
	libfuzzerCall(&__sanitizer_cov_trace_const_cmp2, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceConstCmp4(arg0, arg1 uint32) {
	libfuzzerCall(&__sanitizer_cov_trace_const_cmp4, uintptr(arg0), uintptr(arg1))
}

func libfuzzerTraceConstCmp8(arg0, arg1 uint64) {
	libfuzzerCall(&__sanitizer_cov_trace_const_cmp8, uintptr(arg0), uintptr(arg1))
}

var pcTables []byte

func init() {
	libfuzzerCallWithTwoByteBuffers(&__sanitizer_cov_8bit_counters_init, &__start___sancov_cntrs, &__stop___sancov_cntrs)
	start := unsafe.Pointer(&__start___sancov_cntrs)
	end := unsafe.Pointer(&__stop___sancov_cntrs)

	// PC tables are arrays of ptr-sized integers representing pairs [PC,PCFlags] for every instrumented block.
	// The number of PCs and PCFlags is the same as the number of 8-bit counters. Each PC table entry has
	// the size of two ptr-sized integers. We allocate one more byte than what we actually need so that we can
	// get a pointer representing the end of the PC table array.
	size := (uintptr(end)-uintptr(start))*unsafe.Sizeof(uintptr(0))*2 + 1
	pcTables = make([]byte, size)
	libfuzzerCallWithTwoByteBuffers(&__sanitizer_cov_pcs_init, &pcTables[0], &pcTables[size-1])
}

// We call libFuzzer's __sanitizer_weak_hook_strcmp function
// which takes the following four arguments:
//   1- caller_pc: location of string comparison call site
//   2- s1: first string used in the comparison
//   3- s2: second string used in the comparison
//   4- result: an integer representing the comparison result. Libfuzzer only distinguishes between two cases:
//      - 0 means that the strings are equal and the comparison will be ignored by libfuzzer.
//      - Any other value means that strings are not equal and libfuzzer takes the comparison into consideration.
//      Here, we pass 1 when the strings are not equal.
func libfuzzerHookStrCmp(s1, s2 string, fakePC int) {
	if s1 != s2 {
		libfuzzerCall4(&__sanitizer_weak_hook_strcmp, uintptr(fakePC), cstring(s1), cstring(s2), uintptr(1))
	}
	// if s1 == s2 we could call the hook with a last argument of 0 but this is unnecessary since this case will be then
	// ignored by libfuzzer
}

// This function has now the same implementation as libfuzzerHookStrCmp because we lack better checks
// for case-insensitive string equality in the runtime package.
func libfuzzerHookEqualFold(s1, s2 string, fakePC int) {
	if s1 != s2 {
		libfuzzerCall4(&__sanitizer_weak_hook_strcmp, uintptr(fakePC), cstring(s1), cstring(s2), uintptr(1))
	}
}

//go:linkname __sanitizer_cov_trace_cmp1 __sanitizer_cov_trace_cmp1
//go:cgo_import_static __sanitizer_cov_trace_cmp1
var __sanitizer_cov_trace_cmp1 byte

//go:linkname __sanitizer_cov_trace_cmp2 __sanitizer_cov_trace_cmp2
//go:cgo_import_static __sanitizer_cov_trace_cmp2
var __sanitizer_cov_trace_cmp2 byte

//go:linkname __sanitizer_cov_trace_cmp4 __sanitizer_cov_trace_cmp4
//go:cgo_import_static __sanitizer_cov_trace_cmp4
var __sanitizer_cov_trace_cmp4 byte

//go:linkname __sanitizer_cov_trace_cmp8 __sanitizer_cov_trace_cmp8
//go:cgo_import_static __sanitizer_cov_trace_cmp8
var __sanitizer_cov_trace_cmp8 byte

//go:linkname __sanitizer_cov_trace_const_cmp1 __sanitizer_cov_trace_const_cmp1
//go:cgo_import_static __sanitizer_cov_trace_const_cmp1
var __sanitizer_cov_trace_const_cmp1 byte

//go:linkname __sanitizer_cov_trace_const_cmp2 __sanitizer_cov_trace_const_cmp2
//go:cgo_import_static __sanitizer_cov_trace_const_cmp2
var __sanitizer_cov_trace_const_cmp2 byte

//go:linkname __sanitizer_cov_trace_const_cmp4 __sanitizer_cov_trace_const_cmp4
//go:cgo_import_static __sanitizer_cov_trace_const_cmp4
var __sanitizer_cov_trace_const_cmp4 byte

//go:linkname __sanitizer_cov_trace_const_cmp8 __sanitizer_cov_trace_const_cmp8
//go:cgo_import_static __sanitizer_cov_trace_const_cmp8
var __sanitizer_cov_trace_const_cmp8 byte

//go:linkname __sanitizer_cov_8bit_counters_init __sanitizer_cov_8bit_counters_init
//go:cgo_import_static __sanitizer_cov_8bit_counters_init
var __sanitizer_cov_8bit_counters_init byte

//go:linkname __start___sancov_cntrs __start___sancov_cntrs
//go:cgo_import_static __start___sancov_cntrs
var __start___sancov_cntrs byte

//go:linkname __stop___sancov_cntrs __stop___sancov_cntrs
//go:cgo_import_static __stop___sancov_cntrs
var __stop___sancov_cntrs byte

//go:linkname __sanitizer_cov_pcs_init __sanitizer_cov_pcs_init
//go:cgo_import_static __sanitizer_cov_pcs_init
var __sanitizer_cov_pcs_init byte

//go:linkname __sanitizer_weak_hook_strcmp __sanitizer_weak_hook_strcmp
//go:cgo_import_static __sanitizer_weak_hook_strcmp
var __sanitizer_weak_hook_strcmp byte
