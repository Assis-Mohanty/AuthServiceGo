-- +goose Up
-- +goose StatementBegin
INSERT INTO permissions(name,description,resource,action) VALUES
('user:read','Permission to read user data','user','read'),
('user:write','Permission to write user data','user','write'),
('user:delete','Permission to delete user data','user','delete'),
('role:read','Permission to read role data','role','read'),
('role:write','Permission to write role data','role','write'),
('role:delete','Permission to delete role data','role','delete'),
('role:manage','Permission to manage role assignments','role','manage'),
('permission:read','Permission to read permission data','permission','read'),
('permission:write','Permission to write permission data','permission','write'),
('permission:delete','Permission to delete permission data','permission','delete'),
('permission:manage','Permission to manage permissions','permission','manage')
on duplicate key update description=values(description),resource=values(resource),action=values(action);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
