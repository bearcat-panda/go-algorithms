package mybinarysearch

import (
	"fmt"
	"testing"
)

func TestBSTInset(t *testing.T) {
	//bst := NewBinarySearchTree()
	//fmt.Println("空？", bst.IsEmpty())
	//fmt.Println("inserting...")
	//bst.Insert("28", "a")
	//bst.Insert("16", "b")
	//bst.Insert("30", "c")
	//bst.Insert("13", "d")
	//bst.Insert("22", "e")
	//bst.Insert("29", "f")
	//bst.Insert("42", "g")
	//fmt.Println("空？", bst.IsEmpty())
	//fmt.Println("max的key为：", bst.MaxKey(), "max的value为：", bst.MaxKeyValue())
	//fmt.Println("min的key为：", bst.MinKey(), "min的value为：", bst.MinKeyValue())
	//fmt.Println(bst.Search("100"))
	//fmt.Println(bst.Search("30"))
	//fmt.Println("前序遍历：",bst.PreOrder())
	//fmt.Println("中序遍历：",bst.InOrder())
	//fmt.Println("后序遍历：",bst.PostOrder())
	//fmt.Println("层序遍历：",bst.LevelOrder())
	//fmt.Println("删除最小：")
	//bst.RemoveMin()
	//fmt.Println("中序第一个为最小：",bst.InOrder())
	//fmt.Println("删除最大：")
	//bst.RemoveMax()
	//fmt.Println("中序最后一个为最大：",bst.InOrder())

	fmt.Print("============")
	bst1 := NewBinarySearchTree()
	fmt.Println("inserting...")
	bst1.Insert("41", "a")
	bst1.Insert("45", "b")
	bst1.Insert("43", "c")
	bst1.Insert("48", "d")
	bst1.Insert("42", "e")
	bst1.Insert("44", "f")
	bst1.Insert("46", "g")
	bst1.Insert("50", "h")
	fmt.Println("bst1中序遍历：", bst1.InOrder())

	fmt.Println("删除带有左右孩子的节点45。。。")
	bst1.Remove("45")
	fmt.Println("bst1中序遍历：", bst1.InOrder())

	fmt.Println("test Traverse()")
	bst1.root.Traverse()

	count := 0
	bst1.root.TraverseFunc(func(n *Node) {
		count++
	})
	fmt.Printf("num of nodes is : %d", count)
	fmt.Println()

	c := bst1.root.TraverseWithChannel()
	maxNode := "0"
	for node := range c {
		if node.key > maxNode {
			maxNode = node.key
		}
	}
	fmt.Printf("maxNode value is: %s", maxNode)
}
