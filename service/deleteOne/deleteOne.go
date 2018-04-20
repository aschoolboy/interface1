package deleteOne

import (
	"net/http"
	"fmt"
	"interface1/db/sqlBean"
	"interface1/control"
	"interface1/db/toSQLString"
)

func DeleteOne(w http.ResponseWriter, r *http.Request) {

	json := control.ToDeleteParam(r)
	fmt.Printf("json=%+v\n1",json)
	deleteString := toSQLString.ToDeleteString(json)
	fmt.Printf("sql=%v\n",deleteString)
	sqlBean.Remove(deleteString)

}
