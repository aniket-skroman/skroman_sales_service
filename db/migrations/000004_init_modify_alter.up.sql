alter table sale_leads
    add constraint check_ref_name check (referal_name <> ''),
    add constraint check_ref_cont check (referal_contact <> '');

-- alter table lead_info
--     add constraint check_lead_info check (name <> ''),
--     add constraint check_lead_cont check (contact <> '');

-- alter table lead_order
--     add constraint check_device_type check (device_type <> ''),
--     add constraint check_device_name check (device_name <> ''),
--     add constraint check_device_price check (device_price <> ''),
--     add constraint check_device_device_model check (device_device_model <> '');
