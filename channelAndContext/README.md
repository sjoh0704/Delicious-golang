# 채널과 컨텍스트
채널, 컨텍스트, 고루틴 => 동시성 프로그래밍

### 채널
채널은 고루틴끼리 메시지를 전달할 수 있는 메시지 큐 

채널을 사용하면 locking을 안해도 된다. 

chan은 채널 키워드, string은 메시지의 타입을 의미한다. 
```
var messages chan string = make(chan string)
```

채널에 데이터 넣기 
```
messages <- "this is a data"
```

채널에서 데이터 빼기 
```
var msg string = <- messages
```

### 채널 크기
채널의 기본 크기는 0이다. 
크기 = 보관함
보관함이 없는 경우 수신자가 올때까지 기다린다.
보관함이 있을 때 보관함에 보관하고 다음일을 한다. 
