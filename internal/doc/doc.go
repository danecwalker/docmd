package doc

type Doc struct {
	OutputPath string
}

func NewDoc(outputPath string) *Doc {
	op := outputPath
	if op == "" {
		op = "dist"
	}

	return &Doc{
		OutputPath: op,
	}
}
