create table article
(
    id          int auto_increment
primary key,
    title       text          not null,
    description text          not null,
    content     text          not null,
    tag         text          not null,
    watch_count int default 0 null,
    create_time int default 0 null,
    update_time int default 0 null
);

