module checkBuild

go 1.20

require github.com/no-src/fserver v0.0.2

require github.com/no-src/log v0.3.0 // indirect

replace github.com/no-src/fserver v0.0.2 => ./fserver
