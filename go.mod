module 425-a1

require (
	github.com/timtadh/data-structures v0.5.3 // indirect
	github.com/timtadh/lexmachine v0.2.2
)

replace github.com/timtadh/data-structures => ./vendor/github.com/timtadh/data-structures

replace github.com/timtadh/getopt => ./vendor/github.com/timtadh/getopt

replace github.com/timtadh/lexmachine => ./vendor/github.com/timtadh/lexmachine

go 1.13
