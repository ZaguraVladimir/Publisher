package publisher

type publisher struct {
	sourceDir, resultDir string
	books                map[string]bool
}

func NewPublisher(sourceDir, resultDir string) *publisher {
	p := publisher{
		sourceDir: sourceDir,
		resultDir: resultDir,
		books:     nil,
	}
	p.Scan()
	return &p
}

func (p publisher) Scan() {

	//ioutil.ReadDir(path)
}
