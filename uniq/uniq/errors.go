package uniq

type ProcError struct {
	Message string
}

func (e ProcError) Error() string {
	return e.Message
}

var (
	IncorrectArgs = ProcError{Message: "Incorrect argument\n" +
		"Run the program with the --help flag to output supported commands"}
)
