package sensitive

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	filter := ForceMatchFilter{}
	err := filter.InitFilter()
	if err != nil {
		t.Error("读入敏感词失败")
	}

	if !filter.IsSensitive("QQ") {
		t.Error("检测敏感词失败")
	}

	text := filter.ReplaceSensitiveWord("QQ", "*")
	if text != "*" {
		t.Error("替换敏感词失败")
		fmt.Println("期望输出：*")
		fmt.Println("实际输出：" + text)
	}

}
