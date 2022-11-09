package sensitive

import "testing"

func TestFilter(t *testing.T) {
	filter := AcAutomaton{}
	err := filter.InitFilter()
	if err != nil {
		t.Error(err)
	}

	res := filter.IsSensitive("QQ")
	if !res {
		t.Failed()
	}

	str := filter.ReplaceSensitiveWord("测试QQ", "*")
	if str != "测试**" {
		t.Failed()
	}
}
