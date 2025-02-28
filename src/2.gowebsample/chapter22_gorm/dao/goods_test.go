package dao

import (
	"testing"
	"time"
)

func TestTranscation1(t *testing.T) {
	Transcation1()
}

func TestTranscation2(t *testing.T) {
	Transcation2()
}

func TestSave(t *testing.T) {
	goods := Good{
		Name:       "钩子函数",
		Price:      11111110,
		TypeId:     0,
		CreateTime: time.Now(),
	}
	Save(&goods)
}
