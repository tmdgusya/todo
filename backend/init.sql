CREATE TABLE Posts (
                       id bigint not null auto_increment primary key,
                       title varchar(255) not null,
                       content text null,
                       isChecked char(1) not null,
                       createdAt timestamp default NOW(),
                       updatedAt timestamp default NOW(),
                       tagId bigint not null,
                       categoryId bigint not null
) ENGINE=InnoDB CHARACTER SET utf8;


CREATE TABLE Categories (
                       id bigint not null auto_increment primary key,
                       name varchar(255) not null unique
) ENGINE=InnoDB CHARACTER SET utf8;


CREATE TABLE Tag(
    id bigint not null auto_increment primary key,
    name varchar(255) not null unique,
    color varchar(255) not null
) ENGINE=InnoDB CHARACTER SET utf8;