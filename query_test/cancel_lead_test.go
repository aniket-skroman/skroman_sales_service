package querytest

import (
	"context"
	"testing"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomCancelLead(t *testing.T) {
	lead_id, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")
	args := db.CreateCancelLeadParams{
		LeadID:   lead_id,
		CancelBy: uuid.New(),
		Reason:   "dsds",
	}

	cancel_lead, err := testQueries.CreateCancelLead(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, cancel_lead)
}

func TestCreateCancelLead(t *testing.T) {
	createRandomCancelLead(t)
}

func TestFetchCancelLeadByLeadId(t *testing.T) {
	lead_id, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")

	cancel_lead, err := testQueries.FetchCancelLeadByLeadId(context.Background(), lead_id)

	require.NoError(t, err)
	require.NotEmpty(t, cancel_lead)
}
