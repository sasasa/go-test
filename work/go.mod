module github.com/sasasa/go-test

go 1.19

require (
	local.packages/user v0.0.0
	local.packages/utility v0.0.0
)

replace local.packages/user => ./user
replace local.packages/utility => ./utility