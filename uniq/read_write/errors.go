package read_write

type ParseError struct {
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}

var (
	UnknownFlagError = ParseError{"Unknown flag\n" +
		"Run the program with the --help flag to output the supported commands"}
	TogetherArgs      = ParseError{"Flags [-c -d -u] cannot be used together"}
	SkipNegative      = ParseError{"Count of skipped symbols(words) must be a positive number"}
	IncorrectPosition = ParseError{"Positional arguments should go at the end"}
)
