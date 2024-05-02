package prompt

import (
	"fmt"
	"strconv"
	"xfire/model"
	"xfire/model/request"
	"xfire/service"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
)

type brandPrompt struct {
	args map[string]int //参数名；参数数量
}

func (*brandPrompt) Name() string {
	return "brandPrompt"
}
func (*brandPrompt) Level() int {
	return 2
}

func (*brandPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "查看所有品牌"},
		{Text: "add", Description: "增加品牌 e.g: add xxx"},
		{Text: "update", Description: "修改品牌 e.g: update id xxx"},
		{Text: "rm", Description: "删除品牌 e.g: rm id"},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (*brandPrompt) Show() {
	fmt.Println("品牌管理")
	fmt.Println("list add update rm")
	fmt.Println("")
}

func (b *brandPrompt) Exec(args []string) bool {
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
		data, _, err := service.FactoryService.GetBrandList(request.PageInfo{Page: 0, PageSize: 1000, Keyword: "name"}, "id", true)
		if err != nil {
			fmt.Println("查询失败")
			return false
		}
		list := data.([]model.Brand)
		// Define the data for the first table
		tableData1 := pterm.TableData{
			{"ID", "品牌名称", "创建时间"},
		}
		for i := range list {
			tableData1 = append(tableData1, []string{fmt.Sprintf("%d", list[i].ID), list[i].Name, list[i].CreatedAt.Format("2006-01-02")})
		}
		pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()
	} else if cmd == "add" {
		err := service.FactoryService.CreateBrand(model.Brand{Name: args[1]})
		if err != nil {
			fmt.Println("添加失败")
		} else {
			fmt.Println("添加成功")
		}

	} else if cmd == "update" {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("参数错误")
			return false
		}
		err = service.FactoryService.UpdateBrand(model.Brand{ID: id, Name: args[2]})
		if err != nil {
			fmt.Println("修改失败")
		} else {
			fmt.Println("修改成功")
		}

	} else if cmd == "rm" {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("参数错误")
			return false
		}

		err = service.FactoryService.DelBrand([]int{id})
		if err != nil {
			fmt.Println("删除失败")
		} else {
			fmt.Println("删除成功")
		}
	}

	return false
}
