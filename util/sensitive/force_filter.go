package sensitive

import (
	"os"
	"sort"
	"strings"
)

type ForceMatchFilter struct {
	SensitiveWords []string //敏感词数组
}

func (it *ForceMatchFilter) IsSensitive(Text string) bool {
	for _, SensitiveWord := range it.SensitiveWords { //检查是否存在敏感词
		if strings.Contains(Text, SensitiveWord) {
			return true
		}
	}
	return false
}

func (it *ForceMatchFilter) ReplaceSensitiveWord(Text string, ReplaceString string) string {
	for _, SensitiveWord := range it.SensitiveWords { //替换敏感词
		strings.Replace(Text, SensitiveWord, ReplaceString, -1)
	}
	return Text
}

func (it *ForceMatchFilter) readSensitiveWord() []string {
	Text, err := os.ReadFile("../../resource/广告.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(Text), "\n")
}

func (it *ForceMatchFilter) InitFilter() error {
	it.SensitiveWords = it.readSensitiveWord()          //初始化敏感词数组
	sort.Slice(it.SensitiveWords, func(i, j int) bool { //按照敏感词长度降序排列
		return len(it.SensitiveWords[i]) > len(it.SensitiveWords[j])
	})
	return nil
}
