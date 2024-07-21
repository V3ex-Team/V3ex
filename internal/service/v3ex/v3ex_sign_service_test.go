package v3ex

import (
	"context"
	"github.com/apache/incubator-answer/internal/schema"
	"testing"
)

func TestSign(t *testing.T) {
	ctx := context.Background()
	svc := NewV3exService(nil)
	err := svc.Sign(ctx, schema.SignV3exReq{
		ChainId: 1,
		Amount:  100,
		To:      "0x1234567890",
		Date:    "20210101",
	})
	if err != nil {
		t.Error(err)
	}
}
