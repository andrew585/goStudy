package schema

import "testing"

func TestCreateSchema(t *testing.T) {
	NewDgrapClient()
}

func TestCreateSchema2(t *testing.T) {
	CreateSchema(NewDgrapClient())
}
