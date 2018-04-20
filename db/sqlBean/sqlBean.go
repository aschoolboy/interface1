package sqlBean

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"interface1/util/checkErr"
	"interface1/public/singleSelect"
)

//type Info struct {
//	Id                        string
//	Network_connect_rule_name string
//	Network_connect_type      string
//	Network_src               string
//	Network_src_type          string
//	Network_dest              string
//	Network_dest_type         string
//	Network_desc              string
//	Create_time               string
//	Update_time               string
//	Del_flag                  string
//}

func ConnDB() *sql.DB {
	db, err := sql.Open("mysql", "root:20111412e@tcp(172.105.204.252:3306)/PXTest?charset=utf8")
	checkErr.CheckErr(err, 1)
	return db
}

//插入demo
func Insert(sqlWord string) {
	conn :=ConnDB()
	stmt, err := conn.Prepare(sqlWord)
	checkErr.CheckErr(err, 2)
	res, err := stmt.Exec()
	checkErr.CheckErr(err, 3)
	id, err := res.LastInsertId()
	checkErr.CheckErr(err, 4)
	fmt.Println(id)
}

//查询Infos，将结果放入Info结构体数组
func QueryInfos(sqlWord string) []singleSelect.Info {
	var Result []singleSelect.Info
	fmt.Println("sql=" + sqlWord)
	conn :=ConnDB()
	defer conn.Close()
	rows, err := conn.Query(sqlWord)
	if err != nil {
		fmt.Println("mysql query error", err.Error())
	}
	defer rows.Close()
	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	info := singleSelect.Info{}
	for rows.Next() {
		_ = rows.Scan(scans...)
		for i, col := range values {
			switch expr := columns[i]; expr {
			case "id":
				info.Id = string(col)
			case "network_connect_rule_name":
				info.Network_connect_rule_name = string(col)
			case "network_connect_type":
				info.Network_connect_type = string(col)
			case "network_src":
				info.Network_src = string(col)
			case "network_src_type":
				info.Network_src_type = string(col)
			case "network_dest":
				info.Network_dest = string(col)
			case "network_dest_type":
				info.Network_dest_type = string(col)
			case "network_desc":
				info.Network_desc = string(col)
			case "nreate_time":
				info.Create_time = string(col)
			case "npdate_time":
				info.Update_time = string(col)
			case "del_flag":
				info.Del_flag = string(col)
			}
		}

		Result = append(Result, info)

	}
	return Result
}

//查询num
func QueryNum(sqlWord string) string {
	fmt.Println("sql=" + sqlWord)
	conn :=ConnDB()

	defer conn.Close()
	rows, err := conn.Query(sqlWord)
	if err != nil {
		fmt.Println("mysql query error", err.Error())
	}
	defer rows.Close()
	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scans...)
		for _, col := range values {
			num := string(col)
			return num
		}
	}
	return "-1"
}

//更新数据
//func Update() {
//	db :=ConnDB()
//
//	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
//	checkErr.CheckErr(err,1)
//	res, err := stmt.Exec(21, 2, 1)
//	checkErr.CheckErr(err,2)
//	num, err := res.RowsAffected()
//	checkErr.CheckErr(err,3)
//	fmt.Println(num)
//}

//删除数据
func Remove(str string) {
	conn :=ConnDB()

	stmt, err := conn.Prepare(str)
	checkErr.CheckErr(err, 2)
	res, err := stmt.Exec()
	checkErr.CheckErr(err, 3)
	num, err := res.RowsAffected()
	checkErr.CheckErr(err, 4)
	fmt.Println(num)
}

