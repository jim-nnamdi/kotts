CREATE TABLE users (id int not null primary key AUTO_INCREMENT, username VARCHAR 191, password text, email text, country text, active text)

CREATE TABLE transactions(id int not null, from_user int references users(id), to_user int references users(id), amount bigint, tx_fee bigint, tx_status text, created_at datetime, updated_at datetime);

CREATE TABLE docs(id int not null primary key AUTO_INCREMENT, author int references users(id), content text, created_at datetime, updated_at datetime);

CREATE TABLE courses(id int not null primary key AUTO_INCREMENT, author int references users(id), module1 text, module2 text, module3 text, module4 text, module5 text, created_at datetime, updated_at datetime, total_views bigint);

CREATE TABLE chats(id int not null primary key AUTO_INCREMENT, from_user int references users(id), to_user int references users(id), connection_id int references connections(id), created_at datetime, updated_at datetime);

CREATE TABLE connections (id int not null primary key AUTO_INCREMENT, chat_id int references chats(id), created_at datetime, updated_at datetime);