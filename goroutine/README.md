# 고루틴

### 스레드
스레드는 프로그램 내의 실행 흐름을 말하며 초기에는 하나의 실행 흐름만 존재한다. 

### 멀티 스레드 
코어가 여러 스레드를 빠르게 번갈아가면서 실행시킨다. 
OS는 이러한 스레드를 스케줄링을 통해서 실행시킨다. 

### 컨텍스트 스위칭
스레드 전환을 컨텍스트 스위칭이라 하며 스레드 전환에는 비용이 발생한다.
스택은 스레드마다 따로 있으며, 힙은 모든 스레드가 공유한다.

- 코어가 하나인데 스레드가 3개라면, 컨텍스트 스위칭이 많이 발생
- 코어가 3개인데 스레드가 3개라면, 각자 하나씩 사용하므로 컨텍스트 스위칭이 발생하지 않는다. 

### 멀티 스레드와 멀티 프로세스
- 멀티프로세스: 계산기라는 프로그램이 있고, 이 프로그램을 실행시키면 메모리 상에 계산기 프로세스가 동작한다. 이 때 계산 프로그램을 하나 더 실행시키면 여러개의 계산기 프로세스가 동작하게 되고, 이를 멀티 프로세스라고 한다.
- 멀티 스레드: 하나의 프로그램안에서 여러개의 스레드가 있는 것을 멀티 스레드라고 한다. 

### Go Routine
고에서 만든 경량 스레드 
메인함수도 고루틴이며 이를 메인 고루틴이라고 한다.
그래서 go 프로그램은 최소 하나의 고루틴을 갖는다. 

### 고루틴 동작원리
다른 언어들은 스레드를 지원하지만 고루틴의 경우에는 OS 스레드를 이용하는 경량 스레드이다. 

고루틴 != 스레드 

고루틴은 스레드를 이용하는 것이다. 

< 코어 2개, 고루틴 3개의 경우 >
코어 => OS 스레드1 => 고루틴1 
코어 => OS 스레드2 => 고루틴2
코루틴3는 대기한다. 

실행 중인 고루틴이 끝나면 고루틴3이 놀고 있는 스레드를 사용하기 위해 빈 자리에 들어간다.
상황에 따라 교체가 되기도 한다. 

### 시스템 콜 호출시 
만약 고루틴1이 시스템 콜을 호출하면 고루틴1은 응답이 올 때까지 대기하게 되고, 
해당 자리를 고루틴2가 들어가게 된다.
즉, 고루틴 수준에서의 컨텍스트 스위칭이 발생한다. 


### 결론
- 스레드의 갯수는 코어의 갯수만큼 만들어진다. 
- 아무리 많은 고루틴을 생성해도 스레드 단위의 컨텍스트 스위칭 비용은 발생하지 않는다. 
- OS단에서 스레드의 컨텍스트 스위칭이 발생하지 않는다. 
- 대신 고루틴을 교체하는 스위칭 비용이 들지만 이는 아주 작기 때문에 빠르다!
- 고루틴을 사용하지 않는 일반적인 경우, 커널 스레드의 갯수보다 프로세스 내 스레드의 갯수가 늘어나면 컨텍스트 스위칭 비용이 더 늘어나 성능이 떨어진다. 


### 동시성 프로그래밍의 주의점
- 동일한 자원을 여러 고루틴에서 접근할 때 동시성 문제가 발생한다. 
- 여러개의 스레드가 있을 때 `heap 메모리`를 공유하게 되고,접근시 문제가 될 수 있다. 
- 이러한 문제를 해결하기 위해서는 메모리 자원을 하나의 고루틴에서만 접근하게 해야 한다.
- 간단한 방법으로 `Lock을 거는 방법`이 있는데 이를 `뮤텍스`라 한다. 

### 뮤텍스의 문제점
- 동시성 프로그램으로 인한 성능 향상을 얻을 수 없다.
- 심지어 과도한 locking으로 성능이 하락되기도 한다.(lock 자체의 성능 문제) 
- 동시성 프로그램을 사용하는 의미가 없어진다. 
- 더 큰 문제는 고루틴을 완전히 멈추게 만드는 데드락 문제가 발생할 수 있다.
- 데드락은 필요한 자원들이 모두 락이 걸려 있어서 이러지도 저리지도 못하다가 프로그램이 죽는 것이다. 

### 뮤텍스 결론
- 매우 조심히 사용해야 한다.
- 사용하지 말아야 하는 것은 아니다. 동시성 문제에 대한 simple한 해결책이 될 수 있기 때문

### 자원 관리 기법
1. `영역`을 나누는 방법: 영역을 나누면 서로 부딪힐 일이 없다.
2. `역할`을 나누는 방법: 각기 다른 역할을 함으로써 충돌을 피해간다. 