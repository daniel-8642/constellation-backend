create table session
(
    uid      int auto_increment,
    session  varchar(20) not null,
    lasttime datetime    not null,
    constraint session_session_uindex
        unique (session),
    constraint session_uid_uindex
        unique (uid)
);

alter table session
    add primary key (uid);

create table starLog
(
    consName varchar(10) null,
    ip       varchar(20) null,
    time     date        not null
);

create table user
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

alter table user
    add primary key (uid);

