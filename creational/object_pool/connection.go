package object_pool

// Connection 具体连接池对象
type Connection struct {
	ID string
}

func (c *Connection) getID() string {
	return c.ID
}
