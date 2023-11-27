alter table cancel_leads
    add constraint check_cancel_reason check (reason <> '');