package sensitive

import (
	"Fuever/resource"
	"sort"
	"strings"
)

type ForceMatchFilter struct {
	sensitiveWords []string //敏感词数组
}

func (it *ForceMatchFilter) IsSensitive(text string) bool {
	for _, SensitiveWord := range it.sensitiveWords { //检查是否存在敏感词
		if strings.Contains(text, SensitiveWord) {
			return true
		}
	}
	return false
}

func (it *ForceMatchFilter) ReplaceSensitiveWord(text string, replaceString string) string {
	for _, SensitiveWord := range it.sensitiveWords { //替换敏感词
		if len(SensitiveWord) != 0 {
			ReplaceString := ""
			for index := range SensitiveWord {
				index = 1 + index //无用变量
				ReplaceString = ReplaceString + replaceString
			}
			text = strings.Replace(text, SensitiveWord, ReplaceString, -1)
		}
	}
	return text
}

func (it *ForceMatchFilter) readSensitiveWord() []string {
	return resource.SensitiveWords()
}

func (it *ForceMatchFilter) InitFilter() error {
	it.sensitiveWords = it.readSensitiveWord()          //初始化敏感词数组
	sort.Slice(it.sensitiveWords, func(i, j int) bool { //按照敏感词长度降序排列
		return len(it.sensitiveWords[i]) > len(it.sensitiveWords[j])
	})
	return nil
}
