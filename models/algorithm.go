package models

type Algorithm struct {
	Id          int64
	Name        string
	Version     string
	Description string
	Author      string
	Parameters  string
	Image       string
	Downloads   int64
	Stars       int64
}
