# About Module

golang 1.11 이전 버전에서는 모듈이 없었고, GOPATH를 사용하였다. 
1.12 버전부터 모듈이 만들어졌고, 1.16 버전까지 모듈 사용은 선택사항이었다. 이후 1.16 버전부터 모듈 사용이 default가 되었다.  

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
seung@seung-laptop:~/go/src/go-practice/usepkg$ tree .
.
├── custompkg
│   ├── custompkg2.go
│   └── custompkg.go
├── go.mod
├── go.sum
├── program
│   ├── program
│   └── usepkg.go
└── README.md

2 directories, 7 files
```

1. testmodule 디렉토리 생성 => 루트 디렉토리가 된다. 
2. testmodule 내부에 custompkg 생성하기 => 커스텀 패키지를 만들 것이다. 
3. testmodule 내부에 program 생성하기 => main 패키지를 사용할 것이다. 
4. 루트 디렉토리(testmodule)에서 go mod init test => test라는 모듈명을 사용하겠다. 
5. 이후 go mod tidy를 실행 => 추가적인 패키지를 다운 받는다.(ex. github.com/~ )

