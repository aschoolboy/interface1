package typeQuery

import (
	"net/http"
	"interface1/db/sqlBean"
	"interface1/util/JsonResponse"
	"interface1/util/change/struct2bytes"
	"interface1/control"
	"fmt"
	"interface1/db/toSQLString"
	"interface1/public/batchSelect"
)

func TypeQuery(w http.ResponseWriter, r *http.Request) {
	json := control.ToBatchSelectParam(r)
	fmt.Printf("json=%+v\n1",json)
	queryString:=toSQLString.ToBatchSelectString(json)
	selectResult := sqlBean.QueryInfos(queryString)
	list:=batchSelect.BatchInfo{}
	list.List=selectResult

	queryNumString := toSQLString.ToNumSelectString(json)
	selectNumResult := sqlBean.QueryNum(queryNumString)

	list.Sum=selectNumResult
	bytes := struct2bytes.Struct2Bytes(list)
	JsonResponse.JsonResponse(bytes, w)
}
