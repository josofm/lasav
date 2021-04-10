package lasav_test

import (
	"testing"

	"github.com/josofm/lasav"

	"github.com/stretchr/testify/assert"
)

func TestShouldTestTheTest(t *testing.T) {
	lasav.DoRequest()

	assert.NotNil(t, nil, "Should be null!")

}
