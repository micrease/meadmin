package tests

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func v2() {
	u1 := uuid.NewV2(uuid.DomainGroup).String()
	fmt.Println("go", u1)
}

//都是一样的。。
func TestUuidV2(t *testing.T) {

	go v2()

	u1 := uuid.NewV2(uuid.DomainGroup).String()
	fmt.Println("testv2", u1)

	time.Sleep(time.Second * 3)
	go v2()
	u2 := uuid.NewV2(uuid.DomainGroup).String()
	fmt.Println("testv2", u2)
}

func TestUuidV5(t *testing.T) {

	u1 := uuid.NewV4().String()
	fmt.Println(u1)

	u2 := uuid.NewV4().String()
	fmt.Println(u2)
}
