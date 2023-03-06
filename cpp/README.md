### 1. example.go 작성(main 패키지만 가능)
### 2. 아카이브 파일 및 헤더 파일 생성
``` 
go build -buildmode=c-archive -o libexample.a example.go
```
### 3. cpp 코드 작성
### 4. cpp 빌드
```    
g++ main.cpp libexample.a -o cpptest -lpthread
```
### 5. 실행
```
./cpptest
```

