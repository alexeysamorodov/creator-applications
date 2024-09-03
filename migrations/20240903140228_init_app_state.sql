-- +goose Up
-- +goose StatementBegin
update applications
set data = jsonb_set(data, '{state}', '"Created"')
where data->>'state' is null
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
