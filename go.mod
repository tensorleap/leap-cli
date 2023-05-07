module github.com/tensorleap/cli-go

go 1.20

require (
	github.com/spf13/cobra v1.7.0
	github.com/tensorleap/cli-go/pkg/tensorleapapi v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/tensorleap/cli-go/pkg/tensorleapapi => ./pkg/tensorleapapi
