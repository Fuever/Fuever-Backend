package sensitive

import "Fuever/resource"

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

func (it *AcAutomaton) Replace(str, replaceString string) string { //消除敏感词
	byteStr := []byte(str)
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
			byteStr[i] = '*'
			j--
		}
	}
	return replace(string(byteStr), str, replaceString)
}

func (it *AcAutomaton) IsSensitive(str string) bool {
	return it.Check(str)
}

func (it *AcAutomaton) ReplaceSensitiveWord(str string, replaceString string) string {
	return it.Replace(str, replaceString)
}

func (it *AcAutomaton) readSensitiveWord() []string {
	return resource.SensitiveWords()
}

func (it *AcAutomaton) InitFilter() error {
	strArr := it.readSensitiveWord()
	it.init()
	for _, str := range strArr {
		it.insert(str)
	}
	it.build()
	return nil
}
