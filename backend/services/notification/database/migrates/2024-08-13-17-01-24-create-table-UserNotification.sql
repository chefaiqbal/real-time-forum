CREATE TABLE IF NOT EXISTS UserNotification (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt TEXT DEFAULT CURRENT_TIMESTAMP,
	SenderId INTEGER ,
	ReceiverId INTEGER ,
	Read TEXT NOT NULL
)