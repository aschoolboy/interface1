package bytes2struct

import (
	"fmt"
	"unsafe"
)
type Info struct {
	id                        string
	network_connect_rule_name string
	network_connect_type      string
	network_src               string
	network_src_type          string
	network_dest              string
	network_dest_type         string
	network_desc              string
	create_time               string
	update_time               string
	del_flag                  string
}
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func Bytes2Struct(bytes []byte) (Info){

	var ptestStruct *Info = *(**Info)(unsafe.Pointer(&bytes))
	fmt.Println("ptestStruct is : ", ptestStruct)
	return *ptestStruct
}
