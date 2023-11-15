alter table sale_leads
    add constraint check_referal_name_and_contact check(referal_name is not null and referal_contact is not null);

alter table lead_info
    add constraint check_lead_info check(name is not null and contact is not null);

alter table lead_order
    add constraint check_lead_order check(device_type is not null and device_name is not null and device_price is not null and device_model is not null);