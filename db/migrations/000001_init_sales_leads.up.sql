
CREATE TABLE "sale_leads" (
  "id" uuid default uuid_generate_v4 () primary key,
  "lead_by" uuid NOT NULL,
  "referal_name" varchar NOT NULL,
  "referal_contact" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'INIT',
  "quatation_count" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE "lead_info" (
  "id" uuid default uuid_generate_v4 () primary key,
  "lead_id" uuid UNIQUE,  
  "name" varchar NOT NULL,
  "email" varchar,
  "contact" varchar NOT NULL,
  "address_line_1" varchar,
  "city" varchar,
  "state" varchar,
  "lead_type" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  foreign key ("lead_id") references sale_leads("id")
);


CREATE TABLE "lead_order" (
  "id" uuid default uuid_generate_v4 () primary key,
  "lead_id" uuid,
  "device_type" varchar,
  "device_model" varchar,
  "device_name" varchar,
  "device_price" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  foreign key ("lead_id") references sale_leads("id")
);



