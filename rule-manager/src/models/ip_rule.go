package models

import (
	"errors"
	"fmt"
)

type IpRule struct {
	IpAddress string
	Rule
}

func (self IpRule) GetKey() string {
	return fmt.Sprint("IpAddr::%s", self.IpAddress)
}

func (self IpRule) ValidateConstraint() error {
	if self.Count <= 0 {
		return errors.New("Count should have the valid value")
	}

	if self.Window <= 0 {
		return errors.New("Window should have valid value")
	}
	return nil
}
