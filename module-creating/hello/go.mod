module hello

go 1.18

replace (
example.com/greetings => ../greetings
)

require (
	example.com/greetings v1.0.0
)