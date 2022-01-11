package _4_单词替换

import "strings"

// 哈希表O(n)

func replaceWords2(dictionary []string, sentence string) string {
	mp := map[string]bool{}
	for _, str := range dictionary {
		mp[str] = true
	}
	strArr := strings.Split(sentence, " ")
	for i, str := range strArr {
		var stringB strings.Builder
		for j := 0; j < len(str); j++ {
			stringB.WriteByte(str[j])
			if _, ok := mp[stringB.String()]; ok {
				strArr[i] = stringB.String()
				break
			}
		}
	}
	return strings.Join(strArr, " ")
}

func replaceWords(dictionary []string, sentence string) string {
	My := CreateTire()
	for _, v := range dictionary {
		My.Insert(v)
	}
	str := strings.Split(sentence, " ")
	for i := 0; i < len(str); i++ {
		tmp := My.Search(str[i])
		str[i] = tmp
	}
	return strings.Join(str, " ")
}

// Tire 建立字典树
type Tire struct {
	node [26]*Tire
	isOK bool
}

func CreateTire() *Tire {
	return &Tire{}
}

func (t *Tire) Insert(word string) {
	tmp := t
	for _, v := range word {
		v -= 'a'
		if tmp.node[v] == nil {
			tmp.node[v] = CreateTire()
		}
		tmp = tmp.node[v]
	}
	tmp.isOK = true
}

func (t *Tire) Search(word string) string {
	tmp := t
	length := len(word)
	for i := 0; i < length; i++ {
		if tmp.node[word[i]-'a'] == nil {
			return word
		}
		tmp = tmp.node[word[i]-'a']
		if tmp.isOK {
			return word[:i+1]
		}
	}
	return word
}
