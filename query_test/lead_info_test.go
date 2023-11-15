package querytest

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	db "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomLeadInfoTest(t *testing.T, lead_id uuid.UUID) {

	args := db.CreateLeadInfoParams{
		LeadID: uuid.NullUUID{UUID: lead_id, Valid: true},
		Name:   "test",
		Email: sql.NullString{
			String: "test@gmail.com",
			Valid:  true,
		},
		Contact: "123456789",
		AddressLine1: sql.NullString{
			String: "My Address",
			Valid:  true,
		},
		City: sql.NullString{
			String: "PUNE",
			Valid:  true,
		},
		State: sql.NullString{
			String: "Maharashtra",
			Valid:  true,
		},
		LeadType: sql.NullString{
			String: "Builder",
			Valid:  true,
		},
	}

	lead_info, err := testQueries.CreateLeadInfo(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, lead_info)
}

func TestCreateLeadInfo(t *testing.T) {
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")
	lead_id_1, _ := uuid.Parse("0745cc17-8d9f-4356-a58d-ed29976a6355")
	lead_id_2, _ := uuid.Parse("3ce2b1f8-a738-48da-bb9f-59a664d663d8")

	tests := []struct {
		TestName    string
		ExpectedErr bool
		QueryArgs   db.CreateLeadInfoParams
	}{
		{
			TestName:    "First Test",
			ExpectedErr: true,
			QueryArgs: db.CreateLeadInfoParams{
				LeadID:       uuid.NullUUID{UUID: uuid.New(), Valid: true},
				Name:         "test",
				Email:        sql.NullString{String: "test@gmail.com", Valid: true},
				Contact:      "123456789",
				AddressLine1: sql.NullString{String: "My Address", Valid: true},
				City:         sql.NullString{String: "PUNE", Valid: true},
				State:        sql.NullString{String: "Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "Builder", Valid: true},
			},
		},
		{
			TestName:    "Secound Test",
			ExpectedErr: false,
			QueryArgs: db.CreateLeadInfoParams{
				LeadID:       uuid.NullUUID{UUID: lead_id, Valid: true},
				Name:         "test",
				Email:        sql.NullString{String: "test@gmail.com", Valid: true},
				Contact:      "1234567890",
				AddressLine1: sql.NullString{String: "My Address", Valid: true},
				City:         sql.NullString{String: "PUNE", Valid: true},
				State:        sql.NullString{String: "Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "Builder", Valid: true},
			},
		},
		{
			TestName:    "Third Test",
			ExpectedErr: true,
			QueryArgs: db.CreateLeadInfoParams{
				LeadID:       uuid.NullUUID{UUID: lead_id_1, Valid: true},
				Name:         "",
				Email:        sql.NullString{String: "test@gmail.com", Valid: true},
				Contact:      "123456789",
				AddressLine1: sql.NullString{String: "My Address", Valid: true},
				City:         sql.NullString{String: "PUNE", Valid: true},
				State:        sql.NullString{String: "Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "Builder", Valid: true},
			},
		},
		{
			TestName:    "Fourth Test",
			ExpectedErr: true,
			QueryArgs: db.CreateLeadInfoParams{
				LeadID:       uuid.NullUUID{UUID: lead_id_2, Valid: true},
				Name:         "test",
				Email:        sql.NullString{String: "test@gmail.com", Valid: true},
				Contact:      "",
				AddressLine1: sql.NullString{String: "My Address", Valid: true},
				City:         sql.NullString{String: "PUNE", Valid: true},
				State:        sql.NullString{String: "Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "Builder", Valid: true},
			},
		},
	}

	for i := range tests {
		t.Run(tests[i].TestName, func(t *testing.T) {

			result, err := testQueries.CreateLeadInfo(context.Background(), tests[i].QueryArgs)

			if tests[i].ExpectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, result)
			}
		})
	}
}

func TestFetchLeadInfoByLeadId(t *testing.T) {
	lead_id, _ := uuid.Parse("a1aa3ce2-6f3a-459c-9f27-67b1798bdd45")

	tests := []struct {
		TestName    string
		ExpectedErr bool
		LeadId      uuid.UUID
		Error       error
	}{
		{
			TestName:    "First Test",
			ExpectedErr: true,
			Error:       sql.ErrNoRows,
			LeadId:      uuid.New(),
		},
		{
			TestName:    "Second Test",
			ExpectedErr: false,
			LeadId:      lead_id,
		},
	}

	for i := range tests {
		t.Run(tests[i].TestName, func(t *testing.T) {
			lead_info, err := testQueries.FetchLeadInfoByLeadID(context.Background(), uuid.NullUUID{UUID: tests[i].LeadId, Valid: true})

			if tests[i].ExpectedErr {
				require.Error(t, err)
				require.Equal(t, err, sql.ErrNoRows)
				require.Empty(t, lead_info)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateLeadInfoById(t *testing.T) {
	lead_id, _ := uuid.Parse("b057fcd2-12ce-4d09-a669-ff261ae61839")
	tests := []struct {
		TestName    string
		ExpectedErr bool
		QueryArgs   db.UpdateLeadInfoParams
		Error       error
	}{
		{
			TestName:    "First Test",
			ExpectedErr: true,
			QueryArgs: db.UpdateLeadInfoParams{
				ID:           uuid.New(),
				Name:         "New Update Name",
				Email:        sql.NullString{String: "updatetest@gmail.com", Valid: true},
				Contact:      "123456789",
				AddressLine1: sql.NullString{String: "update My Address", Valid: true},
				City:         sql.NullString{String: " update PUNE", Valid: true},
				State:        sql.NullString{String: "update Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "update Builder", Valid: true},
			},
			Error: sql.ErrNoRows,
		},
		{
			TestName:    "Second Test",
			ExpectedErr: true,
			QueryArgs: db.UpdateLeadInfoParams{
				ID:           lead_id,
				Name:         "New Update Name",
				Email:        sql.NullString{String: "updatetest@gmail.com", Valid: true},
				Contact:      "123456789",
				AddressLine1: sql.NullString{String: "update My Address", Valid: true},
				City:         sql.NullString{String: " update PUNE", Valid: true},
				State:        sql.NullString{String: "update Maharashtra", Valid: true},
				LeadType:     sql.NullString{String: "update Builder", Valid: true},
			},
			Error: errors.New("new row for relation lead_info violates check constraint check_lead_contact"),
		},
	}

	for i := range tests {

		t.Run(tests[i].TestName, func(t *testing.T) {

			lead_info, err := testQueries.UpdateLeadInfo(context.Background(), tests[i].QueryArgs)

			if tests[i].ExpectedErr {
				require.Error(t, err)
				//require.Equal(t, err, tests[i].Error)
			} else {
				require.NoError(t, err)
				require.Equal(t, lead_info.ID, tests[i].QueryArgs.ID)
			}
		})
	}
}
