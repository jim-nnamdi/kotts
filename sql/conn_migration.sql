CREATE TABLE connections (id int not null primary key AUTO_INCREMENT, chat_id int references chats(id), created_at datetime, updated_at datetime);