package promcli

type Label string

func (l Label) String() string {
	return string(l)
}

const (
	LabelTotal Label = "total"
	LabelError Label = "error"
)
