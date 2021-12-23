create table session
(
    uid      int auto_increment
        primary key,
    session  varchar(20) not null,
    lasttime datetime    not null,
    constraint session_session_uindex
        unique (session),
    constraint session_uid_uindex
        unique (uid)
);

create table starLog
(
    id       int auto_increment
        primary key,
    consName varchar(10) null,
    ip       varchar(20) null,
    time     date        not null,
    constraint starLog_id_uindex
        unique (id)
);

create table user
(
    uid   bigint auto_increment
        primary key,
    uname varchar(20) not null,
    upass varchar(64) not null,
    uauth int         not null,
    constraint user_uid_uindex
        unique (uid),
    constraint user_uname_uindex
        unique (uname)
);