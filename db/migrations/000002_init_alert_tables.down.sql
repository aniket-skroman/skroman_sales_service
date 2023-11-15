alter table sale_leads
    drop constraint if exists check_referal_name_and_contact;

alter table lead_info
    drop constraint if exists check_lead_info;

alter table lead_order
    drop constraint if exists check_lead_order;