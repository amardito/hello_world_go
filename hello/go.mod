module hello_world/hello

go 1.13

replace hello_world/greetings => ../greetings

require (
	hello_world/customLog v0.0.0-00010101000000-000000000000
	hello_world/greetings v0.0.0-00010101000000-000000000000
)

replace hello_world/randomize => ../randomize

replace hello_world/customLog => ../customLog
