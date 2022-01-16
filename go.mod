module hello_world/main

go 1.17

replace hello_world/hello => ./hello

replace hello_world/greetings => ./greetings

replace hello_world/customLog => ./customLog

replace hello_world/randomize => ./randomize

require hello_world/hello v0.0.0-00010101000000-000000000000

require (
	hello_world/customLog v0.0.0-00010101000000-000000000000 // indirect
	hello_world/greetings v0.0.0-00010101000000-000000000000 // indirect
	hello_world/randomize v0.0.0-00010101000000-000000000000 // indirect
)
