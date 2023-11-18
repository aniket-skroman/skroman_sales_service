-- name: CreateLeadOrder :one
insert into lead_order (
    lead_id,
    device_type,
    device_model,
    device_price,
    device_name,
    quantity
) values (
    $1,$2,$3,$4,$5,$6
) returning *;


/* fetch all order by lead id */
-- name: FetchOrdersByLeadId :many
select * from lead_order
where lead_id = $1
order by created_at desc;


/* delete a order by order id */
-- name: DeleteLeadOrder :execresult 
delete from lead_order
where id = $1 and lead_id = $2;

/* update a specific order */
-- name: UpdateLeadOrder :one
update lead_order
set device_type = $3,
device_model = $4,
device_price = $5,
updated_at = CURRENT_TIMESTAMP
where id = $1 and lead_id = $2
returning *;

/* fetch lead order by order id*/
-- name: FetchLeadOrderByOrderId :one
select * from lead_order
where id = $1
limit 1
;

/* check is there any order or all order get deleted */
-- name: CheckLeadHasOrder :one
select count(*) from lead_order
where lead_id = $1;