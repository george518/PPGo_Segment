/************************************************************
** @Description: PPGo_Segment
** @Author: george hao
** @Date:   2018-03-22 14:27
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-22 14:27
*************************************************************/
package PPGo_Segment

import (
	"fmt"
	"testing"
)

func TestPrintSegment(t *testing.T) {

}

func TestSegmenter_LoadDictionary(t *testing.T) {
	//载入字典
	var seg Segmenter
	seg.LoadDictionary("testData/dictionary.txt")

	str := "吃葡萄不吐葡萄皮，不吃葡萄倒吐葡萄皮"
	segs := seg.Segment([]byte(str))
	fmt.Print(SegmentsToString(segs, false))
}
