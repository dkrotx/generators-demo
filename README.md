# This is a demo of using go generators

Just run make. It automatically:
- build special binaries (./bin/gowrap)
- run `go generate` which uses the above binaries
- runs the final code

## gowrap
gowrap is designed to wrap interfaces. Very flexible with templates language.  
go:generate call is defined client/interface.go and it uses `client/withlog.tmpl` as a template.  
The generation itself called by `make go-generate`.
