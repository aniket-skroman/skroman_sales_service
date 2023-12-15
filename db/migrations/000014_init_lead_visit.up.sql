create table lead_visit(
    "id" uuid default uuid_generate_v4 () primary key,
    "lead_id" uuid not null,
    "visit_by" uuid not null,
    "visit_discussion" text not null,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    foreign key ("lead_id") references sale_leads("id")
);