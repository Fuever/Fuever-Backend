package sensitive

//TODO implement this！
// @Garliczz
// 请实现一个敏感词过滤器
// 放在这个 包 里面就可以
// 如果您对实现接口有所疑虑
// 请参考 dummy_filter.go
// 倒不是说一定要AC自动机
// 越给力的算法越好啦~
// 文件的输入输出也要自己写哦

// WordFilter
// 敏感词过滤器接口
// 所有的敏感词过滤器都应该实现它
type WordFilter interface {

	// IsSensitive 判断某个句子是否含有敏感词
	// 例如:
	// s := "天生万物以养人 人无一物以报天"
	// println(IsSensitive(s))
	// console output: true
	IsSensitive(string) bool

	// ReplaceSensitiveWord 替换某个词句中的敏感词
	// 例如:
	// s := "我是小熊维尼"
	// replaceString := "*"
	// println(ReplaceSensitiveWord(s, replaceString))
	// console output: 我是****
	// 哦对了 如果你看到这里
	// 匹配不成功也是很正常的
	// 这只是举个匹配成功的例子
	// 如果没匹配到就原值返回
	ReplaceSensitiveWord(string, string) string
	// InitFilter
	// 初始化敏感词过滤器
	// 数据结构的建立就是这个时候啦！
	InitFilter() error

	// 读入敏感词数据
	// 文件在 Fuever/resource 下
	// 值得一提的是
	// 数据量很小
	readSensitiveWord() []string
}

var filter WordFilter

func GetFilter() WordFilter {
	if filter == nil {
		filter = &AcAutomaton{}
		filter.InitFilter()
	}
	return filter
}
