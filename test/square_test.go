package main

import (
	"testing"
	"github.com/stretchr/testify/assert" // 테스트를 더 쉽게 만들어주는 패키지(go mod tidy)
)

func TestSquare1(t *testing.T){
	rst := square(9)
	if rst != 81{
		t.Errorf("square(9) should be 81 but returns %d", rst)
	}
}


func TestSquare2(t *testing.T){
	rst := square(3)
	if rst != 9{
		t.Errorf("square(3) should be 9 but returns %d", rst)
	}
}


func TestSquare3(t *testing.T){
	assert.Equal(t, 1, square(1), "square(1) should be 1")
	assert.Equal(t, 4, square(2), "square(2) should be 4")
	assert.Equal(t, 9, square(3), "square(3) should be 9")
	assert.Equal(t, 16, square(4), "square(4) should be 16")
	assert.Equal(t, 25, square(5), "square(5) should be 25")
	assert.Equal(t, 100, square(10), "square(10) should be 100") // 테스트를 한줄로 표현 가능하고 더 자세하게 나온다. 
	assert.Equal(t, 121, square(11), "square(11) should be 121")
}
