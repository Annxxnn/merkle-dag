package merkledag

// Has(key []byte) (bool, error)：检查存储中是否存在给定的键。如果存在，返回 true，否则返回 false。如果在检查过程中发生错误，返回错误信息。
// Put(key, value []byte) error：将键值对存储到存储系统中。使用给定的键和值进行存储。如果存储过程中发生错误，返回错误信息。
// Get(key []byte) ([]byte, error)：根据给定的键从存储中检索对应的值。如果找到对应的值，则返回该值的字节数组。如果在检索过程中发生错误或找不到对应的值，返回错误信息。
// Delete(key []byte) error：根据给定的键从存储中删除对应的键值对。如果删除成功，返回 nil。如果在删除过程中发生错误，返回错误信息。
type KVStore interface {
	Has(key []byte) (bool, error)
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
}
