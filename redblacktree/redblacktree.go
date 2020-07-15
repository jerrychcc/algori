package redblacktree

//Tree record redblack meta data here
type Tree struct {
	size int64
	root *Node
}

//Node record node info
type Node struct {
	key                 interface{}
	val                 interface{}
	color               color
	left, right, parent *Node
}

type color bool

const (
	//RED mark the tri-node
	RED color = true
	//BLACK mark the double node
	BLACK color = false
)

//Init init the redblacktree
func Init() *Tree {
	return &Tree{}
}

//Size ret the size of the tree
func (t *Tree) Size() int64 {
	return t.size
}

//Get ret the val,the key mark
func (t *Tree) Get(key interface{}) interface{} {

	return "test for locate"
}

//Put put key and val into the tree
func (t *Tree) Put(key interface{}, val interface{}) {

}

//Del del the val key point to
func (t *Tree) Del(key interface{}) bool {

	return true
}

func (t *Tree) setColor() {

}
