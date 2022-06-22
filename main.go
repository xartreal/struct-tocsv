package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//in: a, out: []string
func structtoarr(a interface{}) []string {
	structType := reflect.ValueOf(a) // could be any underlying type

	out := []string{}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		str := ""
		switch field.Type() {
		case reflect.TypeOf(""):
			str = field.Interface().(string)
		case reflect.TypeOf(1):
			x := field.Interface().(int)
			str = strconv.Itoa(x)
		case reflect.TypeOf(1.1):
			str = fmt.Sprintf("%.2f", field.Interface().(float64)) // or universal "%f"
		default:
			continue
		}
		out = append(out, str)
	}
	return out
}

//in: s, out: a
func arrtostruct(s []string, a interface{}) int { //call: s,&a
	ValuePtr := reflect.ValueOf(a)
	if ValuePtr.Kind() != reflect.Ptr {
		return 0
	}
	Value := ValuePtr.Elem()
	cmax := len(s)
	if cmax > Value.NumField() {
		cmax = Value.NumField()
	}

	for i := 0; i < cmax; i++ {
		field := Value.Field(i)
		switch field.Type() {
		case reflect.TypeOf(""):
			field.SetString(s[i])
		case reflect.TypeOf(1):
			x, _ := strconv.ParseInt(s[i], 10, 64)
			field.SetInt(x)
		case reflect.TypeOf(1.1):
			x, _ := strconv.ParseFloat(s[i], 64)
			field.SetFloat(x)
		default:
			continue
		}
	}
	return cmax
}

func mkcsvrec(in []string) string {
	return strings.Join(in, ";")
}

func splitcsv(in string) []string {
	return strings.Split(in, ";")
}

type TRec struct {
	ID     string
	Name   string
	Price  float64
	Amount int
}

func main() {
	fmt.Printf("hello!\n")

	// struct -> csv
	data := []TRec{
		TRec{"1", "Втулка С1", 65.7, 100},
		TRec{"2", "Гайка-барашек", 10, 70},
		TRec{"3", "Кружка", 110.12, 50},
	}

	out := ""

	for _, item := range data {
		//		fmt.Printf("%+v\n", item)
		arr := structtoarr(item)
		out += mkcsvrec(arr) + "\n"
	}

	ioutil.WriteFile("result.csv", []byte(out), 0644)

	// ------------------------------------------------

	// csv -> struct
	csvdata := []string{
		"1;Втулка С1;65.70;100",
		"2;Гайка-барашек;10;70",
		"3;Кружка;110.12;50",
	}

	for _, item := range csvdata {
		rec := TRec{}
		arr := splitcsv(item)
		arrtostruct(arr, &rec)
		fmt.Printf("%+v\n", rec)
	}

}
