package dependency

// Autogenerated by Typical-Go. DO NOT EDIT.

import "github.com/typical-go/typical-rest-server/typical"

func init() {
	typical.Context.MockTargets.Append("app/repository/book_repo.go")
	typical.Context.MockTargets.Append("app/service/book_service.go")
}