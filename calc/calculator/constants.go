package calculator

var priority = map[string]int{
	"+": 2,
	"-": 2,
	"/": 3,
	"*": 4,
	"(": 1,
}

var validOperations = map[string]bool{
	"+": true,
	"-": true,
	"/": true,
	"*": true,
	"(": true,
}

var validTokens = map[string]bool{
	"(": true,
	")": true,
	"-": true,
	"+": true,
	"/": true,
	"*": true,
}
