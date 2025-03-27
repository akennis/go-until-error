# go-until-error
Streamlines Go code wrapping functions so they can be run in succession without checking their error returns one by one, instead capturing any first error that occurs and then no-op'ing further function calls with the error available for final checking.

# example



## code-gen

go run .\internal\codegen\main.go | Out-File -encoding utf8 go-until-error.go