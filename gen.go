// +build generate
package main

//go:generate find ./candig-schemas/src/main/proto -type f -fprint gen2.go
//go:generate sed -i "s|^|//go:generate protoc -I candig-schemas/src/main/proto --go_out=./vendor/ |" gen2.go
//go:generate sed -i "1i// +build generate" gen2.go
//go:generate sed -i "2ipackage main" gen2.go
