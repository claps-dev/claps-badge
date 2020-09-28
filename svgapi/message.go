package main
//获取汇率及项目信息
import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
)
type Fiat struct {
	Code   string  `gorm:"type:varchar(25);column:code;primary_key"`
	Rate float64 `gorm:"type:decimal;column:rate"`
}

type Data struct {
	Name    string
	Id       string
	Total    string
}
type Employee struct {
	Code    float64
	Data    []Data
	Message string
}
func message() map[string][2] string {
	message := make(map[string][2] string)
	var about [2] string
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("error")
		}
	}()
	url := "http://121.89.171.193:4507/api/projects/"
	resp, _ := http.Get(url)
	var s Employee
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal([]byte(body), &s)
	for _, cur := range s.Data {
		about[0]=cur.Name//项目名
		about[1]=cur.Total//项目总捐赠额
		message[cur.Id] = about
	}
	return message
}
func fiat(code string) float64 {//返回传入货币对于美元的汇率
	db,err:= gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3308)/database")//MySQL数据库地址，端口和密码
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	fs := make([]Fiat, 0)
	if err := db.Select("rate").Where("code = ? ", code).Find(&fs).Error; err != nil {
		log.Fatal(err)
	}
	if len(fs)<1{
		return 0
	}
	return ((fs[0]).Rate)
}
