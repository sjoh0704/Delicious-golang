# 함수 고급

### 가변인자 
...를 사용하면 가변인자로 받을 수 있다. 
 
```
func Print(args ...interface{}){
    for _, arg := range args{
        switch t:= arg.(type){
            case bool: ~
            case float64: ~
            default: ~ 
        }
    }
}
```
빈 인터페이스가 나오는 경우 타입 변환이 따라온다. 

### defer
함수 종료전에 실행을 보장한다
주로 OS 자원을 반납할 때 사용한다.
마지막에 잊지 않고 자원을 반납할 수 있다.


### 함수 타입 변수
함수를 값으로 갖는 변수 

함수 타입은 함수 시그니처로 표현
```
func (int, int) int 
```