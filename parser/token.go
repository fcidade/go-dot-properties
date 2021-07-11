package parser

const (
	TypeIdentifier = "identifier"
	TypeValue      = "value"
	TypeSeparator  = "separator"
)

type Token struct {
	Text string
	Type string
	// TODO: Line, column
}
