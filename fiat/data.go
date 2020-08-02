package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/robfig/cron"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "strconv"
)
//定义api接口数据结构和序列化json字段
type Data struct {
	Code    string
	Rate    decimal.Decimal
}
type Employee struct {
	Data    []Data
}
type Fiat struct {
	Code   string  `gorm:"type:varchar(25);column:code;primary_key"`
	Rate decimal.Decimal `gorm:"type:decimal;column:rate"`
}
func data(){
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("error")
		}
	}()
	url := "https://echo.yiplee.com/fiats"
	resp, _ := http.Get(url)
	var s Employee
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))
	json.Unmarshal([]byte(body), &s)
	data := make(map[string]decimal.Decimal)
	for _, cur := range s.Data {
		data[cur.Code] = cur.Rate
	}
	db,err:= gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3308)/database")//MySQL地址端口和密码
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.Delete(&Fiat{})
	for c,b := range data{
		test := &Fiat{
			Code:c,
			Rate:b,
		}
		db.Create(test)
		fmt.Println(c,b)
	}
}
