package print


type computer interface {
	print() string
	setPrinter(p printer)
}
