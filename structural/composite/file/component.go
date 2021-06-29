package file

type component interface {
	search(s string) (string, error)
}
