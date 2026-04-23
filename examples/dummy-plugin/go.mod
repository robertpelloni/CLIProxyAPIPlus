module dummy-plugin

go 1.26.0

replace github.com/router-for-me/CLIProxyAPI/v6 => ../..

require github.com/router-for-me/CLIProxyAPI/v6 v6.9.31

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	golang.org/x/sys v0.38.0 // indirect
)
