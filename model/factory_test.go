package model

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func guozijiudian() *Build {
	bdic := make(map[string]float64)
	bdic["面粉-冀南香5号粉"] = 25000
	bdic["糖"] = 4000
	bdic["盐"] = 250
	bdic["葡萄干"] = 1000
	bdic["苹果丁"] = 500
	bdic["橙皮丁"] = 500
	bdic["复配面包乳化剂"] = 1000
	bdic["面包改良剂"] = 125
	bdic["脱氢乙酸钠"] = 10
	bdic["丙酸钙"] = 60
	bdic["高糖酵母"] = 200
	bdic["无水酥油"] = 3000
	bdic["鸡蛋"] = 800
	bdic["老面"] = 1500

	buildData := &Build{
		Name:         "果子面包-酒店专供",
		Remake:       "",
		Info:         "xxx",
		MaterialUnit: 110,
		Water:        10500,
		Bakingtime:   18,
		Bakingtem:    180,
		BuildData:    bdic,
	}
	return buildData
}

func guozi() *Build {
	bdic := make(map[string]float64)
	bdic["面粉-冀南香5号粉"] = 25000
	bdic["糖"] = 4000
	bdic["盐"] = 250
	bdic["葡萄干"] = 2000
	bdic["苹果丁"] = 1000
	bdic["橙皮丁"] = 1000
	bdic["复配面包乳化剂"] = 1000
	bdic["面包改良剂"] = 125
	bdic["脱氢乙酸钠"] = 10
	bdic["丙酸钙"] = 60
	bdic["高糖酵母"] = 200
	bdic["无水酥油"] = 3000
	bdic["鸡蛋"] = 1500
	bdic["老面"] = 1500

	buildData := &Build{
		Name:         "果子面包",
		Remake:       "",
		Info:         "xxx",
		MaterialUnit: 125,
		Water:        10500,
		Bakingtime:   18,
		Bakingtem:    180,
		BuildData:    bdic,
	}
	return buildData
}

func dalieba() *Build {
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
	bdic["老面"] = 1500

	return &Build{
		Name:         "大列巴",
		Remake:       "夏季大列巴 260g面团 80g葡萄干核桃 8号粉",
		Info:         "xxx",
		Water:        11500,
		MaterialUnit: 320,
		Bakingtime:   32,
		Bakingtem:    160,
		BuildData:    bdic,
	}
}
func TestBuildData(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/xfire?charset=utf8&parseTime=true"))
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	g := guozi()
	dlb := dalieba()
	gzj := guozijiudian()
	db.Create(gzj)

	db.Create(g)
	db.Create(dlb)
	fmt.Println(err)
}
