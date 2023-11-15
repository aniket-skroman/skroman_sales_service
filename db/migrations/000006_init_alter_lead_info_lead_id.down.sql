alter table lead_info
    drop constraint if exists check_lead_name,
    drop constraint if exists check_lead_contact;