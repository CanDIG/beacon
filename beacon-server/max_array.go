package server

const MaxWorkingStringArray = 1_000

type toomany struct{}

func (t toomany) Error() string {
	return "Too many values to return."
}
