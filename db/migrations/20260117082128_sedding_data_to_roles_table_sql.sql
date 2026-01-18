-- +goose Up
-- +goose StatementBegin
INSERT INTO roles (name,description) VALUES ('admin','Administator with full access'),('user','User with limited access'),('moderator','Moderator with elivated previlages');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
