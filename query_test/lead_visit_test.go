package querytest

import (
	"context"
	"testing"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomLeadVisit(t *testing.T) {
	lead, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")
	visit_by, _ := uuid.Parse("30c1dfa2-0aa2-4653-9c8e-ae02bb209a06")
	args := db.CreateLeadVisitParams{
		LeadID:          lead,
		VisitBy:         visit_by,
		VisitDiscussion: "test",
	}

	visit, err := testQueries.CreateLeadVisit(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, visit)

}

func TestCreateLeadVisit(t *testing.T) {
	createRandomLeadVisit(t)
}

func TestFetchLeadVisitByLead(t *testing.T) {
	lead, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")

	result, err := testQueries.FetchAllVisitByLead(context.Background(), lead)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestDeleteLeadVisit(t *testing.T) {
	lead, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")
	visit_id, _ := uuid.Parse("41deaf7a-cb20-4a47-a331-94405843267f")

	args := db.DeleteLeadVisitParams{
		LeadID: lead,
		ID:     visit_id,
	}

	result, err := testQueries.DeleteLeadVisit(context.Background(), args)
	require.NoError(t, err)

	a_rows, _ := result.RowsAffected()

	require.NotZero(t, a_rows)

}

func TestCountLeadVisits(t *testing.T) {
	lead, _ := uuid.Parse("81da42b4-9eeb-4f43-a14f-1dc33de56aef")

	count, err := testQueries.CountOfLeadVisit(context.Background(), lead)
	require.NoError(t, err)
	require.NotZero(t, count)
}
