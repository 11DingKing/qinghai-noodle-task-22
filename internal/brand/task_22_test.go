package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask22(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	c := CultureCampaign{StoreID: "store-1", Title: "大美青海", ContentVersion: 2, ApprovedVersion: 2, StartsAt: now.Add(-time.Hour), EndsAt: now.Add(time.Hour), FeaturedSKUs: []string{"yak"}, RegionalBrandLogo: true}
	listings := map[string]ProductListing{"yak": {StoreID: "store-1", Published: true}}
	require.NoError(t, s.CheckCampaign(context.Background(), c, compliantStore(now), listings))
}
