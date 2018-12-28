package main

import (
	s "github.com/vladcostea/ginstarter"
)

func main() {
	r := s.NewRoutes()
	r.Run(":8080")
}
