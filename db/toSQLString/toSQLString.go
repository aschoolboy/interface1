package toSQLString

import (
	"reflect"
	"fmt"
	"interface1/public/singleSelectParam"
	"bytes"
	"strconv"
	"interface1/util/checkErr"
	"interface1/public/batchSelectParam"
	"interface1/public/deleteParam"
)

func ToSingleSelectString(param singleSelectParam.Param) string {
	b := bytes.Buffer{}
	b.WriteString("select * from t_project_network_connect where ")
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)
	first := 0
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		if v.Field(k).String() != "" {
			first++
			if first == 1 {
				b.WriteString(t.Field(k).Name + " = '" + v.Field(k).String() + "' ")
			} else {
				b.WriteString(" and " + t.Field(k).Name + " = '" + v.Field(k).String() + "' ")
			}
		}
	}
	return b.String()
}

func ToBatchSelectString(param batchSelectParam.Param) string {
	b := bytes.Buffer{}
	b.WriteString("select * from t_project_network_connect where ")
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)
	first := 0
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		if v.Field(k).String() != "" {
			first++
			switch expr := t.Field(k).Name; expr {
			case "Page":
				break
			case "Page_size":
				break
			case "Start_time":
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString("create_time >= '" + v.Field(k).String() + "'")
			case "End_time":
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString("create_time <= '" + v.Field(k).String() + "'")
			default:
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString(t.Field(k).Name + " = '" + v.Field(k).String() + "' ")
			}
		}
	}

	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		if v.Field(k).String() != "" {
			if t.Field(k).Name == "Page" {
				iPage, err := strconv.Atoi(v.Field(k).String())
				checkErr.CheckErr(err, 5)
				iPage_size, err := strconv.Atoi(v.Field(k).String())
				checkErr.CheckErr(err, 6)
				start_num := strconv.Itoa((iPage - 1) * iPage_size)
				b.WriteString(" order by id limit " + start_num)
			}
			if t.Field(k).Name == "Page_size" {
				b.WriteString("," + v.Field(k).String())
			}
		}
	}

	return b.String()
}

func ToNumSelectString(param batchSelectParam.Param) string {
	b := bytes.Buffer{}
	b.WriteString("select count(id) from t_project_network_connect where ")
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)
	first := 0
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		if v.Field(k).String() != "" {
			first++
			switch expr := t.Field(k).Name; expr {
			case "Page":
				break
			case "Page_size":
				break
			case "Start_time":
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString("create_time >= '" + v.Field(k).String() + "'")
			case "End_time":
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString("create_time <= '" + v.Field(k).String() + "'")
			default:
				if first != 1 {
					b.WriteString(" and ")
				}
				b.WriteString(t.Field(k).Name + " = '" + v.Field(k).String() + "' ")
			}
		}
	}
	return b.String()
}

func ToDeleteString(param deleteParam.Param) string {
	b := bytes.Buffer{}
	b.WriteString("DELETE FROM t_project_network_connect WHERE id='")
	b.WriteString(param.Id)
	b.WriteString("'")
	return b.String()
}
