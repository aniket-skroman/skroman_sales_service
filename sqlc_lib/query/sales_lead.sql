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

/* fetch all leads */
-- name: FetchAllLeads :many
select * from sale_leads
order by created_at desc
limit $1
offset $2;

/* fetch lead by id */
-- name: FetchLeadByLeadId :one
select * from sale_leads
where id = $1
limit 1;

/* make a flags true or false for info or orde */
-- name: UpdateIsLeadInfo :one
update sale_leads
set is_lead_info = $2
where id = $1
returning *;

-- name: UpdateIsLeadOrder :one
update sale_leads
set is_order_info = $2
where id = $1
returning *;