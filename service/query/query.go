package query

import (
	"net/http"
	"interface1/control"
	"fmt"
	"interface1/db/toSQLString"
	"interface1/db/sqlBean"
	"interface1/util/change/struct2bytes"
	"interface1/util/JsonResponse"
)

func Query(w http.ResponseWriter, r *http.Request) {
	json := control.ToSingleSelectParam(r)
	fmt.Printf("json=%+v\n1",json)
	queryString := toSQLString.ToSingleSelectString(json)
	fmt.Printf("sql=%v\n",queryString)
	selectResult := sqlBean.QueryInfos(queryString)
	bytes := struct2bytes.Struct2Bytes(selectResult)
	JsonResponse.JsonResponse(bytes, w)
}
