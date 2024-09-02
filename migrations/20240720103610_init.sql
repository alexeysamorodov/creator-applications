-- +goose Up
-- +goose StatementBegin
create table applications
(
    id               uuid      not null PRIMARY KEY,
    external_id      bigint    not null unique,
    created_at       timestamp not null,
    updated_at       timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists applications;
-- +goose StatementEnd
