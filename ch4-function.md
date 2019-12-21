# Chapter4: Function

## 1 값 넘겨주고 넘겨받기

1. 값 넘겨주기

    - &: 주소값 반환
    - *: 참조값 반환

1. 둘 이상의 반환값

    ```go
    func Foo() error{}
    func Foo2() (int, error) {} // 둘 이상의 return
    integer, err := Foo2()
    _, err := Foo2() // return 생략
    // (convention) error는 가장 마지막에 반환
    ```

1. 에러값 주고 받기

    ```go
    // error handling
    func Foo3() error {
        if err := Foo(); err != nil { // err인지 확인하고 if구문 내에서 소멸
            return err // Foo의 error를 받아서 그대로 반환
            return errors.New("Foo: sth bla bla") // 새 에러 반환
            return fmt.Errorf("Foo has problem: line at %d", count) // 추가 정보 구성 가능
        }
    }
    ```

1. 명명된 결과 인자

    ```go
    func Foo() err error { // return 값에 이름을 붙여서 사용
        return
    }
    ```

1. 가변인자

    ```go
    func Foo(nums []int){} // slice parameter
    func Foo2(nums...) {} // another method
    
    integers := []int{1,2,3}
    Foo2(integers)
    Foo2(1,2,3)
    Foo(integers)
    Foo(1,2,3) // Imposible
    ```

## 2 값으로 취급되는 함수

1. function literal

    - function is first-class object(citizen)

    ```go
    func (a, b int) int {}
    ```

1. higher-order function

    ```go
    // ReadForm function with parameter reader stream and function do scan reader and do function
    func ReadForm(r io.Reader, f func(line string)) error {
        scanner := bufio.Newscanner(r)
        for scanner.Scan() {
            f(scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            return err
        }
        return nil
    }

    // call ReadForm with literal function
    r := strings.NewReader("leo\nleohn\nlee")
    err := ReadForm(r, func(line string) {
        fmt.Println(line)
    })
    ```

1. Closure

    - Closure: 외부에서 선언된 변수를 literal function 내에서 접근할 수 있는 코드

    ```go
    // continue the higher-order function
    var lines = []string
    err := ReadForm(r, func(line string){
        lines = append(lines, line) // global에서 선언된 lines var를 literal 안에서 접근한다.
    })
    ```

1. Generator

    ```go
    func Foo() func() int {
        var next int // define int -> default 0
        return func() int {
            next++
            return next
        }
    }
    generator := Foo()
    fmt.Println(generator())

    generator2 := Foo() // var next is different btw generator and generator2

    // Output: 1
    ```

1. Named Type

    - type inspection
    - 명명되지 않은 자료형(원형)이라고 해도 타입 검사 불통

    ```go
    type NewInt int
    newInteger := 1
    integer := 1

    func Foo(num NewInt) {}
    Foo(integer) // Compile Error
    Foo(newInteger) // Work well
    ```

1. Named Function

    - 원형(명명되지 않은 함수형)을 parameter로 사용 시, 타입 검사 통과

    ```go
    type FooType1 func(int, int) int
    type FooType2 func(int, int) int
    
    func Foo(f Footype1) {}

    Foo(FooType1(a, b)) // work
    Foo(func(a,b) int{
        return a + b
    }) // work
    Foo(FooType2(a, b)) // Not Work
    ```

1. 인자 고정

    - 뭔가 어렵다...??: 나중에 다시 볼 것..

1. 패턴의 추상화

    - higher order function을 이용해 추상화...

1. 자료구조에 담은 함수

    - ..!??

## 3 Method

- method: function + receiver
- 모든 Named Type에 대해 Method 선언 가능
- public: method 이름이 대문자로 시작
- private: method 이름이 소문자로 시작

    ```go
    // method MethodName
    // data type T
    // receiver name recv
    // when call MethodName as type T
    func (recv T) MethodName(p1 T1) R1 {}
    ```

1. 단순 자료형 메서드
1. 문자열 다중 집합

    ```go
    type MultiSet map[string]int
    
    func (m MultiSet) Insert(val string) {
        m[val]++
    }

    func (m MultiSet) Erase(val string) {
        if m[val] <= 1 {
            delete(m, val)
        } else {
            m[val]--
        }
    }

    func (m MultiSet) Count(val string) int {
        return m[val]
    }

    func (m MultiSet) String() string {
        s := "{ "
        for val, count := range m {
            s += strings.Repeat(val+" ", count)
        }
        return s + "}"
    }
    ```

1. 포인터 리시버

    ```go
    func (adjList *Graph) ReadFrom(r io.Reader) error {}
    ```

1. 공개 및 비공개

    - 모든 객체명에 동일하게 적용.
    - private은 해당 모듈에서만 접근 가능
    - 여러 파일에 나눠져 있더라도 같은 모듈이면 private도 접근 가능
    - ***public은 주석을 꼭 달자!! godoc이 문서를 만들어준다***
    - public: method 이름이 대문자로 시작
    - private: method 이름이 소문자로 시작

## 4 활용

1. blocking(synchronous), non-blocking(asynchronous): 타이머 활용하기

    - blocking or synchronous

    ```go
    // 1초에 한 번씩 몇 초 남았는지 프린트
    func countDown(seconds int) {
        for seconds > 0 {
            fmt.Println(seconds)
            time.Sleep(time.Second)
            seconds--
        }
    }
    ```

    - non-blocking or asynchronous

    ```go
    // 이후 코드들이 실행되다가 비동기적으로 5초 뒤에 literal func 수행
    time.Afterfunc(5*time.Second, func(){
        // sth do
    })
    ```

1. path/filepath 패키지

    - 파일 이름 경로 관리 패키지
    - https://golang.org/pkg/path/filepath/#Walk
