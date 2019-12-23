# Chapter7: Concerrency

## 1 goroutine

- 메모리를 공유하는 별개의 흐름
- 그냥 go routine을 실행시키면 메인 흐름이 끝나기 전에 go routine이 완료되지 않을 수도 있다.

```go
go f(x, y, z)
```

1. paralleism, concerrency
    - paralleism: 물리적으로 별개의 업무를 수행
    - concerrency: 서로 의존관계 없이 논리적으로 병행해서 수행

1. go routine 기다리기
    - 싱크 라이브러리

    ```go
    // go routine이 몇 개 생길지 아는 경우
    var wg sync.WaitGroup // default 0인 counter
    wg.Add(len(urls)) // counter 증가
    for _, url := range urls {
        go func(url string) {
            defer wg.Done() // counter--, wg.Add(-1)과 같다
            if _, err := download(url); err != nil {
                log.Fatal(err)
            }
        }(url)
    }
    wg.Wait() // counter가 0이 될 때까지 기다린다.
    ```

    ```go
    // go routine이 몇 개 생길지 모르는 경우
    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1) // go routine이 시작되기 전에 counter++
        go func(url string) {
            defer wg.Done()
            if _, err := download(url); err != nil {
                log.Fatal(err)
            }
        }(url)
    }
    wg.Wait()
    ```

1. 공유 메모리와 병렬 최소값 찾기

    ```go
    func ParallelMin(a []int, n int) int {
        if len(a) < n {
            return Min(a)
        }
        mins := make([]int, n)
        size := (len(a) + n - 1) / n
        var wg sync.WaitGroup
        for i := 0; i < n; i++ {
            wg.Add(1)
            go func(i int) {
                defer wg.Done()
                begin, end := i*size, (i+1)*size
                if end > len(a) {
                    end = len(a)
                }
                mins[i] = Min(a[begin:end])
            }(i)
        }
        wg.Wait()
        return Min(mins)
    }
    ```

## 2 Chennel

- channel: 넣은 데이터를 뽑아낼 수 있는 파이프와 같은 형태의 자료구조
- channel은 포인터 같이 reference형 자료구조

```go
// map처럼 선언해야 한다.
// int형 channel 생성
c1 := make(chan int)
// 새로운 channel c2에 c1을 할당
// c2와 c1은 동일한 채널이 된다.
var chan int c2 = c1
var <-chan int c3 = c1 // c3는 receiver: 자료를 꺼내는 것만 가능
var chan<- int c4 = c1 // c4는 sender: 자료를 보내는 것만 가능

c1 <- 100 // channel c1에 100을 보낸다.
data := <-c1 // c1에서 자료를 받는다.
```

1. 일대일 단방향 채널 소통

    ```go
    func channel() {
        c := func() <-chan int {
            c := make(chan int)
            go func() {
                defer close(c)
                c <- 1
                c <- 2
                d <- 3
            }()
            return c
        }()
        for num := range c {
            fmt.Println(num)
        }
        // Output:
        // 1
        // 2
        // 3
    }
    ```

1. 생성기 패턴

    - 채널을 이용하는 방법의 장점
        - 생성하는 쪽에서 상태 저장 방법을 복잡하게 고민할 필요가 없다.
        - 받는 쪽에서 for의 range를 이용할 수 있다.
        - 채널 버퍼를 이용하면 멀티 코어를 활용하거나 입출력 성능상의 장점을 이용할 수 있다.

    ```go
    func BabyNames(first, second string) <-chan string {
        c:= make(chan string)
        go func() {
            defer close(c)
            for _, f := range first {
                for _, s := range second {
                    c <- string(f) + string(s)
                }
            }
        }()
        return c
    }

    func ExampleBabyNames() {
        for n := range BabyNames("성정명재경", "준호우훈진"){
            fmt.Print(n, ", ")
        }
    }
    ```

1. 버퍼 있는 채널

- 받는 쪽이 준비되지 않아도 보내는 쪽이 미리 버퍼에 보내 놓는 패턴
- 알려진 패턴을 따르는 것을 권장
- 버퍼 없는 채널로 동작하는 코드를 만들고,
- 필요에 따라 성능 향상을 위하여 버퍼 값을 조절해주자.

```go
// Working but bad example
c := make(chan int, 1)
c <- 5
fmt.Println(<-c)
```

1. 닫힌 채널

```go
val, ok := <-c // value, ok(채널이 열려 있는지 여부)
// 채널이 닫혀 있는 경우, val에는 default, ok에는 false
// 채널이 열려 있는데 받을 값이 없을 때: 받을 값이 생길 때까지 멈춰 있는다.
// 이미 닫혀 있는 채널을 또 닫으면 panic
```

## 3 동시성 패턴

1. 파이프라인 패턴

1. 채널 공유로 팬아웃하기

1. 팬인하기

1. 분산처리

1. select

    - 팬인하기

    - 채널을 기다리지 않고 받기

    - 시간 제한

1. 파이프라인 중단하기

1. 컨텍스트(context.Context) 활용하기

1. 용청과 응답 짝짓기

1. 동적으로 고루틴 이어붙이기

1. 주의점

## 4 경쟁 상태

1. 동시성 디버그

1. atomic과 sync.WaitGroup

1. sync.Once

1. Mutex와 RWMutex

1. 문맥 전환