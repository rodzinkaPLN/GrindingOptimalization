-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE datasets (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR,
    created_at timestamptz
);

CREATE TABLE datapoints (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    data_time timestamptz,
    val double precision,
    dataset_id uuid references datasets(id),
    created_at timestamptz
);

-- CREATE TABLE prediction_runs (
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     name VARCHAR,
--     created_at timestamptz
-- );
--
-- CREATE TABLE predictions (
--     id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
--     data_time timestamptz,
--     val double precision,
--     dataset_id uuid references datasets(id),
--     created_at timestamptz
-- );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

DELETE TABLE datasets;

DELETE TABLE datapoints;

-- +goose StatementEnd