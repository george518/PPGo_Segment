/************************************************************
** @Description: PPGo_Segment
** @Author: george hao
** @Date:   2018-03-22 13:19
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-22 13:19
*************************************************************/
package PPGo_Segment

import (
	"github.com/george518/PPGo_Trie/double_trie"
)

type Dictionary struct {
	trie           *double_trie.Dat // Cedar 前缀树
	maxTokenLength int              // 词典中最长的分词
	tokens         []Token          // 词典中所有的分词，方便遍历
	totalFrequency int64            // 词典中所有分词的频率之和
}

// 字串类型，可以用来表达
//	1. 一个字元，比如"中"又如"国", 英文的一个字元是一个词
//	2. 一个分词，比如"中国"又如"人口"
//	3. 一段文字，比如"中国有十三亿人口"
type Text []byte

// 一个分词
type Token struct {
	// 分词的字串，这实际上是个字元数组
	text []Text

	// 分词在语料库中的词频
	frequency int

	// log2(总词频/该分词词频)，这相当于log2(1/p(分词))，用作动态规划中
	// 该分词的路径长度。求解prod(p(分词))的最大值相当于求解
	// sum(distance(分词))的最小值，这就是“最短路径”的来历。
	distance float32

	// 词性标注
	pos string

	// 该分词文本的进一步分词划分，见Segments函数注释。
	segments []*Segment
}

// 文本中的一个分词
type Segment struct {
	// 分词在文本中的起始字节位置
	start int

	// 分词在文本中的结束字节位置（不包括该位置）
	end int

	// 分词信息
	token *Token
}

// 分词器结构体
type Segmenter struct {
	dict *Dictionary
}

// 该结构体用于记录Viterbi算法中某字元处的向前分词跳转信息
type jumper struct {
	minDistance float32
	token       *Token
}
