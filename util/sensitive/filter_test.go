package sensitive

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	filter := ForceMatchFilter{}
	err := filter.InitFilter()
	if err != nil {
		t.Error("暴力	读入敏感词失败")
	}
	acFilter := AcAutomaton{}
	err = acFilter.InitFilter()
	if err != nil {
		t.Error("自动机读入敏感词失败")
	}

	if !filter.IsSensitive("扣扣") {
		t.Error("暴力检测敏感词失败")
	}
	if !acFilter.IsSensitive("扣扣") {
		t.Error("自动机检测敏感词失败")
	}

	content := "Q兼职扣扣淘宝Q"
	replaceString := "哈哈哈"
	text := filter.ReplaceSensitiveWord(content, replaceString)
	acText := acFilter.ReplaceSensitiveWord(content, replaceString)

	if text != acText {
		t.Error("替换敏感词失败")
		fmt.Println("暴力输出：" + text)
		fmt.Println("自动机输出：" + acText)
	}

}
