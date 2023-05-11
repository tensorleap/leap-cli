module github.com/tensorleap/cli-go

go 1.20

require (
	github.com/spf13/cobra v1.7.0
	github.com/tensorleap/cli-go/pkg/tensorleapapi v0.0.0-00010101000000-000000000000
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tensorleap/cli-go/pkg/tensorleapapi => ./pkg/tensorleapapi
