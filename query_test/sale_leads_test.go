package querytest

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSalesLead(t *testing.T) db.SaleLeads {

	args := db.CreateNewLeadParams{
		LeadBy:         uuid.New(),
		ReferalName:    "test 1",
		ReferalContact: "123045678",
		QuatationCount: sql.NullInt32{Int32: 1, Valid: true},
	}

	sale_lead, err := testQueries.CreateNewLead(context.Background(), args)
	fmt.Println(err)
	require.NoError(t, err)
	require.NotEmpty(t, sale_lead)

	return sale_lead
}

func TestCreatSaleLead(t *testing.T) {
	createRandomSalesLead(t)
}

func TestUpdateSaleLeadReferal(t *testing.T) {

	lead := createRandomSalesLead(t)

	print("Created Referal Name : ", lead.ReferalName)
	print("Created Referal Contact : ", lead.ReferalContact)

	args := db.UpdateSaleLeadReferalParams{
		ReferalName:    "eerr",
		ReferalContact: "456321789",
		ID:             lead.ID,
	}

	updated_lead, err := testQueries.UpdateSaleLeadReferal(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updated_lead)

	require.NotEqual(t, updated_lead.ReferalName, lead.ReferalName)
	require.NotEqual(t, updated_lead.ReferalContact, lead.ReferalContact)
	require.WithinDuration(t, lead.UpdatedAt, lead.CreatedAt, time.Second)
}

func TestIncreaseQuataionCount(t *testing.T) {
	lead_id, _ := uuid.Parse("8c9880d8-8633-4820-8771-dffd959ceb4a")
	dummy_id, _ := uuid.Parse("bjbdbfdjbfbfdfndjkhbf dbf")
	args := []struct {
		TestName    string
		LeadId      uuid.UUID
		ExpectedErr bool
	}{
		{
			TestName:    "First test",
			LeadId:      uuid.New(),
			ExpectedErr: true,
		},
		{
			TestName:    "Second Test",
			LeadId:      lead_id,
			ExpectedErr: false,
		},
		{
			TestName:    "Third Test",
			LeadId:      dummy_id,
			ExpectedErr: true,
		},
	}

	for i := range args {
		t.Run(args[i].TestName, func(t *testing.T) {
			_, err := testQueries.IncreaeQuatationCount(context.Background(), args[i].LeadId)

			if !args[i].ExpectedErr {
				require.NoError(t, err)
			}
			print("Error : ", err)
			//print("result for test :  ", args[i].TestName, " and result is : ", result)
		})
	}
}

func print(args ...interface{}) {
	fmt.Println("DEBUG : ", args)
}
