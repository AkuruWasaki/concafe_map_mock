use concafe_map_db;
create table shops (
  id bigint NOT NULL,
  name varchar(255) NOT NULL,
  tel varchar(32),
  address varchar(255) NOT NULL,
  created_at datetime default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);