module expresso_example

go 1.19

require github.com/pr47h4m/expresso v0.0.10

replace github.com/pr47h4m/expresso v0.0.10 => ../

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
