package model

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IdentifierTestSuite struct {
	suite.Suite
}

func TestIdentifierTestSuite(t *testing.T) {
	suite.Run(t, new(IdentifierTestSuite))
}

func (suite *IdentifierTestSuite) TestGetStringValue() {

	id := "2754407e-56f2-4e24-ad24-1f8d2994b8f8"

	identifier := Identifier(id)
	suite.Equal(id, identifier.String())
}
