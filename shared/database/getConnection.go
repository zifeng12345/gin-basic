package database

func GetConnection() IConnection {
	return &conn
}

func StartTransaction() IConnection {
	connTrx.db = conn.db.Begin()
	return &connTrx
}
