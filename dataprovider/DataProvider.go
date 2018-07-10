package dataprovider

type DataType struct {
	ID int32
	Name string
}

type DataProvider interface {
	GetData() ([]DataType, error)
}
