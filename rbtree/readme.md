# Red-Black Tree
Yet another red-black trees implementation written in Go, strong testing covered.


## Example

```go
import "github.com/chasecs/tree/rbtree"

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func main() {

    tree := rbtree.New()
    tree.Insert(key(1), "val_1")
    tree.Insert(key(3), "val_3")

    n := tree.Find(key(3))
    d := tree.Delete(key(1))

    fmt.Println(tree.IsBalance())
}
```

## Similar Projects
- [sakeven/RbTree](https://github.com/sakeven/RbTree)
- [petar/GoLLRB](https://github.com/petar/GoLLRB)
