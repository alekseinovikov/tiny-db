package tree

import "tiny-db/tools"

type BTree struct {
	// pointer (a non-negative integer) to the root node of the tree
	root uint64

	// callbacks for mapping on-disk -pages
	get func(uint64) BNode // dereference a node from disk
	new func() uint64      // create a new node on disk
	del func(uint64)       // delete a node from disk
}

const (
	BTREE_PAGE_SIZE    = 4096
	BTREE_MAX_KEY_SIZE = 1000
	BTREE_MAX_VAL_SIZE = 3000
)

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	tools.Assert(node1max <= BTREE_PAGE_SIZE)
}
