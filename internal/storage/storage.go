package storage

type AddMaterialDto struct {
	Title        string
	Measure      string
	Cost         int64
	Weight       int64
	Manufacturer string
	Barcode      string
	Img          string
	SectionID    string
}
