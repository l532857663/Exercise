package model

import "fmt"

type Service struct {
	Name string
	Age  int
}

func (s *Service) Ceshi() string {
	return fmt.Sprintf("the service is test '%s'!\n", s.Name)
}
