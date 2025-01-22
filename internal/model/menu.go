package model

type MenuType string 

type MenuItem struct {
	Name string
	Price int
	KodeRahasia string
	Type MenuType
}