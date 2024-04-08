use website;

# account表，保存的是账号信息，该表包含用户最基本的涉及安全的信息，很少会变动
drop table if exists account;
create table account
(
    id           bigint unsigned  not null unique auto_increment comment '用户ID',
    account      char(8)          not null unique default '' comment '账号',
    password     char(64)         not null        default '' comment '密码，用sha256加密',
    phone        char(11)         not null unique default '' comment '手机号',
    email        varchar(64)      not null unique default '' comment '邮箱',
    nickname     char(32)         not null        default '' comment '昵称',
    age          tinyint unsigned not null        default 0 comment '年龄',
    gender       tinyint unsigned not null        default 0 comment '性别, 0: 未知, 1: 男, 2: 女',
    campus       char(32)         not null        default '' comment '校区',
    dormitory    char(32)         not null        default '' comment '宿舍',
    is_deleted   tinyint unsigned not null        default 0 comment '是否删除, 0: 未删除, 1: 已删除',
    gmt_create   datetime         not null        default current_timestamp comment '创建时间',
    gmt_modified datetime         not null        default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id)
) comment '账号表' charset = utf8
                   engine = InnoDB;

insert into account (account, password, phone, email)
values ('admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '12345678901',
        '@email.com');

SELECT @@global.time_zone;
select @@session.time_zone;