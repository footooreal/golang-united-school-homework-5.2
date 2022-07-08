package cache

import "time"

type Cache struct {
	kv         map[string]string
	expire     map[string]bool
	expireDate map[string]time.Time
}

func NewCache() Cache {
	return Cache{map[string]string{}, map[string]bool{}, map[string]time.Time{}}
}

func (c *Cache) Get(key string) (string, bool) {

	x, err1 := c.kv[key]
	if err1 == false {
		return "", false
	}

	d, err2 := c.expireDate[key]
	if err2 == true {
		if !time.Now().Before(d) {
			c.expire[key] = true
			return "", false
		}
	}

	return x, true
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.kv[key] = value
	c.expireDate[key] = deadline
}

func (c *Cache) Put(key, value string) {
	c.kv[key] = value
	c.expire[key] = false
}

func (c *Cache) Keys() []string {
	k := []string{}
	for i := range c.expire {
		if c.expire[i] == false {
			k = append(k, i)
		}
	}
	return k
}
