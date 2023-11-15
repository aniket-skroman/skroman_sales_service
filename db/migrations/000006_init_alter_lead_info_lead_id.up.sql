alter table lead_info
    add constraint check_lead_name check (name <> ''),
    add constraint check_lead_contact check (contact <> '' and length(contact)=10);    
