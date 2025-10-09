package one_of

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OneOf(t *testing.T) {
	require := require.New(t)
	asDay := []Day{
		{ValueType: "after_n_days", Value: "two"},
		{ValueType: "before_n_days", Value: "ten"},
		{ValueType: "last_n_days", Value: "five"},
	}
	asHour := []Hour{
		{ValueType: "after_n_hours", Value: 4},
		{ValueType: "after_n_hours", Value: 16},
		{ValueType: "after_n_hours", Value: 256},
	}
	for _, operatorValue := range asDay {
		t.Run(string(operatorValue.ValueType), func(t *testing.T) {
			preset := DateTimePreset{}
			err := preset.FromDay(operatorValue)
			require.NoError(err)

			value, err := preset.AsDay()
			require.NoError(err)
			require.Equal(operatorValue, value)

			valueByDiscriminator, err := preset.ValueByDiscriminator()
			require.NoError(err)
			require.Equal(value, valueByDiscriminator)
		})
	}
	for _, operatorValue := range asHour {
		t.Run(string(operatorValue.ValueType), func(t *testing.T) {
			preset := DateTimePreset{}
			err := preset.FromHour(operatorValue)
			require.NoError(err)

			value, err := preset.AsHour()
			require.NoError(err)
			require.Equal(operatorValue, value)

			valueByDiscriminator, err := preset.ValueByDiscriminator()
			require.NoError(err)
			require.Equal(value, valueByDiscriminator)
		})
	}
}
