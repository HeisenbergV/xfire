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

type buildPrompt struct {
	args map[string]int //参数名；参数数量
}

func (*buildPrompt) Name() string {
	return "buildPrompt"
}
func (*buildPrompt) Level() int {
	return 2
}

func (*buildPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "查看所有制作工艺信息"},
		{Text: "info", Description: "查看具体制作细节 e.g:info id "},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (*buildPrompt) Show() {
	fmt.Println("工艺管理")
	fmt.Println("list info")
	fmt.Println("")
}

func (b *buildPrompt) Exec(args []string) bool {
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
		data, _, err := service.FactoryService.GetBuildList(request.PageInfo{Page: 0, PageSize: 10000}, "", true)
		if err != nil {
			fmt.Println("查询失败")
			return false
		}
		list := data.([]model.Build)
		tableData1 := pterm.TableData{
			{"ID", "名称", "备注", "水", "烘焙时间", "烘焙温度"},
		}
		for i := range list {
			tableData1 = append(tableData1, []string{fmt.Sprintf("%d", list[i].ID), list[i].Name,
				list[i].Remake, fmt.Sprintf("%.2f", list[i].Water),
				fmt.Sprintf("%.2f", list[i].Bakingtime), fmt.Sprintf("%.2f", list[i].Bakingtem),
			})
		}

		pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

	} else if cmd == "info" {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("参数错误")
			return false
		}
		binfo, err := service.FactoryService.GetBuildInfo(id)
		if err != nil {
			fmt.Println("查询失败")
			return false
		}
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("%s \n水:%.0f斤 \n"), binfo.Name, binfo.Water*2/1000)
		pterm.DefaultBasicText.Printf(pterm.LightMagenta("烘焙时间: %.0f分钟  温度:%.0f度 \n"), binfo.Bakingtime, binfo.Bakingtem)

		bulletListItems := []pterm.BulletListItem{
			{
				Level:       0,                           // Level 0 (top level)
				Text:        "配方",                        // Text to display
				TextStyle:   pterm.NewStyle(pterm.FgRed), // Text color
				BulletStyle: pterm.NewStyle(pterm.FgRed), // Bullet color
			},
		}

		for mname, ratio := range binfo.BuildData {
			bulletListItems = append(bulletListItems,
				pterm.BulletListItem{
					Level:       1,                                     // Level 0 (top level)
					Text:        fmt.Sprintf("%s %.2fg", mname, ratio), // Text to display
					TextStyle:   pterm.NewStyle(pterm.FgBlue),          // Text color
					BulletStyle: pterm.NewStyle(pterm.FgRed),           // Bullet color
				})
		}
		// Create a bullet list with the defined items and render it.
		pterm.DefaultBulletList.WithItems(bulletListItems).Render()
	}

	return false
}
