module example.com/hello

go 1.19

require (
	local.packages/user v0.0.0
)

replace local.packages/user => ./user