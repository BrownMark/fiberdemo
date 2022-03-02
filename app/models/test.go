package models

type Test struct {
	Name   string `json:"name" validate:"required,min=3,max=32"`
	Age    int    `json:"age" validate:"required,number"`
	TestId *int   `json:",omitempty"`
} //@name Test
