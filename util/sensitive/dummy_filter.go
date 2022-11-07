package sensitive

type DummyFilter struct {
}

func (d *DummyFilter) IsSensitive(sentence string) bool {
	return true
}

func (d *DummyFilter) ReplaceSensitiveWord(sentence string, replaceString string) string {
	return sentence
}

func (d *DummyFilter) InitFilter() error {
	d.readSensitiveWord()
	return nil
}

func (d *DummyFilter) readSensitiveWord() []string {
	return nil
}

var _ WordFilter = (*DummyFilter)(nil)
