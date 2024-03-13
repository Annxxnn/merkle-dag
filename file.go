package merkledag

// 定义了一个常量 FILE 和 DIR，分别表示文件和目录的类型
const (
	FILE = iota
	DIR
)

// Node 接口，表示 MerkleDag 中的节点。它包含两个方法：
// Size() 方法返回节点的大小（以字节为单位）。
// Type() 方法返回节点的类型（FILE 或 DIR）。
type Node interface {
	Size() uint64
	Name() string
	Type() int
}

// File 接口，表示 MerkleDag 中的文件节点。它继承了 Node 接口，并添加了一个方法：
// Bytes() 方法返回文件节点的内容（以字节数组的形式）。
type File interface {
	Node
	Bytes() []byte
}

// Dir 接口，表示 MerkleDag 中的目录节点。它继承了 Node 接口，并添加了一个方法：
// It() 方法返回一个 DirIterator 对象，用于遍历目录中的子节点。
type Dir interface {
	Node

	It() DirIterator
}

// DirIterator 接口，表示目录迭代器，用于遍历目录中的子节点。它包含两个方法：
// Next() 方法用于移动到目录中的下一个子节点，如果存在下一个子节点则返回 true，否则返回 false。
// Node() 方法返回当前指向的子节点。
type DirIterator interface {
	Next() bool

	Node() Node
}
