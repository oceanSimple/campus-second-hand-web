use website;

drop table if exists product;

create table product
(
    id           bigint unsigned auto_increment comment '商品id',
    title        varchar(255)     not null default '' comment '商品标题',
    description  varchar(255)     not null default '' comment '商品描述',
    price        decimal(10, 2)   not null default 0.00 comment '商品价格',
    stock        int unsigned     not null default 0 comment '商品库存',
    campus       varchar(255)     not null default '' comment '校区',
    address      varchar(255)     not null default '' comment '地址',
    seller_id    bigint unsigned  not null default 0 comment '卖家id',
    buyer_id     bigint unsigned  not null default 0 comment '买家id',
    picture1     varchar(255)     not null default '' comment '商品图片1',
    picture2     varchar(255)     not null default '' comment '商品图片2',
    picture3     varchar(255)     not null default '' comment '商品图片3',
    type1        int unsigned     not null default 0 comment '商品类型1',
    type2        int unsigned     not null default 0 comment '商品类型2',
    is_deleted   tinyint unsigned not null default 0 comment '是否删除, 0: 未删除, 1: 已删除',
    gmt_create   datetime         not null default current_timestamp comment '创建时间',
    gmt_modified datetime         not null default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id),
    index title (title)
) comment '商品表' charset = utf8
                   engine = InnoDB;

insert into product (title, description, price, stock, campus, address, seller_id, buyer_id, picture1, picture2,
                     picture3, type1, type2)
values ('iphone 12', 'iphone 12 128G', 6999.00, 100, '南湖校区', '二舍', 1, 1,
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg', 1, 1),
       ('iphone 11', 'iphone 11 128G', 5999.00, 100, '南湖校区', '二舍', 1, 1,
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg', 1, 1),
       ('iphone 10', 'iphone 10 128G', 4999.00, 100, '南湖校区', '二舍', 1, 1,
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg',
        'https://img.alicdn.com/imgextra/i1/2200725021/O1CN01Zz1Q2o1Q5Q6Q6Q5Q6_!!2200725021.jpg', 1, 1);