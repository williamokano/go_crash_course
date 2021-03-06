package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Order is struct for orders.
type Order struct {
	OrderID    int `json:"order_id" column:"order_id"`
	CustomerID int `json:"customer_id" column:"customer_id"`
}

// Employee is whatever you want
type Employee struct {
	Name    string `json:"name"`
	ID      int    `json:"id" column:"employee_id"`
	Address string `json:"address"`
	Salary  int    `json:"salary" column:"annual_salary"`
	Country string `json:"country"`
}

// Clear is better than clever... Reflection is not clear!
func main() {
	order := Order{
		OrderID:    321321,
		CustomerID: 987645,
	}

	employee := Employee{
		Name:    "Jorge Amado",
		ID:      95195195,
		Address: "Rua dos pinhais, 470",
		Salary:  120000,
		Country: "Brasil",
	}
	orderInsertQuery, _ := createGenericInsertQuery(order)
	employeeInsertQuery, _ := createGenericInsertQuery(employee)

	fmt.Println(orderInsertQuery)
	fmt.Println(employeeInsertQuery)
}

func createGenericInsertQuery(data interface{}) (string, error) {
	typeOfData := reflect.TypeOf(data)
	kindOfData := typeOfData.Kind()

	if kindOfData == reflect.Struct {
		valueOfData := reflect.ValueOf(data)
		fieldsOfDataStruct := make(map[string]string)
		indirectValueOfData := reflect.Indirect(valueOfData)
		for i := 0; i < valueOfData.NumField(); i++ {

			var fieldName string
			if fieldNameTag, ok := indirectValueOfData.Type().Field(i).Tag.Lookup("column"); ok {
				fieldName = fieldNameTag
			} else {
				fieldName = indirectValueOfData.Type().Field(i).Name
			}

			if valueOfData.Field(i).Type().Kind() == reflect.String {
				fieldsOfDataStruct[fieldName] = "'" + valueOfData.Field(i).String() + "'"
			} else {
				fieldsOfDataStruct[fieldName] = strconv.FormatInt(valueOfData.Field(i).Int(), 10)
			}
		}

		fields := ""
		values := ""

		for field, value := range fieldsOfDataStruct {
			fields += field + ", "
			values += value + ", "
		}

		fields = strings.TrimSuffix(fields, ", ")
		values = strings.TrimSuffix(values, ", ")

		if len(fields) == 0 || len(values) == 0 {
			return "", errors.New("Could not generate fields and values for the query")
		}

		return fmt.Sprintf("insert into %s (%s) values (%s)", typeOfData.Name(), fields, values), nil
	}

	return "", errors.New("Type is not a struct")
}
