// AUTOGENERATED FILE
// +build !codeanalysis
// +build !syz_target syz_target,syz_os_trusty,syz_arch_arm

package gen

import . "github.com/google/syzkaller/prog"
import . "github.com/google/syzkaller/sys/trusty"

func init() {
	RegisterTarget(&Target{OS: "trusty", Arch: "arm", Revision: revision_arm, PtrSize: 4, PageSize: 4096, NumPages: 4096, DataOffset: 536870912, Syscalls: syscalls_arm, Resources: resources_arm, Structs: structDescs_arm, Consts: consts_arm}, InitTarget)
}

var resources_arm = []*ResourceDesc(nil)

var structDescs_arm = []*KeyedStruct{
	{StructKey{Name: "dma_pmem"}, &StructDesc{TypeCommon: TypeCommon{TypeName: "dma_pmem", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4}}},
	}}},
	{StructKey{Name: "ipc_msg"}, &StructDesc{TypeCommon: TypeCommon{TypeName: "ipc_msg", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4}}},
	}}},
	{StructKey{"ipc_msg", 1}, &StructDesc{TypeCommon: TypeCommon{TypeName: "ipc_msg", TypeSize: 4, ArgDir: 1}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4, ArgDir: 1}}},
	}}},
	{StructKey{Name: "ipc_msg_info"}, &StructDesc{TypeCommon: TypeCommon{TypeName: "ipc_msg_info", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4}}},
	}}},
	{StructKey{Name: "uevent"}, &StructDesc{TypeCommon: TypeCommon{TypeName: "uevent", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4}}},
	}}},
	{StructKey{"uevent", 1}, &StructDesc{TypeCommon: TypeCommon{TypeName: "uevent", TypeSize: 4, ArgDir: 1}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4, ArgDir: 1}}},
	}}},
	{StructKey{"uuid", 1}, &StructDesc{TypeCommon: TypeCommon{TypeName: "uuid", TypeSize: 4, ArgDir: 1}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "todo", TypeSize: 4, ArgDir: 1}}},
	}}},
}

var syscalls_arm = []*Syscall{
	{NR: 18, Name: "accept", CallName: "accept", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle_id", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "peer_uuid", TypeSize: 4}, &StructType{Key: StructKey{"uuid", 1}}},
	}},
	{NR: 2, Name: "brk", CallName: "brk", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "brk", TypeSize: 4}}},
	}},
	{NR: 19, Name: "close", CallName: "close", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle_id", TypeSize: 4}}},
	}},
	{NR: 17, Name: "connect", CallName: "connect", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "string", IsVarlen: true}, Kind: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
	}},
	{NR: 3, Name: "exit_etc", CallName: "exit_etc", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "status", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
	}},
	{NR: 11, Name: "finish_dma", CallName: "finish_dma", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "uaddr", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"uaddr"}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
	}},
	{NR: 32, Name: "get_msg", CallName: "get_msg", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "msg_info", TypeSize: 4}, &StructType{Key: StructKey{Name: "ipc_msg_info"}}},
	}},
	{NR: 7, Name: "gettime", CallName: "gettime", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "clock_id", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "time", TypeSize: 4}, &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", TypeSize: 8, ArgDir: 1}}}},
	}},
	{NR: 21, Name: "handle_set_create", CallName: "handle_set_create"},
	{NR: 22, Name: "handle_set_ctrl", CallName: "handle_set_ctrl", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "cmd", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "evt", TypeSize: 4}, &StructType{Key: StructKey{Name: "uevent"}}},
	}},
	{NR: 5, Name: "ioctl", CallName: "ioctl", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "fd", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "req", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "buf", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "array", IsVarlen: true}}},
	}},
	{NR: 8, Name: "mmap", CallName: "mmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "uaddr", TypeSize: 4}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"uaddr"}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
	}},
	{NR: 9, Name: "munmap", CallName: "munmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "uaddr", TypeSize: 4}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"uaddr"}},
	}},
	{NR: 6, Name: "nanosleep", CallName: "nanosleep", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "clock_id", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "sleep_time", TypeSize: 8}}},
	}},
	{NR: 16, Name: "port_create", CallName: "port_create", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "path", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "string", IsVarlen: true}, Kind: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "num_recv_bufs", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "recv_buf_size", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
	}},
	{NR: 10, Name: "prepare_dma", CallName: "prepare_dma", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "uaddr", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"uaddr"}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "flags", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "pmem", TypeSize: 4}, &StructType{Key: StructKey{Name: "dma_pmem"}}},
	}},
	{NR: 34, Name: "put_msg", CallName: "put_msg", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "msg_id", TypeSize: 4}}},
	}},
	{NR: 4, Name: "read", CallName: "read", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "fd", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "msg", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "array", ArgDir: 1, IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"msg"}},
	}},
	{NR: 33, Name: "read_msg", CallName: "read_msg", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "msg_id", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "offset", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "msg", TypeSize: 4}, &StructType{Key: StructKey{"ipc_msg", 1}}},
	}},
	{NR: 35, Name: "send_msg", CallName: "send_msg", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "msg", TypeSize: 4}, &StructType{Key: StructKey{Name: "ipc_msg"}}},
	}},
	{NR: 20, Name: "set_cookie", CallName: "set_cookie", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "cookie", TypeSize: 4}}},
	}},
	{NR: 24, Name: "wait", CallName: "wait", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "handle_id", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "event", TypeSize: 4}, &StructType{Key: StructKey{Name: "uevent"}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "timeout_msecs", TypeSize: 4}}},
	}},
	{NR: 25, Name: "wait_any", CallName: "wait_any", Args: []Type{
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "event", TypeSize: 4}, &StructType{Key: StructKey{"uevent", 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "timeout_msecs", TypeSize: 4}}},
	}},
	{NR: 1, Name: "write", CallName: "write", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "fd", TypeSize: 4}}},
		&PtrType{TypeCommon{TypeName: "ptr", FldName: "msg", TypeSize: 4}, &BufferType{TypeCommon: TypeCommon{TypeName: "array", IsVarlen: true}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "size", TypeSize: 4}}, Path: []string{"msg"}},
	}},
}

var consts_arm = []ConstValue{
	{"__NR_accept", 18},
	{"__NR_brk", 2},
	{"__NR_close", 19},
	{"__NR_connect", 17},
	{"__NR_exit_etc", 3},
	{"__NR_finish_dma", 11},
	{"__NR_get_msg", 32},
	{"__NR_gettime", 7},
	{"__NR_handle_set_create", 21},
	{"__NR_handle_set_ctrl", 22},
	{"__NR_ioctl", 5},
	{"__NR_mmap", 8},
	{"__NR_munmap", 9},
	{"__NR_nanosleep", 6},
	{"__NR_port_create", 16},
	{"__NR_prepare_dma", 10},
	{"__NR_put_msg", 34},
	{"__NR_read", 4},
	{"__NR_read_msg", 33},
	{"__NR_send_msg", 35},
	{"__NR_set_cookie", 20},
	{"__NR_wait", 24},
	{"__NR_wait_any", 25},
	{"__NR_write", 1},
}

const revision_arm = "ab417848f2bf8ecdc807d9ee618bdf7ef646eeaa"
