alter table sale_leads
    add column is_lead_info boolean default 'FALSE',
    add column is_order_info boolean default 'FALSE';