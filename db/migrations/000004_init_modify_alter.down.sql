alter table sale_leads
    drop constraint if exists check_ref_name,
    drop constraint if exists check_ref_cont;

alter table lead_info
    drop constraint if exists check_lead_info,
    drop constraint if exists check_lead_cont;

alter table lead_order
    drop constraint if exists check_device_type,
    drop constraint if exists check_device_name,
    drop constraint if exists check_device_price,
    drop constraint if exists check_device_device_model;
