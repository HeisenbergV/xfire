package prompt

import (
	"fmt"
	"strconv"
	"xfire/model"
	"xfire/service"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
)

type materialPrompt struct {
	args map[string]int //参数名；参数数量
}

func (*materialPrompt) Name() string {
	return "materialPrompt"
}
func (*materialPrompt) Level() int {
	return 2
}

func (*materialPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "查看所有原料信息"},
		{Text: "add", Description: "增加新原料 e.g: add name brand info unit price"},
		{Text: "update", Description: "修改原料信息 e.g: update id 字段名称 修改结果"},
		{Text: "rm", Description: "删除原料 e.g: rm id"},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (*materialPrompt) Show() {
	fmt.Println("原料信息管理")
	fmt.Println("list add update rm")
	fmt.Println("")
}

func (b *materialPrompt) Exec(args []string) bool {
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
		list, err := service.FactoryService.GetMaterialList()
		if err != nil {
			fmt.Println("查询失败")
			return false
		}

		// Define the data for the first table
		tableData1 := pterm.TableData{
			{"ID", "原料", "品牌", "重量(kg)", "单价", "备注", "创建时间"},
		}
		for i := range list {
			tableData1 = append(tableData1, []string{fmt.Sprintf("%d", list[i].ID),
				list[i].Name, list[i].Brand, fmt.Sprintf("%.2f", list[i].Unit),
				fmt.Sprintf("%.2f", list[i].Unit), list[i].Info, list[i].CreatedAt.Format("2006-01-02")})
		}
		pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()
	} else if cmd == "add" {
		unit, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			fmt.Println("参数错误")
			return false
		}
		p, err := strconv.ParseFloat(args[5], 64)
		if err != nil {
			fmt.Println("参数错误")
			return false
		} ////add 酵母 安琪 三合的 10 190
		//name brand info unit price
		err = service.FactoryService.AddMaterial(model.Material{Name: args[1], Brand: args[2],
			Info: args[3], Unit: unit, Price: p,
		})
		if err != nil {
			fmt.Println("添加失败")
		} else {
			fmt.Println("添加成功")
		}

	} else if cmd == "update" {
		fmt.Println("todo")

	} else if cmd == "rm" {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("参数错误")
			return false
		}

		err = service.FactoryService.DelMaterial(id)
		if err != nil {
			fmt.Println("删除失败")
		} else {
			fmt.Println("删除成功")
		}
	}

	return false
}
