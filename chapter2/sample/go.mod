// https://go.dev/ref/mod
// https://go.dev/doc/modules/gomod-ref
// https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
// https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/

module sample
go 1.18

replace github.com/goinaction/code/chapter2/sample/matchers => ./matchers
replace github.com/goinaction/code/chapter2/sample/search => ./search
require (
	github.com/goinaction/code/chapter2/sample/matchers v1.0.0
)
require (
	github.com/goinaction/code/chapter2/sample/search v1.0.0
)

