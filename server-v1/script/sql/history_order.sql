use website;

drop table if exists history_order;

create table history_order
(
    id           bigint unsigned  not null unique auto_increment comment '用户ID',
    product_id   bigint unsigned  not null default 0 comment '商品ID',
    buyer_id     bigint unsigned  not null default 0 comment '买家ID',
    seller_id    bigint unsigned  not null default 0 comment '卖家ID',
    title        varchar(255)     not null default '' comment '商品标题',
    price        decimal(10, 2)   not null default 0.0 comment '交易金额',
    is_deleted   tinyint unsigned not null default 0 comment '是否删除, 0: 未删除, 1: 已删除',
    gmt_create   datetime         not null default current_timestamp comment '创建时间',
    gmt_modified datetime         not null default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id)
) comment '历史订单表' charset = utf8
                       engine = InnoDB;

insert into history_order (product_id, buyer_id, seller_id, title, price, is_deleted)
values (1, 1, 1, '商品1', 100.00, 0),
       (2, 1, 1, '商品2', 200.00, 0)
