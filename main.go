package main

import "fmt"

type MyErr string

func (m MyErr) Error() string {
	return "big fail"
}

func doSomething(i int) error {
	switch i {
	default:
		return nil // == nil
	case 1:
		var p *MyErr
		return p // != nil
	case 2:
		return (*MyErr)(nil) // != nil
	case 3:
		var p *MyErr
		return error(p) // != nil because the interface points to a
		// nil item but is not nil itself.
	case 4:
		var err error // == nil: zero value is nil for the interface
		return err    // This will be true because err is already interface type
	}
}

func main() {
	for i := 0; i <= 4; i++ {
		err := doSomething(i)
		fmt.Println(i, err, err == nil)
	}
}
