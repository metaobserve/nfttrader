create table wallet
(
    id int auto_increment primary key,
    # user's wallet
    address nvarchar(200),
    # access token, which is created by login.
    token nvarchar(200),
    # login Time
    loginTime datetime,
    # logout Time
    logoutTime datetime,
    createTime datetime,
    updateTime datetime
);
# user list records who join the airdrop
create table userList
(
    id int auto_increment primary key ,
    address nvarchar(200),
    airdropId int,
    # may be there is a white list to protect users who made an appointment before.
    joinedWhiteList bool,
    createTime datetime,
    updateTime datetime
);

# airdrop has three stages.
create table airdrop
(
    id int auto_increment primary key ,
    name nvarchar(200),
    description nvarchar(400),
    startTime datetime,
    endTime datetime,
    # airDrop stage: 0 ready; 1 action; 2 cut
    stage integer,
    createTime datetime,
    updateTime datetime
);
# ntf
create table nft
(
    id int auto_increment primary key ,
    name nvarchar(200),
    description nvarchar(400),
    # nft token
    token varchar(200),
    airdropId int,
    # nft status : 0 invalid 1 valid
    status int,
    createTime datetime,
    updateTime datetime
);