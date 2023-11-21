package querytest

import (
	"context"
	"fmt"
	"testing"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomOrderQuatation(t *testing.T) db.OrderQuatation {
	lead_id, _ := uuid.Parse("996db888-bde9-4f30-b335-0b777b70673b")

	args := db.CreateNewOrderQuatationParams{
		LeadID:        lead_id,
		GeneratedBy:   uuid.New(),
		QuatationLink: "sdfdf",
	}

	quatation, err := testQueries.CreateNewOrderQuatation(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, quatation)

	return quatation
}

func TestCreateOrderQuatation(t *testing.T) {
	createRandomOrderQuatation(t)
}

func TestFetchOrderQutationsByLeadId(t *testing.T) {
	lead_id, _ := uuid.Parse("996db888-bde9-4f30-b335-0b777b70673b")

	quatations, err := testQueries.FetchQuatationByLeadId(context.Background(), lead_id)

	require.NoError(t, err)
	require.NotEmpty(t, quatations)

	for i := range quatations {
		//path := "http://localhost:9000/api/device-file/quatation/"
		path := fmt.Sprintf("http://localhost:9000/api/device-file/%s", quatations[i].QuatationLink)
		fmt.Println("Quatation : ", path)
	}
}

func TestDeleteOrderQuotation(t *testing.T) {

}
