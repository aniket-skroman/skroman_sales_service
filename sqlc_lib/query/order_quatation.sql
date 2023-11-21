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

/* delete order quotation by lead_id and quotation id */
-- name: DeleteOrderQuotation :execresult
delete from order_quatation
where id = $1 and lead_id = $2;

/* fetch quotation by quotation id */
-- name: FetchQuotationById :one
select * from order_quatation
where id = $1
limit 1;