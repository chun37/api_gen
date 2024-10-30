module github.com/go-generalize/api_gen/v2/pkg/server/tmp

go 1.18

replace github.com/go-generalize/api_gen/v2/samples => ../../../samples

require (
	github.com/go-generalize/api_gen/v2/samples v0.0.0-00010101000000-000000000000
	github.com/go-utils/echo-multipart-binder v0.2.1
	github.com/labstack/echo/v4 v4.12.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
