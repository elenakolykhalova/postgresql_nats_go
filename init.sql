CREATE TABLE IF NOT EXISTS Orders (
		order_uid	varchar(50) NOT NULL,
		data		text NOT NULL,
		PRIMARY KEY (order_uid)
);