-- +goose Up
-- +goose StatementBegin
INSERT INTO role_permissions(role_id,permission_id) 
SELECT 1,id from permissions;

-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO role_permissions(role_id,permission_id) 
SELECT 2,id from permissions where name IN ("user:read");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
