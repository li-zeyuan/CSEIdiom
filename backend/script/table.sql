create table if not exists user_profile
(
    created_at         timestamp             default CURRENT_TIMESTAMP,
    updated_at         timestamp             default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    deleted_at         bigint                default 0,
    id                 bigint PRIMARY KEY,
    name               varchar(64)  not null default '',
    gender             tinyint      not null default 0,
    portrait           varchar(255) not null default '',
    openid             varchar(128) not null default '',
    session_key        varchar(128) not null default '',
    current_subject_id bigint       not null default 0,
    study_total        int          not null default 0,
    study_unm          int          not null default 0,
    unique (openid)
    ) ENGINE = INNODB;