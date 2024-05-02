package model

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestBuildData(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/xfire?charset=utf8&parseTime=true"))
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	bdic := make(map[string]float64)
	bdic["麦麸"] = 750
	bdic["面粉-冀南香8号粉"] = 17500
	bdic["糖"] = 4000
	bdic["盐"] = 300
	bdic["低筋粉"] = 7500
	bdic["可可粉"] = 80
	bdic["核桃仁"] = 3680
	bdic["葡萄干"] = 11040
	bdic["脱氢乙酸钠"] = 12.5
	bdic["丙酸钙"] = 50
	bdic["高糖酵母"] = 200
	bdic["无水酥油"] = 3500
	bdic["鸡蛋"] = 1500
	bdic["捞面"] = 1500

	buildData := Build{
		Name:       "大列巴",
		Remake:     "夏季大列巴 260g面团 80g葡萄干核桃 8号粉",
		Info:       "xxx",
		Water:      11500,
		Bakingtime: 32,
		Bakingtem:  160,
		BuildData:  bdic,
	}
	err = db.Create(&buildData).Error
	fmt.Println(err)
}
