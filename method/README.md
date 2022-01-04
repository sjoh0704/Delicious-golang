# 메서드

메서드는 타입에 속한 함수 
메서드는 함수와 다르다!
리시버라는 것이 추가되어 있다. 

```
func (r Rabbit) info() int{
    return r.width * r.height
}

```

info는 Rabbit에 포함된다. 
리시버는 type으로 선언된 모든 지역 타입이 가능하다. 
int, string 등으로 리시버를 만들 수 없다. 
구조체나 별칭 타입 모두 가능하다. 

메서드는 왜 만들어졌을까?
=> 객체에게 기능을 제공함으로써, 객체 지향 프로그래밍을 가능하게 만든다. (결합도를 높인다)
객체는 데이터와 기능을 묶은 것

메서드는 객체와 객체의 관계를 만든다. 


### 언제 값 타입을 쓰고 언제 포인터 타입을 쓰는가?

time과 timer 패키지를 예로 들었을 때 
- time(시각 - 2022년 1월 10일 3시 10분)
    - 값 타입으로 사용
    - 객체가 의미가 없는 경우 

- timer(타이머 - 17분 30초 남음)
    - 포인터 타입으로 사용 
    - 필드 값이 바뀌더라도 객체가 바뀌면 안되는 경우 

하지만 정해진 것은 없다.
만약 값 타입으로 만들면 모두 값으로 만들고, 
포인터 타입으로 만들면 모두 포인터 타입으로 만드는게 좋다. 

go에서는 생성자와 소멸자가 없다. 
생성자가 없기 때문에 그냥 만들거나 초기화하는 함수를 만든다. 

ex) timer
NewTimer라는 함수를 만들어서 초기화하는 방식 
=> time.NewTimer(time.second)

## embedded field 메서드 