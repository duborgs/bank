drop database if exists q2bank;

create database q2bank;

use q2bank;

create table users(
    id      int primary key not null auto_increment,
    name_user    varchar(50) not null,
    type_user    enum('commom', 'shopkeeper'),
    CPF_CNPJ       varchar(50) unique,
    email   varchar(50) unique not null,
    password_user   varchar(50) not null,
    wallet   FLOAT(10.2) not null
);

create table transactions(
    id int primary key auto_increment,
    amount FLOAT (10.2) not null,
    id_origin   int not null,
    id_dest int not null,
    date_time datetime not null
);

alter table transactions
add foreign key (id_origin) references users(id);

alter table transactions
add foreign key (id_dest) references users(id);

insert into users (name_user, type_user, CPF_CNPJ, email, password_user, wallet) values ('Anderson Amaral Santos ', 'commom', '123.123.123-90','anderson@gmail.com', '123456', 1000.10);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, wallet) values ('Maria Santa Cruz', 'commom', '444.123.123-90', 'maria@gmail.com', '444444', 800.10);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, wallet) values ('Roberto Firmino da Silva', 'shopkeeper', '444.123.988-90', 'roberto@hotmail.com', 'robertinho', 3000);
insert into users (name_user, type_user, CPF_CNPJ, email, password_user, wallet) values ("Junior da Silva Ferreira", "shopkeeper", "825.510.380-90", "juninho123@", "juninho2022@outlook.com", 100000.00);