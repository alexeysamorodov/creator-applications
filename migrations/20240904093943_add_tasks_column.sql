-- +goose Up
-- +goose StatementBegin
alter table applications add tasks jsonb
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
