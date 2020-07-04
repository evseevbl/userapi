-- +goose Up
-- +goose StatementBegin
alter table "user"
    add column phone varchar(10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table "user"
    drop column phone;
-- +goose StatementEnd
