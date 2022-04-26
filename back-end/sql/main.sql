CREATE database Todo DEFAULT CHARACTER SET utf8;
use Todo;
create table User (
    id char(10) UNIQUE,
    name varchar(30) not null,
    email varchar(255) not null UNIQUE,
    password varchar(255) not null
);

INSERT into User(id,name,email,password) values(?,?,?,?)
UPDATE User set email = 'daniel081009a@gmail.com' where id = 'mscUdwjUnQ'

create table block (
    user_id char(10) NOT NULL,
    id char(10) UNIQUE,
    type int, 
    data text,
    top_id char(10),
    create_date timestamp default now(),
    final_date timestamp, 
    CONSTRAINT fk_block_top_id 
    FOREIGN KEY (top_id) 
    REFERENCES block (id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE,
    CONSTRAINT fk_block_user_id 
    FOREIGN KEY (user_id) 
    REFERENCES User(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE
);
INSERT into block(user_id,id,type,data,top_id) values(?,?,?,?,?);
INSERT into block(user_id,id,type,data,top_id,final_date) values("PfNuazVaAV","NCSTaGiqGe",1,"superDan","BWstccdAct","2022-10-09");