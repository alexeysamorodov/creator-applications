-- +goose Up
-- +goose StatementBegin
create table applications
(
    id               uuid      not null PRIMARY KEY,
    external_id      bigint    not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists applications;
-- +goose StatementEnd
