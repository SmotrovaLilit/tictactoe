package board

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_cellValue_IsEmpty(t *testing.T) {
	cv := EmptyValue
	require.True(t, cv.IsEmpty())
	cv = XValue
	require.False(t, cv.IsEmpty())
}
