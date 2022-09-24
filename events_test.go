package unifi_test

import (
	"testing"

	"github.com/chrisparish/unifi"
	"github.com/stretchr/testify/assert"
)

func TestIPGeoUnmarshalJSON(t *testing.T) {
	t.Parallel()

	a := assert.New(t)
	i := &unifi.IPGeo{}

	a.Nil(i.UnmarshalJSON([]byte(`[]`)))
	a.EqualValues(0, i.Asn)
	a.Nil(i.UnmarshalJSON([]byte(`{"asn": 123}`)))
	a.EqualValues(123, i.Asn)
}
