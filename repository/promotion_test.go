package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPromotionRepository(t *testing.T) {
	t.Run("should not return error, and promotion repository is initialize", func(t *testing.T) {
		promotionRepo, err := NewPromotionRepository("data/sample_promotion.json")
		assert.Nil(t, err)
		assert.NotNil(t, promotionRepo)
	})
}
