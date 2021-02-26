create database urlshoten;

create table urls(id varchar(50) not null primary key unique, shorturl varchar(50) not null, longurl varchar(150) not null unique);