package struct2bytes

import (
	"encoding/json"
	"interface1/util/checkErr"
)



func Struct2Bytes(info interface{}) []byte{

	bytes,err:=json.Marshal(info)
	checkErr.CheckErr(err,1)

	return bytes
}
