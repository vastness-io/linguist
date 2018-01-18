package model

import (
	"strings"
)

type RepositoryInfo struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	Repository string `json:"repository"`
}

func (r RepositoryInfo) String() string {
	return strings.Join([]string{
		r.Host,
		r.User,
		r.Repository}, "/")
}
