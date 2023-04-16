-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "name" text NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_name_key" to table: "users"
CREATE UNIQUE INDEX "users_name_key" ON "public"."users" ("name");
