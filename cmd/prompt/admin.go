package prompt

import (
	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
)

type adminPrompt struct{}

func (*adminPrompt) Level() int {
	return 1
}
func (*adminPrompt) Name() string {
	return "adminPrompt"
}
func (*adminPrompt) Exec(args []string) bool {
	return true
}

func (*adminPrompt) Show() {
	pterm.DefaultSection.Println(pterm.Blue("    后台管理"))
	pterm.Println(pterm.Blue("1. brand-品牌管理   2. build-工艺管理"))
	pterm.Println(pterm.Blue("3. customer-客户管理  4. Goods-产品管理"))
	pterm.Println(pterm.Blue("5. material-原料管理"))
}

func (*adminPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "brand", Description: "品牌管理"},
		{Text: "build", Description: "工艺管理"},
		{Text: "customer", Description: "客户管理"},
		{Text: "Goods", Description: "产品管理"},
		{Text: "material", Description: "原料管理"},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
