create table operator_tab
(
    op_id   bigint auto_increment comment '主键',
    email   varchar(255) DEFAULT "" not null comment 'email',
    pwd     varchar(255) DEFAULT "" not null comment '密码md5',
    status  int(8) DEFAULT 0 not null comment '状态 0-无效 1-有效',
    ctime   varchar(255) DEFAULT "" not null comment '创建时间',
    mtime   varchar(255) DEFAULT "" not null on update CURRENT_TIMESTAMP comment '最后修改时间',
    name    varchar(100) DEFAULT "" not null comment '姓名',
    phone   varchar(100) DEFAULT "" not null comment '电话',
    op_type tinyint(4) unsigned default 0 DEFAULT 0 not null comment '操作者类型 0-普通 1-admin',
    PRIMARY KEY (`op_id`)
)

