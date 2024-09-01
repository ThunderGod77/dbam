package core

type TableType int

const (
	BaseTable TableType = iota + 1 // EnumIndex = 1
	View                           // EnumIndex = 2
	Foreign
	LocalTemporary
)

func (t TableType) String() string {
	return [...]string{"BASE TABLE", "VIEW", "FOREIGN", "LOCAL TEMPORARY"}[t-1]
}

func (t TableType) EnumIndex() int {
	return int(t)
}
