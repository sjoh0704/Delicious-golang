# About Module

golang 1.11 이전 버전에서는 모듈이 없었고, GOPATH를 사용하였다. 
1.12 버전부터 모듈 기능을 제공하기 시작했으며, 1.16 버전까지 모듈 사용은 선택사항이었다. 이후 1.16 버전부터 모듈 사용이 default가 되었다.  

<br/>

## 모듈이란? 
패키지들을 모아 놓은 것


### 참고 
패키지: 코드를 묶는 단위로 모든 코드는 패키지로 묶여야 한다.

프로그램: 실행 파일로 즉 main 함수를 포함한 main 패키지 

<br/>

## 커스텀 모듈 사용하기 

다음과 같은 구조의 모듈을 구성한다. 
```
seung@seung-laptop:~/go/src/go-practice/testmodule$ tree .
.
├── custompkg
│   ├── custompkg2.go
│   └── custompkg.go
├── exinit
│   └── exinit.go
├── go.mod
├── go.sum
├── program
│   ├── main.go
│   └── program
└── README.md

3 directories, 8 files
```

<br/>


## 개요

1. testmodule 모듈 생성하기  
2. testmodule 내부에 커스텀 패키지 생성하기
3. testmodule 내부에 main 패키지를 실행할 program 생성하기
4. 추가적인 패키지를 다운 받기
5. 패키지 변수 초기화하기 

<br/>

## 실습
### 1. testmodule 디렉토리 생성하기 
testmodule 디렉토리가 루트 디렉토리가 된다. 
```
mkdir testmodule
```

testmodule을 모듈로 사용하기 위해서 다음 커맨드를 실행한다.

```
go mod init test
```

실행하면 go.mod라는 모듈이 생성되고, test라는 모듈을 사용할 수 있다.

<br/>


### 2. testmodule 내부에 커스텀 패키지 생성하기
```
cd testmodule
mkdir custompkg
```
custompkg.go을 생성한다. 

custompkg 디렉토리 내부에 생성되는 go파일들은 package custompkg라는 패키지명을 명시해야 한다. 

이후 main에서 사용할 함수를 하나 구현한다. 

코드는 testmodule/custompkg/custompkg.go를 참고한다. 

<br/>


### 3. testmodule 내부에 main 패키지를 실행할 program 생성하기
```
cd testmodule
mkdir program
```
main.go 파일 생성한 후, test/custompkg 라는 이름으로 패키지를 import한다. 

custompkg.원하는 함수명으로 함수를 호출한다. 

코드는 testmodule/program/main.go를 참고한다. 

<br/>

### 4. 추가적인 패키지를 다운 받기

이후 추가적인 패키지를 다운받기 위해 다음 커맨드를 실행한다. 

```
    go mod tidy 
```
<br/>

예를 들어, 위 커맨드를 통해 임포트된 github.com/tuckersGo/musthaveGo/ch16/expkg 패키지를 설치할 수 있다. 

<br/>

### 5. 패키지 변수 초기화하기 

exinit이라는 패키지를 생성하고 변수, init(), f(), PrintD를 선언한다. 

코드는 testmodule/exinit/exinit.go를 참고한다. 

실행 순서는 다음과 같다. 
1. 변수 초기화
2. 변수 초기화를 위해 f() 사용
3. init()
4. main에서 호출된 PrintD()

패키지를 중복해서 선언하는 경우, 처음 한번만 패키지 변수 초기화가 일어난다. 