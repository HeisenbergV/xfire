package prompt

import (
	"fmt"
	"strconv"
	"xfire/global"
	"xfire/model"
	"xfire/model/request"
	"xfire/service"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
)

type GoodsPrompt struct {
	args map[string]int //参数名；参数数量
}

func (*GoodsPrompt) Name() string {
	return "GoodsPrompt"
}
func (*GoodsPrompt) Level() int {
	return 2
}

func (*GoodsPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "查看产品"},
		{Text: "info", Description: "产品细节 e.g:info id "},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (*GoodsPrompt) Show() {
	fmt.Println("产品管理")
	fmt.Println("list info")
	fmt.Println("")
}
func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}
func (b *GoodsPrompt) Exec(args []string) bool {
	if args[0] == "q" {
		return true
	}
	cmd := args[0]
	argNum, ok := b.args[cmd]
	if !ok {
		return false
	}
	if len(args) != argNum {
		fmt.Println("参数数量错误")
		return false
	}
	if cmd == "list" {
		data, _, err := service.FactoryService.GetGoodsList(model.Product, request.PageInfo{Page: 0, PageSize: 10000}, "", true)

		if err != nil {
			fmt.Println("查询失败")
			return false
		}
		list := data.([]model.Goods)
		tableData1 := pterm.TableData{
			{"ID", "名称", "编码", "重量/g", "单价"},
		}

		for _, Goods := range list {
			tableData1 = append(tableData1, []string{fmt.Sprintf("%d", Goods.ID), fmt.Sprintf("%s(%s)", Goods.Name, Goods.Brand),
				Goods.Barcode, fmt.Sprintf("%.2f", Goods.Unit),
				fmt.Sprintf("%.2f", Goods.Price),
			})
		}

		pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

	} else if cmd == "info" {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("参数错误")
			return false
		}
		GoodsInfo, err := service.FactoryService.GetGoodsInfo(id)
		fmt.Println(GoodsInfo)
		if err != nil {
			fmt.Println("查询失败")
			return false
		}

		binfo, err := service.FactoryService.GetBuildInfo(GoodsInfo.BuildID)
		if err != nil {
			fmt.Println("查询失败")
			return false
		}

		bulletListItems := []pterm.BulletListItem{
			{
				Level:       0,                                 // Level 0 (top level)
				Text:        fmt.Sprintf("配方(%s)", binfo.Name), // Text to display
				TextStyle:   pterm.NewStyle(pterm.FgRed),       // Text color
				BulletStyle: pterm.NewStyle(pterm.FgRed),       // Bullet color
			},
		}

		cost := 0.0
		sumWeight := 0.0
		for mname, ratio := range binfo.BuildData {
			bulletListItems = append(bulletListItems,
				pterm.BulletListItem{
					Level:       1,                                     // Level 0 (top level)
					Text:        fmt.Sprintf("%s %.2fg", mname, ratio), // Text to display
					TextStyle:   pterm.NewStyle(pterm.FgBlue),          // Text color
					BulletStyle: pterm.NewStyle(pterm.FgRed),           // Bullet color
				})

			materialinfo := global.GetMaterialInfo(mname)
			if materialinfo == nil {
				fmt.Println("异常：没有此材料:", mname)
				return false
			}
			cost += ratio * (materialinfo.Price / materialinfo.Unit)
			fmt.Println(mname, ratio*(materialinfo.Price/materialinfo.Unit))
			// fmt.Printf("%s:%.2fg - %.2f \n", v.Name, useG, useG/(v.Unit*1000)*v.Price)
			sumWeight += ratio
		}
		sumWeight += binfo.Water
		num := int(sumWeight / 80)
		costSum := float64(num)*GoodsInfo.Price - cost
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("%s\n"), binfo.Name)
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("一袋面原料成本约:%.2f \n"), cost)
		pterm.DefaultBasicText.Printf("大约做 %d 个 单品价格:%.2f \n", num, GoodsInfo.Price)
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("一袋面利润约（不算人工水电费等）:%.2f元; 单个利润约: %.2f元 \n"), costSum, costSum/float64(num))

		pterm.DefaultBulletList.WithItems(bulletListItems).Render()
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("水:%.0f斤 \n"), binfo.Water*2/1000)
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("烘焙时间: %.0f分钟  温度:%.0f度 \n"), binfo.Bakingtime, binfo.Bakingtem)

	}

	return false
}
