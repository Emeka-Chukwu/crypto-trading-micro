
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "auths" (
  "id" uuid DEFAULT uuid_generate_v4(),
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar NOT NULL UNIQUE,
  "active" boolean NOT NULL DEFAULT false,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

