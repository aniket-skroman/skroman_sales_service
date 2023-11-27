create table "cancel_leads" (
    "id" uuid default uuid_generate_v4 () primary key,
    "cancel_by" uuid not null,
    "lead_id" uuid not null unique,
    "reason" text not null,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    foreign key ("lead_id") references sale_leads("id")
);