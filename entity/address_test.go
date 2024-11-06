package entity_test

import (
	"testing"

	"github.com/felipehrs/go-expert-cloud-run/entity"
	"github.com/stretchr/testify/assert"
)

func TestIsValidZipCode(t *testing.T) {
	assert.True(t, entity.IsValidZipCode("01001-000"))
	assert.True(t, entity.IsValidZipCode("01001000"))
	assert.False(t, entity.IsValidZipCode("01001-0000"))
	assert.False(t, entity.IsValidZipCode("0100-100"))
	assert.False(t, entity.IsValidZipCode("abcde-fgh"))
}
