create table account
(
    id int null
);

create table airdrop
(
    id          int auto_increment
        primary key,
    name        varchar(200) charset utf8mb3 null,
    description varchar(400) charset utf8mb3 null,
    startTime   datetime                     null,
    endTime     datetime                     null,
    stage       int                          null,
    createTime  datetime                     null,
    updateTime  datetime                     null,
    tokenId     varchar(50)                  null
);

create table mint
(
    id         int auto_increment
        primary key,
    address    varchar(200) charset utf8mb3 null,
    nftId      int                          null,
    airdropId  int                          null,
    createTime datetime                     null,
    updateTime datetime                     null
);

create table nft
(
    id             int auto_increment
        primary key,
    name           varchar(200) charset utf8mb3  null,
    description    varchar(400) charset utf8mb3  null,
    token          varchar(200)                  null,
    airdropTokenId int                           null,
    status         int                           null,
    metadata       varchar(2000)                 null,
    theme          varchar(200)                  null,
    category       varchar(40)                   null,
    price          varchar(50)                   null,
    priceUnit      varchar(50)                   null,
    nftAddress     varchar(2000) charset utf8mb3 null,
    isBanner       tinyint(1)                    null,
    author         varchar(100)                  null,
    authorAddress  varchar(2000)                 null,
    createTime     datetime                      null,
    updateTime     datetime                      null
);

create table userList
(
    id              int auto_increment
        primary key,
    walletAddress   varchar(200) charset utf8mb3 null,
    airdropTokenId  varchar(50)                  null,
    joinedWhiteList tinyint(1)                   null,
    createTime      datetime                     null,
    updateTime      datetime                     null
);

create table wallet
(
    id         int auto_increment
        primary key,
    address    varchar(200) charset utf8mb3 null,
    token      varchar(200) charset utf8mb3 null,
    loginTime  datetime                     null,
    logoutTime datetime                     null,
    createTime datetime                     null,
    updateTime datetime                     null
);

create table whiteList
(
    id             int auto_increment
        primary key,
    name           varchar(50) charset utf8mb3  null,
    description    varchar(400) charset utf8mb3 null,
    startTime      datetime                     null,
    endTime        datetime                     null,
    airdropTokenId varchar(50)                  null,
    status         int                          null,
    walletAddress  varchar(50)                  null,
    createTime     datetime                     null,
    updateTime     datetime                     null
);

