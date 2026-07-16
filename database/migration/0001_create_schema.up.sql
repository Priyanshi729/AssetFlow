BEGIN;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE asset_type AS ENUM (
    'laptop',
    'keyboard',
    'mouse',
    'mobile'
    );

CREATE TYPE asset_status AS ENUM (
    'available',
    'assigned',
    'in_service',
    'for_repair',
    'damaged'
    );

CREATE TYPE user_role AS ENUM (
    'admin',
    'employee',
    'project-manager'
    );

CREATE TYPE user_type AS ENUM (
    'full-time',
    'intern',
    'freelancer'
    );

CREATE TYPE owner_type AS ENUM (
    'client',
    'remotestate'
    );

CREATE TYPE connection_type AS ENUM (
    'wired',
    'wireless'
    );

CREATE TABLE IF NOT EXISTS users (
                                     user_id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     name          TEXT NOT NULL,
                                     email         TEXT NOT NULL,
                                     phone_no      TEXT NOT NULL,
                                     role          user_role DEFAULT 'employee',
                                     user_type     user_type NOT NULL,
                                     password      TEXT NOT NULL,
                                     created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     updated_at    TIMESTAMPTZ ,
                                     archived_at   TIMESTAMPTZ
);

CREATE UNIQUE INDEX idx_unique_email
    ON users (email)
    WHERE archived_at IS NULL;

CREATE TABLE IF NOT EXISTS assets (
                        asset_id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        brand           TEXT          NOT NULL,
                        model           TEXT          NOT NULL,
                        serial_number       TEXT      NOT NULL,
                        asset_type            asset_type    NOT NULL,
                        status          asset_status  DEFAULT 'available',
                        owner_type           owner_type   DEFAULT 'remotestate',

                        warranty_start  TIMESTAMPTZ   NOT NULL,
                        warranty_end    TIMESTAMPTZ   NOT NULL,



                        created_at      TIMESTAMPTZ   DEFAULT CURRENT_TIMESTAMP,
                        updated_at      TIMESTAMPTZ,
                        archived_at     TIMESTAMPTZ

);



CREATE TABLE laptops (
                         laptop_id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         asset_id         UUID NOT NULL UNIQUE REFERENCES assets(asset_id),
                         processor        TEXT,
                         ram              TEXT,
                         storage          TEXT,
                         operating_system TEXT,
                         charger           BOOLEAN,
                         device_password  TEXT NOT NULL
);

CREATE TABLE keyboards (
                           keyboard_id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           asset_id     UUID NOT NULL UNIQUE REFERENCES assets(asset_id),
                           layout       TEXT,
                           connectivity connection_type
);

CREATE TABLE mouses (
                        mouse_id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        asset_id     UUID NOT NULL UNIQUE REFERENCES assets(asset_id),
                        dpi          INT,
                        connectivity connection_type
);

CREATE TABLE mobiles (
                         mobile_id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         asset_id         UUID NOT NULL UNIQUE REFERENCES assets(asset_id),
                         operating_system               TEXT NOT NULL,
                         ram              TEXT NOT NULL,
                         storage          TEXT NOT NULL,
                         charger          BOOLEAN,
                         device_password  TEXT NOT NULL
);

CREATE TABLE asset_assignments (
                                   assignment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                   asset_id UUID NOT NULL REFERENCES assets (asset_id),
                                   assigned_to UUID NOT NULL REFERENCES users (user_id),
                                   assigned_on TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                   returned_at TIMESTAMPTZ,

                                   created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMPTZ,
                                   archived_at TIMESTAMPTZ
);

CREATE TABLE asset_repairs (
    repair_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES assets (asset_id),
    sent_for_repair_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    repaired_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    archived_at TIMESTAMPTZ
    );

COMMIT;
