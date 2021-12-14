create table starWeb.starLog
(
    consName varchar(10) null,
    ip       varchar(20) null,
    time     datetime    not null
);

create table starWeb.user
(
    uid   bigint auto_increment,
    uname varchar(20) not null,
    upass varchar(64) not null,
    uauth int         not null,
    constraint user_uid_uindex
        unique (uid),
    constraint user_uname_uindex
        unique (uname)
);

alter table starWeb.user
    add primary key (uid);