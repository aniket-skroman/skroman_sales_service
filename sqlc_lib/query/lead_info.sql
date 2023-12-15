/* create lead info */
-- name: CreateLeadInfo :one
insert into lead_info (
    lead_id,
    name,
    email,
    contact,
    address_line_1,
    city,
    state,
    lead_type
) values (
    $1,$2,$3,$4,$5,$6,$7,$8
) returning *;

-- name: FetchLeadInfoByLeadID :one
select * from lead_info
where lead_id = $1;

-- name: UpdateLeadInfo :one
update lead_info
set name = $2,
email = $3, contact=$4,
address_line_1=$5, city=$6,
state=$7, lead_type=$8,
updated_at = CURRENT_TIMESTAMP
where id = $1
returning *;

-- name: DeleteLeadInfoByLeadId :execresult
delete from lead_info
where lead_id = $1;

/*9022811746*/