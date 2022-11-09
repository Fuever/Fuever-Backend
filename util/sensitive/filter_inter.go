package sensitive

import (
	"io/ioutil"
	"os"
	"strings"
)

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

type Node struct { //trie树节点

	next [256]*Node
	fail *Node
	len  int
}

type AcAutomaton struct { //AC自动机

	root *Node
}

func (it *AcAutomaton) init() { //初始化trie树根节点
	it.root = new(Node)
}

func (it *AcAutomaton) insert(str string) { //往trie树插入匹配串
	node := it.root
	lim := len(str)
	for i := 0; i < lim; i++ {
		if node.next[str[i]] == nil {
			node.next[str[i]] = new(Node)
		}
		node = node.next[str[i]]
	}
	node.len = len(str)
}

func (it *AcAutomaton) build() { //构建自动机
	var q []*Node
	front, tear := 0, -1
	node := it.root
	for i := 0; i < 256; i++ {
		if node.next[i] != nil {
			q = append(q, node.next[i])
			node.next[i].fail = node
			tear++
		} else {
			node.next[i] = node
		}
	}
	for ; front <= tear; front++ {
		node = q[front]
		for i := 0; i < 256; i++ {
			if node.next[i] != nil {
				node.next[i].fail = node.fail.next[i]
				q = append(q, node.next[i])
				tear++
			} else {
				node.next[i] = node.fail.next[i]
			}
		}
	}
}

func replace(str1, str2, str3 string) string { //*
	var str []byte
	lim := len(str1)
	for i := 0; i < lim; {
		if str1[i] == '*' {
			str = append(str, '*')
			if str2[i] < 128 {
				i++
			} else {
				i += 3
			}
		} else {
			str = append(str, str1[i])
			i++
		}
	}
	var res []byte
	replacestring := []byte(str3)
	lim = len(str)
	for i := 0; i < lim; i++ {
		if str[i] == '*' {
			res = append(res, replacestring...)
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}

func (it *AcAutomaton) Check(str string) bool { //检查是否存在敏感词
	node := it.root
	lim := len(str)
	for i := 0; i < lim; i++ {
		node = node.next[str[i]] //转移
		if node.len != 0 {
			return true
		}
	}
	return false
}

func (it *AcAutomaton) Replace(str, replacestring string) string { //消除敏感词
	byte_str := []byte(str)
	node := it.root
	var tag []int
	lim := len(str)
	for i := 0; i < lim; i++ {
		node = node.next[str[i]]    //转移
		tag = append(tag, node.len) //记录出现的敏感词长度
	}
	for i, j := lim-1, 0; i > -1; i-- {
		if tag[i] > j {
			j = tag[i]
		}
		if j != 0 {
			byte_str[i] = '*'
			j--
		}
	}
	return replace(string(byte_str), str, replacestring)
}

func (it *AcAutomaton) IsSensitive(str string) bool {
	return it.Check(str)
}

func (it *AcAutomaton) ReplaceSensitiveWord(str string, replacestring string) string {
	return it.Replace(str, replacestring)
}

func (it *AcAutomaton) readSensitiveWord() []string {
	f, err := os.Open("../../resource/广告.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	// err
	str := string(b)
	return strings.Split(str, "\n")
}

func (it *AcAutomaton) InitFilter() error {
	strs := it.readSensitiveWord()
	it.init()
	for _, str := range strs {
		it.insert(str)
	}
	it.build()
	return nil
}
