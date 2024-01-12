package models

type RuleInterface interface {
	GetKey() string
	ValidateConstraint() error
}

type Rule struct {
	Name   string `yaml:"name" json:"name"`
	Window int    `yaml:"window" json:"window"`
	Count  int    `yaml:"count" json:"count"`
}
