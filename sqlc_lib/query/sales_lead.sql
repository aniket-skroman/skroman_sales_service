-- name: CreateNewLead :one
insert into sale_leads (
    lead_by,
    referal_name,
    referal_contact,
    status,
    quatation_count
) values (
    $1,$2,$3,$4,$5
) returning *;

/* update a lead */
-- name: UpdateSaleLeadReferal :one
update sale_leads
set referal_name = $2,
referal_contact = $3,
updated_at = CURRENT_TIMESTAMP
where id = $1
returning *;

/* increase a quatation count */
-- name: IncreaeQuatationCount :execrows
update sale_leads
set quatation_count = quatation_count + 1,
updated_at = CURRENT_TIMESTAMP
where id = $1;