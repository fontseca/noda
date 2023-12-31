CREATE TABLE IF NOT EXISTS "tag"
(
  "tag_id"      UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "owner_id"    UUID NOT NULL REFERENCES "user" ("user_id"),
  "name"        VARCHAR(50) NOT NULL UNIQUE,
  "description" VARCHAR(512) DEFAULT NULL,
  "color"       tag_color_t NOT NULL DEFAULT 'fff',
  "created_at"  TIMESTAMPTZ NOT NULL DEFAULT now (),
  "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT now ()
);

ALTER TABLE "tag"
   OWNER TO "noda";

COMMENT ON TABLE "tag"
              IS 'Labels and categorizes enhance organization and searchability.';
