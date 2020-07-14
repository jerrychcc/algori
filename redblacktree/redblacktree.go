package redblacktree

// record meta data here
type tree struct {
	size int64
	root *Node
}

// record node info
type Node struct {
	key interface{}
	val interface{}
	color Color
	left,right,parent *Node
}

type Color bool

const (
	RED   Color= true
	BLACK Color= false
)

func Init() {
	return &tree{}
}

func (t *tree)Size() int64 {
	return t.size
}

func (t *tree) Get(key interface{}) interface{} {

}

func (t *tree) Put(key interface{}, val interface{}) {

}