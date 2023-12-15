-- name: CreateLeadVisit :one
insert into lead_visit (
    lead_id,
    visit_by,
    visit_discussion
) values (
    $1,$2,$3
) returning *;

-- name: FetchAllVisitByLead :many
select * from lead_visit
where lead_id = $1
order by created_at desc;

-- name: DeleteLeadVisit :execresult
delete from lead_visit
where lead_id = $1 and id = $2;

-- name: CountOfLeadVisit :one
select count(*) from lead_visit
where lead_id = $1;