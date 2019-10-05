package dataStruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	a := assert.New(t)
	list := NewLinkedList()
	list.Append("1")
	list.Append("2")
	a.Equal(list.lens,2)
	a.Equal(list.Search("1"),true)
	a.Equal(list.Search("2"),true)
	a.Equal(list.Search("3"),false)
	a.Equal(list.Remove("2"),true)
	a.Equal(list.Search("1"),true)
	a.Equal(list.Lens(),1)
}
