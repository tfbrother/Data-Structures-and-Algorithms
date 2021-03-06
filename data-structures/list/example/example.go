package main

import "github.com/tfbrother/Data-Structures-and-Algorithms/data-structures/list"

var (
	seq  *list.SequenceList
	ele1 list.Element
	ele2 list.Element
	ele3 list.Element
	ele4 list.Element
	ele5 list.Element
)

func main() {
	seq = list.NewSequenceList(10)
	ele1 = 1
	ele2 = 2
	ele3 = 3
	ele4 = 4
	ele5 = 5

	seq.InsertIndex(0, &ele1)
	seq.InsertIndex(1, &ele2)
	seq.InsertIndex(2, &ele3)
	seq.InsertIndex(3, &ele4)
	seq.Dump() // 1234

	seq.InsertIndex(1, &ele5)
	seq.Dump() // 15234

	seq.DelIndex(2)
	seq.Dump() // 1534

	ele5 = 8   // 使用引用传递时，外部修改了值会影响内部，慎重使用
	seq.Dump() // 1834
}
