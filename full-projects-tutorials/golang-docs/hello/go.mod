module hello

go 1.15

replace example.com/greetings => ../greetings

replace example.com/user/hello/morestrings => ./morestrings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	example.com/user/hello/morestrings v0.0.0-00010101000000-000000000000
)
