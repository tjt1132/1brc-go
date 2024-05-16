package aggregator

type FileOpenErr struct {
	Err error
}

func (e *FileOpenErr) Error() string {
	return "could not open the file - reason: " + e.Err.Error()
}

func (e *FileOpenErr) Unwrap() error {
	return e.Err
}

type ParceFloatErr struct {
	Err error
}

func (e *ParceFloatErr) Error() string {
	return "could not parce to float64 - reason: " + e.Err.Error()
}

func (e *ParceFloatErr) Unwrap() error {
	return e.Err
}
