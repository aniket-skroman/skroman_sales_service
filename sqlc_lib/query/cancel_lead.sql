/* cancel lead */
-- name: CreateCancelLead :one
insert into cancel_leads (
    reason,
    cancel_by,
    lead_id
) values (
    $1, $2, $3
) returning *;

-- name: FetchCancelLeadByLeadId :one
select * from cancel_leads
where lead_id = $1
limit 1;
