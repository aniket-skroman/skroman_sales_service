-- alter table sale_leads
--     add column is_order_placed boolean;

alter table sale_leads 
    alter COLUMN is_order_placed SET DEFAULT False;
