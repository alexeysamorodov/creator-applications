-- +goose Up
-- +goose StatementBegin
create table applications
(
    id               uuid      not null PRIMARY KEY,
    external_id      bigint    not null unique,
    tax_id           text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists applications;
-- +goose StatementEnd
