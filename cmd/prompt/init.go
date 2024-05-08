package prompt

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/pterm/pterm"
)

type PromptHandler interface {
	Prompt(d prompt.Document) []prompt.Suggest
	Show()
	Exec(args []string) bool
	Level() int
	Name() string
}

type basicPrompt struct{}

func (*basicPrompt) Name() string {
	return "basicPrompt"
}
func (*basicPrompt) Exec(args []string) bool {
	return true
}

func (*basicPrompt) Level() int {
	return 1
}

func (*basicPrompt) Show() {
	pterm.DefaultSection.Println(pterm.Blue("    工厂管理系统\n1. admin-后台管理   2. drp-进销存   3. finance-财务"))
}

func (*basicPrompt) Prompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "admin", Description: "管理后台"},
		{Text: "drp", Description: "进销存货物原料"},
		{Text: "finance", Description: "财务管理"},
		{Text: "q", Description: "离开"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

var PromptMap map[string]PromptHandler

func register() {
	PromptMap = make(map[string]PromptHandler)
	PromptMap["basic"] = &basicPrompt{}
	PromptMap["admin"] = &adminPrompt{}
	PromptMap["Goods"] = &GoodsPrompt{args: map[string]int{
		"list": 1,
		"info": 2,
	}}
	PromptMap["build"] = &buildPrompt{args: map[string]int{
		"list": 1,
		"info": 2,
	}}
	PromptMap["material"] = &materialPrompt{args: map[string]int{
		"list": 1,
		"add":  6,
		"rm":   2,
	}}
	PromptMap["brand"] = &brandPrompt{args: map[string]int{
		"list":   1,
		"add":    2,
		"update": 3,
		"rm":     2,
	}}
}

func handler(h PromptHandler) string {
	if h == nil {
		fmt.Println("没有这个指令")
		return ""
	}

	h.Show()

	input := prompt.Input("> ", h.Prompt)
	if input == "q" {
		return "q"
	}

	if h.Level() != 2 {
		handler(PromptMap[input])
	}

	execInput := input
	for {
		inputargs := strings.Split(execInput, " ")
		if h.Exec(inputargs) {
			break
		}
		execInput = prompt.Input("> ", h.Prompt)
	}
	return ""
}

func Run() {
	register()
	for {
		if handler(PromptMap["basic"]) == "q" {
			fmt.Println("bye bye")
			break
		}
	}
}
