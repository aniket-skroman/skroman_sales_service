/* create a new order quatation */
-- name: CreateNewOrderQuatation :one
insert into order_quatation(
    lead_id,
    generated_by,
    quatation_link   
) values (
    $1,$2,$3
) returning *;

/* fetch order quatation by lead id */
-- name: FetchQuatationByLeadId :many
select * from order_quatation
where lead_id = $1;