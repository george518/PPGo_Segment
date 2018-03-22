/************************************************************
** @Description: PPGo_Segment
** @Author: george hao
** @Date:   2018-03-22 14:06
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-22 14:06
*************************************************************/
package PPGo_Segment

import (
	"github.com/george518/PPGo_Trie/double_trie"
)

func NewDictionary() *Dictionary {
	return &Dictionary{trie: double_trie.New()}
}

// 向词典中加入一个分词
func (dict *Dictionary) addToken(token Token) {
	bytes := textSliceToBytes(token.text)

	_, err := dict.trie.Get(bytes)
	if err == nil {
		return
	}

	dict.trie.Insert(bytes, dict.NumTokens())
	dict.tokens = append(dict.tokens, token)
	dict.totalFrequency += int64(token.frequency)
	if len(token.text) > dict.maxTokenLength {
		dict.maxTokenLength = len(token.text)
	}
}

// 在词典中查找和字元组words可以前缀匹配的所有分词
// 返回值为找到的分词数
func (dict *Dictionary) lookupTokens(words []Text, tokens []*Token) (numOfTokens int) {
	var id, value int
	var err error
	for _, word := range words {
		id, err = dict.trie.Jump(word, id)
		if err != nil {
			break
		}
		value, err = dict.trie.Value(id)
		if err == nil {
			tokens[numOfTokens] = &dict.tokens[value]
			numOfTokens++
		}
	}
	return
}

// 词典中分词数目
func (dict *Dictionary) NumTokens() int {
	return len(dict.tokens)
}
