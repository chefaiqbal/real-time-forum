CREATE TABLE IF NOT EXISTS ReactionPost (
	Value TEXT ,
	PostId INTEGER ,
	UserId INTEGER ,
	FOREIGN KEY (PostId) REFERENCES Post (Id)
)