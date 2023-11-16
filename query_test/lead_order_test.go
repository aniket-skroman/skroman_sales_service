package querytest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomLeadOrder(t *testing.T) {
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")
	args := db.CreateLeadOrderParams{
		LeadID:      uuid.NullUUID{UUID: lead_id, Valid: true},
		DeviceType:  sql.NullString{String: "Test Device", Valid: true},
		DeviceModel: sql.NullString{String: "Test Device Model", Valid: true},
		DevicePrice: sql.NullInt32{Int32: 123455, Valid: true},
		DeviceName:  sql.NullString{String: "test Device Name", Valid: true},
	}

	result, err := testQueries.CreateLeadOrder(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestCreateLeadOrder(t *testing.T) {
	createRandomLeadOrder(t)
}

func TestFetchLeadOrders(t *testing.T) {
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")

	orders, err := testQueries.FetchOrdersByLeadId(context.Background(), uuid.NullUUID{UUID: lead_id, Valid: true})

	require.NoError(t, err)
	require.NotEmpty(t, orders)
	fmt.Println("LEN Of Orders : ", len(orders))
}

func TestDeleteLeadOrder(t *testing.T) {
	order_id, _ := uuid.Parse("17b2154f-f615-48d1-867d-1dd41bf935df")
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")

	args := db.DeleteLeadOrderParams{
		ID:     order_id,
		LeadID: uuid.NullUUID{UUID: lead_id, Valid: true},
	}

	result, err := testQueries.DeleteLeadOrder(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	affectedRows, err := result.RowsAffected()

	require.NoError(t, err)
	fmt.Println("Affected Rows : ", affectedRows)
	require.NotZero(t, affectedRows)
}

func TestUpdateLeadOrder(t *testing.T) {
	order_id, _ := uuid.Parse("a9f2034c-4b69-4008-89a3-8cf9bf48eef5")
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")

	args := db.UpdateLeadOrderParams{
		ID:          order_id,
		LeadID:      uuid.NullUUID{UUID: lead_id, Valid: true},
		DeviceType:  sql.NullString{String: "Test Device Update", Valid: true},
		DeviceModel: sql.NullString{String: "Test Device Model Update", Valid: true},
		DevicePrice: sql.NullInt32{Int32: 123455, Valid: true},
	}

	order, err := testQueries.UpdateLeadOrder(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, order)
}

func TestFetchLeadOrdersByOrderId(t *testing.T) {
	order_id, _ := uuid.Parse("a9f2034c-4b69-4008-89a3-8cf9bf48eef0")

	order, err := testQueries.FetchLeadOrderByOrderId(context.Background(), order_id)

	require.NoError(t, err)
	require.NotEmpty(t, order)

}
