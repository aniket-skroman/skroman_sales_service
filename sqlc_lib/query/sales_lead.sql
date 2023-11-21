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
select sl.id as lead_id,
sl.lead_by as lead_by, sl.referal_name as referal_name,
sl.referal_contact as referal_contact, sl.status as status,
sl.created_at as lead_created_at, sl.updated_at as lead_updated_at,
sl.is_lead_info as is_lead_info, sl.is_order_info as is_order_info,
sl.quatation_count as quatation_count, li.id as lead_info_id,
li.name as name, li.email as email, li.contact as contact,
li.address_line_1 as address_line_1, li.city as city, li.state as state,
li.lead_type as lead_type, li.created_at as lead_info_created_at,
li.updated_at as lead_info_updated_at
from sale_leads as sl
left join lead_info as li 
on sl.id = li.lead_id
order by sl.created_at desc
limit $1
offset $2;

/* fetch lead by id */
-- name: FetchLeadByLeadId :one
select sl.id as lead_id,
sl.lead_by as lead_by, sl.referal_name as referal_name,
sl.referal_contact as referal_contact, sl.status as status,
sl.created_at as lead_created_at, sl.updated_at as lead_updated_at,
sl.is_lead_info as is_lead_info, sl.is_order_info as is_order_info,
sl.quatation_count as quatation_count, li.id as lead_info_id,
li.name as name, li.email as email, li.contact as contact,
li.address_line_1 as address_line_1, li.city as city, li.state as state,
li.lead_type as lead_type, li.created_at as lead_info_created_at,
li.updated_at as lead_info_updated_at
from sale_leads as sl
inner join lead_info as li 
on sl.id = li.lead_id
where sl.id = $1
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

-- name: CountOfLeads :one
select count(*) from sale_leads;

/* counts for landing page */
-- name: FetchLeadCounts :one
select (select count(*) from sale_leads where status = 'INIT') as init_leads,
(select count(*) from sale_leads where status = 'PLACED') as placed_leads,
(select count(*) from sale_leads where status = 'CANCLED') as cancled_leads,
(select count(*) from order_quatation where date_trunc('month', created_at) = date_trunc('month', current_date)) as total_quotations,
count(*) as total_leads     
from sale_leads;

-- name: FetchLeadCountByMonth :many
with lm as 
(
SELECT
	to_char(d, 'Month') as n_month
FROM
    GENERATE_SERIES(
        now(),
        now() - interval '12 months',
        interval '-1 months'
    ) AS d
)


select  l.n_month as month,
count(distinct sl.id)
from lm as l
left join sale_leads as sl 
on l.n_month = to_char(sl.created_at, 'Month')
group by to_char(sl.created_at, 'Month'),l.n_month
order by l.n_month desc
;


/* lead by status */
-- name: FetchLeadsByStatus :many
select sl.id as lead_id,
sl.lead_by as lead_by, sl.referal_name as referal_name,
sl.referal_contact as referal_contact, sl.status as status,
sl.created_at as lead_created_at, sl.updated_at as lead_updated_at,
sl.is_lead_info as is_lead_info, sl.is_order_info as is_order_info,
sl.quatation_count as quatation_count, li.id as lead_info_id,
li.name as name, li.email as email, li.contact as contact,
li.address_line_1 as address_line_1, li.city as city, li.state as state,
li.lead_type as lead_type, li.created_at as lead_info_created_at,
li.updated_at as lead_info_updated_at
from sale_leads as sl
inner join lead_info as li 
on sl.id = li.lead_id
where sl.status = $1
order by sl.created_at
limit $2
offset $3
;

/* lead count by status */
-- name: PGCountByLeadStatus :one
select count(*) from sale_leads as sl 
inner join lead_info as li 
on sl.id = li.lead_id 
where status = $1;