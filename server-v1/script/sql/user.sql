use website;

# account表，保存的是账号信息，该表包含用户最基本的涉及安全的信息，很少会变动
create table account
(
    id       bigint unsigned not null unique auto_increment comment '用户ID',
    account  char(8)         not null unique default '' comment '账号',
    password char(32)        not null        default '' comment '密码',
    phone    char(11)        not null unique default '' comment '手机号',
    email    varchar(64)     not null unique default '' comment '邮箱',
    primary key (id)
) comment '账号表' charset = utf8
                   engine = InnoDB;


insert into account (account, password, phone, email)
values ('admin', 'admin', '12345678901',
        '@email.com');

select * from account where id = 1;

# 账号信息表，保存的是用户的个人信息，该表包含用户的个人信息，会经常变动
create table account_info
(
    id        bigint unsigned  not null default 0 comment '外键，用户ID',
    nickname  char(32)         not null default '' comment '昵称',
    age       tinyint unsigned not null default 0 comment '年龄',
    gender    tinyint unsigned not null default 0 comment '性别, 0: 未知, 1: 男, 2: 女',
    campus    char(32)         not null default '' comment '校区',
    dormitory char(32)         not null default '' comment '宿舍',
    primary key (id)
) comment '账号信息表' charset = utf8
                       engine = InnoDB;

insert into account_info (id, nickname, age, gender, campus, dormitory)
values (1, 'admin nickname', 20, 1, '南湖', '1号楼');