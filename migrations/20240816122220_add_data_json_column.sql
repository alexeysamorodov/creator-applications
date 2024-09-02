-- +goose Up
-- +goose StatementBegin
alter table applications add data jsonb
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
