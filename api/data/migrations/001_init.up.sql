CREATE TABLE "users" (
                         "id" bigint PRIMARY KEY NOT NULL,
                         "username" varchar(50) UNIQUE NOT NULL,
                         "password" varchar(50) NOT NULL,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "cards" (
                         "id" bigint PRIMARY KEY NOT NULL,
                         "number" varchar(50) NOT NULL,
                         "expired_date" date,
                         "user_id" bigint NOT NULL,
                         "cvv" varchar(20),
                         "balance" bigint,
                         "deleted_at" timestamp,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "audit_trail" (
                               "id" bigint PRIMARY KEY NOT NULL,
                               "user_id" bigint NOT NULL,
                               "order_id" bigint NOT NULL,
                               "status" varchar(50),
                               "created_at" timestamp,
                               "updated_at" timestamp
);

CREATE TABLE "orders" (
                          "id" bigint PRIMARY KEY NOT NULL,
                          "amount" bigint NOT NULL,
                          "deleted_at" timestamp,
                          "created_at" timestamp,
                          "updated_at" timestamp
);

CREATE TABLE "transactions" (
                                "id" bigint PRIMARY KEY NOT NULL,
                                "order_id" bigint NOT NULL,
                                "card_id" bigint NOT NULL,
                                "otp" varchar(20),
                                "status" varchar(50),
                                "created_at" timestamp,
                                "updated_at" timestamp
);

ALTER TABLE "cards" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "audit_trail" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "audit_trail" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("id");
