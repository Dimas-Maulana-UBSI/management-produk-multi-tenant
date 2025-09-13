create table tenants (
    id int(11) auto_increment primary key,
    name varchar(100),
    api_key varchar(100),
    db_port varchar(100),
    db_host int(11),
    db_name varchar(100),
    db_user varchar(100),
    db_password varchar(100),
    status varchar(100)
);