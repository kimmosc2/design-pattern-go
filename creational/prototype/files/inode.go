package files

type iNode interface {
	print(indentation string)
	clone() iNode
}
