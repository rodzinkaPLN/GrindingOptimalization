-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE datasets (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR,
    created_at timestamptz
);

CREATE TABLE parameters(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR,
    dataset_id uuid references datasets(id),
    unit VARCHAR,
    created_at timestamptz
);

CREATE TABLE datapoints (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    data_time timestamptz,
    dataset_id uuid references datasets(id),
    vals JSONB,
    created_at timestamptz
);

CREATE TABLE predictions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    data_time timestamptz,
    vals JSONB,
    created_at timestamptz
);

CREATE TABLE userinputs (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    min double precision,
    max double precision,
    step double precision,
    defaultval double precision,
    name varchar
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

DELETE TABLE datapoints;

DELETE TABLE parameters;

DELETE TABLE datasets;

-- +goose StatementEnd