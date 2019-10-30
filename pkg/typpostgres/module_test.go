package typpostgres_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
	"github.com/typical-go/typical-rest-server/pkg/typpostgres"
)

func TestModule(t *testing.T) {
	m := typpostgres.Module()
	require.True(t, typiobj.IsProvider(m))
	require.True(t, typiobj.IsDestructor(m))
	require.True(t, typiobj.IsCommandLiner(m))
	require.True(t, typiobj.IsConfigurer(m))
}
