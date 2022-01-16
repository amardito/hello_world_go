module hello_world/hello

go 1.17

replace hello_world/greetings => ../greetings

replace hello_world/randomize => ../randomize

replace hello_world/customLog => ../customLog

require (
	hello_world/customLog v0.0.0-00010101000000-000000000000
	hello_world/greetings v0.0.0-00010101000000-000000000000
)

require hello_world/randomize v0.0.0-00010101000000-000000000000 // indirect
