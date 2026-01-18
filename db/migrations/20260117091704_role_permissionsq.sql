-- +goose Up
-- +goose StatementBegin
create table if not exists role_permissions(
    id INT primary key auto_increment,
    role_id int not null,
    permission_id int not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    foreign key (role_id) references roles(id) on delete cascade,
    foreign key (permission_id) references permissions(id) on delete cascade

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table role_permissions;
-- +goose StatementEnd
