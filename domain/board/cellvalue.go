package board

const (
	// EmptyValue represents an empty cell
	EmptyValue CellValue = 0
	// XValue represents a cell with X
	XValue CellValue = 1
	// OValue represents a cell with O
	OValue CellValue = -1
)

// CellValue represents the value of a cell
type CellValue int

// IsEmpty returns true if the CellValue is empty
func (v CellValue) IsEmpty() bool {
	return v == EmptyValue
}

// String returns the string representation of the CellValue
func (v CellValue) String() string {
	switch v {
	case XValue:
		return "X"
	case OValue:
		return "O"
	default:
		return "-"
	}
}
