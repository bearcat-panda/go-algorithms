package unionfind

import (
	"fmt"
	"testing"
)

func TestNewUnionFind1(t *testing.T) {
	fmt.Println("============UF1==========")
	uf := NewUnionFind1(5)
	fmt.Println("uf:", uf)
	fmt.Println("1的id:", uf.Find(1))
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("0和3连接中...")
	uf.Union(0, 3)
	fmt.Println("0和3的连接关系：", uf.IsConnected(0, 3))
	fmt.Println("uf:", uf)
}
