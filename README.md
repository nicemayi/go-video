# go-video

1. `go env`

2. `go build`

3. `env GOOS=linux GOARCH=amd64 go build`

4. `go get`

5. `go fmt`
6. `go test -v`

7. `fucn TestPrint(t *testing.T)`

8. `TestXXX`

9. `t.SkipNow()`

10. `t.Errorf("something something")`

11. `go test ./02`

12. `go test ./... -v`

13. `go test ./... -bench=.`

mysql -u root -h 157.230.169.141 -p
123456

401 通过了验证 403并没有

delete 204 no content

go build -o ../bin/govideo

如果是`go mod init demo`的形式，然后demo里两个包， app1, app2, 则应该引用的方式是`demo/app1`

go json tag CAN NOT HAVE SPACE CHAR

第三范式 3NF

````
create database videoserver;
use videoserver;

create table users (
	id int unsigned not null auto_increment,
	login_name varchar(64),
	pwd text not null,
	unique key (login_name),
	primary key (id)
);

create table video_info (
	id varchar(64) not null,
	author_id int(10),
	name text,
	display_ctime text,
	create_time datetime default current_timestamp,
	primary key (id)
);

create table comments (
	id varchar(64) not null,
	video_id varchar(64),
	author_id int(10),
	content text,
	time datetime default current_timestamp, primary key(id)
);

create table sessions (
	session_id VARCHAR(255) not null,
	TTL tinytext,
	login_name text,
    primary key (session_id(255))
);

create table video_del_rec (
	video_id varchar(64) not null,
	primary key (video_id)
);
