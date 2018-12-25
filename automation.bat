 protoc.exe -I .\src\github.com\MilanSavaliya\go-poc-proto-files --go_out=$env:GOPATH\src .\src\github.com\MilanSavaliya\go-poc-proto-files\types\department.proto
 protoc.exe -I .\src\github.com\MilanSavaliya\go-poc-proto-files --go_out=$env:GOPATH\src .\src\github.com\MilanSavaliya\go-poc-proto-files\types\teacher.proto

 protoc.exe -I .\src\github.com\MilanSavaliya\go-poc-proto-files --go_out=plugins=grpc:$env:GOPATH\src .\src\github.com\MilanSavaliya\go-poc-proto-files\controller\teacher-controller.proto
