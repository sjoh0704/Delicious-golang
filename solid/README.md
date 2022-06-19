# SOLID
OOP를 위한 5가지 설계 원칙 
1. 단일 책임 원칙(single responsibility)
2. 개방 폐쇄 원칙(open-closed)
3. 리스코프 원칙(liskov substitution)
4. 인터페이스 분리 원칙(interface segregation)
5. 의존 관계 역전 원칙(dependency inversion)

solid 원칙은 세계평화와 같은 목표이다.
이룰 수 없지만, 지향할 수는 있는 것이다. 

## 나쁜 설계
1. 경직성: 모듈간의 결합도가 너무 높아서 코드를 변경하기 어려운 구조 
2. 부서지기 쉬움: 한 부분을 건드렸더니 다른 부분까지 망가지는 경우
3. 부동성: 모듈간 결합도가 너무 높아서 코드 일부분을 분리해서 사용할 수 없는 경우. 코드 재사용률이 감소한다. 


## SOLID
1. 단일 책임 원칙(제일 중요)
모든 객체는 책임을 하나만 져야 한다. 
코드 재사용성을 높여 준다. 

다음은 단일 책임에 원칙에 위배된다. 
```
type FinanceReport struct{ // 회계 보고서 
    report string
}
func (r *FinanceReport) SendReport(email string){ // 보고서 전송 
    ... 
}
```
만약 회계보고서가 아닌 마케팅 보고서가 나오면 SendReport를 한번 더 만들어줘야하는 문제가 생긴다. 

즉 전송과 보고서를 따로 만들어야 한다. => 중간에 인터페이스를 구현한다.  
FinanceReport => Report(interface) => ReportSender
구체화된 객체간의 커플링(의존도)이 낮아진다. 
커플링이 없을 수는 없다. 


2. 개방 폐쇄 원칙
확장에는 열려 있고, 기존 코드 변경에는 닫혀 있다.
상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 된다. 

다음은 잘못된 경우
```
func SendReport(r *Report, method SendType, receiver string){
    switch method{
        case Email: ...
        case Fax: ...
        case PDF: ...
        case Printer: ...
    }
}

```
위 코드는 method의 종류가 추가될 때마다 SendReport 함수는 계속 변경되어야 한다. 
기존 구현이 변경되므로 이는 좋은 코드가 아니다.
이를 기능별로 구별해야  한다.

```
type EmailSender struct{
}

func (e *EmailSender) Send(r * Report){
}

type FaxSender struct{
}

func (f *FaxSender) Send(r *Report){
}
```
그리고 이를 switch case로 묶는다. 

3. 리스코프 치환 원칙
상위 타입에서 동작한다면 하위 타입에서도 동작해야 한다.
```
func q(c T){

}
// 위 함수가 동작한다면 T의 하위 타입 S(T를 상속한)에 대해서도 다음이 동작해야 한다. 
var y S = ...
q(y) 
```
`일반적으로 Go에서는 상속을 지원하지 않기 때문에 리스코프 치환 원칙을 위반하지 않는다.` 

하지만 인터페이스 타입 변환시 리스코프 치환 법칙을 위배할 수 있다.
내부에서 interface를 concrete object로 바꾸는 경우 동작하지 않는 경우가 발생할 수 있다.

ex) Report 인터페이스(상위 객체), Report 인터페이스를 이용한 FinanceReport struct(하위 객체)가 있을 때,  
Report 인터페이스를 이용한 함수 func SendReport(r Report)는 내부에서 FinanceReport로의 타입 변환(concrete object로의) 구문만 있다면, FinanceReport의 경우 동작하지만 다른 객체인 MarketingReport에서는 동작하지 않을 수 있다.
즉, 상위 객체에 대해서는 동작하지만 하위 객체에서는 동작하지 않는 상황이 생기고, 이는 LSP를 위반한다. 

`웬만하면 인터페이스 타입 변환 같은 다이나믹 캐스팅은 하지 말자. `

4. 인터페이스 분리 원칙
`클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.`

다음 경우 인터페이스 분리 원칙을 위반한다. 

```
type Report interface{
    Report() string
    Pages() int
    Author() string
    WrittenDate time.Time
}
// 위와 같이 Report 인터페이스는 굉장히 많은 메서드를 가지고 있음...

func SendReport(r Report){
    send(r.Report())
}
```
Report 인터페이스를 사용하기 위해서는 반드시 4개의 메서드를 가져야 하며, 
SendReport를 사용하기 위해서는 Report()라는 메서드 하나만 가지면 된다.
즉 필요없는 메서드가 3개나 있는 것.
이는 인터페이스 분리 법칙을 위반한 것이다. 


5. 의존 관계 역전 원칙
상위 계층이 하위 계층에 의존하는 전통적인 `의존 관계를 역전시킴`으로써 `상위 계층이 하위 계층의 구현으로부터 독립`되게 할 수 있다.
원칙 1. 상위 모듈은 하위 모듈에 의존해서는 안된다. 둘 다 추상 모듈에 의존해야 한다.
원칙 2. 추상 모듈은 구체화된 모듈에 의존해서는 안된다. 구체화된 모듈은 추상 모듈에 의해 의존해야 한다. 

### 원칙1
전송(상위 인터페이스), 키보드(전송의 하위 객체), 네트워크(전송의 하위 객체) 라고 했을 때
[키보드로부터 입력 => 전송 => 네트워크로 전달]는 강하게 커플링되어 있다.
이는 `상위 객체가 하위 객체에게 의존성`을 갖게 되고, 이는 `의존 관계가 역전`된 것이다.   
[키보드 <= 입력 인터페이스 <= 전송 => 출력 인터페이스 => 네트워크] 
사실상 개방 폐쇄 원칙의 예와 유사하다. 

### 원칙2
다음은 메일이 알람을 소유하고 있다. 
이 경우는 `구체화된 모듈이 다른 구체화된 모듈에 의존`하고 있으므로 의존 관계 역전 원칙을 위배한다. 
```
type Mail struct{
    alarm Alarm
}

type Alarm struct{
}

func (m *Mail) onRecv(){
    m.alarm.Alarm()
}
```

다음과 같은 관계로 설계한다.
[메일 => 이벤트 인터페이스 <=> 이벤트 리스너 인터페이스 => 알람] 
구체화된 모듈이 추상화된 모듈과 추상화된 모듈과 추상화된 모듈이 연결 

이러한 원칙을 지키는 것이 좋지만 항상 모든 원칙을 지킬 수는 없다.
더 복잡해지거나 비효율적이 될 수도 있다. 

