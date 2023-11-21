create table "order_quatation"(
    "id" uuid default uuid_generate_v4 () primary key,
    "lead_id" uuid not null,
    "generated_by" uuid not null,
    "quatation_link" varchar not null,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    foreign key ("lead_id") references sale_leads("id")
);