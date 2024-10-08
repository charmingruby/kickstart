package entity

import (
	"testing"
	"time"

	"github.com/charmingruby/kickstart/internal/common/core/validation"
	"github.com/stretchr/testify/assert"
)

func Test_NewExample(t *testing.T) {
	var name string = "dummy example"

	t.Run("it should be able to create an example", func(t *testing.T) {
		example, err := NewExample(name)
		dateInThePast := time.Now().Add(-time.Minute)

		assert.NoError(t, err)
		assert.NotNil(t, example.GetID())
		assert.Equal(t, name, example.GetName())
		assert.Less(t, dateInThePast, example.GetCreatedAt())
	})

	t.Run("it should be not able to create an example with a short name", func(t *testing.T) {
		example, err := NewExample("na")

		assert.Error(t, err)
		assert.Nil(t, example)
		assert.Equal(t, validation.ErrMinLength("name", "3"), err.Error())
	})

	t.Run("it should be not able to create an example with a long name", func(t *testing.T) {
		var longName string
		maxSize := 17

		for range maxSize {
			longName += "a"
		}

		example, err := NewExample(longName)

		assert.Error(t, err)
		assert.Nil(t, example)
		assert.Equal(t, validation.ErrMaxLength("name", "16"), err.Error())
	})
}
