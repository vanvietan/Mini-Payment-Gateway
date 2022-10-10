CREATE TABLE "users" (
                         "id" bigint PRIMARY KEY NOT NULL,
                         "username" varchar(50) UNIQUE NOT NULL,
                         "password" varchar(50) NOT NULL,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "cards" (
                         "id" bigint UNIQUE PRIMARY KEY NOT NULL,
                         "expired_date" date,
                         "user_id" bigint NOT NULL,
                         "cvv" int,
                         "balance" int,
                         "deleted_at" timestamp,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "audit_trail" (
                               "id" bigint PRIMARY KEY NOT NULL,
                               "user_id" bigint NOT NULL,
                               "card_id" bigint NOT NULL,
                               "status" varchar(50),
                               "created_at" timestamp,
                               "updated_at" timestamp
);

ALTER TABLE "cards" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "audit_trail" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "audit_trail" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("id");
