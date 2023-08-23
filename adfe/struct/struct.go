package main

import "fmt"

type Student struct {
	Name  string
	Score int
}
type Node struct {
	Student Student
	Left    *Node
	Right   *Node
}

type BinaryTree struct {
	Root *Node
}

// 中序遍历打印学生信息（按分数从小到大排序）
func (t *BinaryTree) PrintInOrder(node *Node) {
	if node != nil {
		fmt.Println("=========", node.Student.Name)
		t.PrintInOrder(node.Left)
		fmt.Printf("%s, %d\n", node.Student.Name, node.Student.Score)
		t.PrintInOrder(node.Right)
	}
}

// 在二叉树中插入学生信息
func (t *BinaryTree) InsertStudent(s Student) {
	newNode := &Node{Student: s}

	if t.Root == nil {
		t.Root = newNode
		return
	}

	current := t.Root
	for {
		if s.Score < current.Student.Score {
			if current.Left == nil {
				current.Left = newNode
				return
			} else {
				current = current.Left
			}
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			} else {
				current = current.Right
			}
		}
	}
}

func main() {
	tree := BinaryTree{}

	s1 := Student{Name: "Alice", Score: 85}
	s2 := Student{Name: "Bob", Score: 95}
	s3 := Student{Name: "Charlie", Score: 76}
	s4 := Student{Name: "David", Score: 92}

	tree.InsertStudent(s1)
	tree.InsertStudent(s2)
	tree.InsertStudent(s3)
	tree.InsertStudent(s4)

	tree.PrintInOrder(tree.Root)

}
