package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	kv         map[string]string
	expire     map[string]bool
	expireDate map[string]time.Time
}

func NewCache() Cache {
	return Cache{map[string]string{}, map[string]bool{}, map[string]time.Time{}}
}

func (c *Cache) Get(key string) (string, bool) {

	x, err := c.kv[key]
	if err == false {
		return "", false
	}

	d, err := c.expireDate[x]
	if err == true {
		if !time.Now().Before(d) {
			return x, true
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

func main() {
	x := NewCache()
	x.PutTill("a", "value1", time.Now()) //
	x.Put("b", "value2")                 //
	x.Keys()                             // "b"

	x1, err := x.Get("a")
	fmt.Println(x1)  // ""
	fmt.Println(err) // false         pass

	x2, err := x.Get("b")
	fmt.Println(x2)  // "value2"
	fmt.Println(err) // true

	x3, err := x.Get("x")
	fmt.Println(x3)  // ""
	fmt.Println(err) // false

	x4 := x.Keys()
	fmt.Println(x4) // "b"

	fmt.Println(x)

}
