# Chapter 5: Structures and Interfaces

- structure: set of fields
- interface: set of method

## 1 Structure

1. 구조체 사용법

    ```go
    type Task struct {
        title string
        done bool
        due *time.Time // 기한은 없을 수도 있으니 포인터 사용
    }

    task1 := Task{"laundry", false, nil}
    task2 := Task{
        title: "laundry", 
        done: true,
    }
    ```

1. const, iota

    - bool 대신 enum: 확장성이 좋다.
    - iota: auto increament

    ```go
    // Task struct
        // done bool
        status status

    type status int

    const (
        UNKNOWN status = iota
        TODO
        DONE
    )
    ```
    
    ```go
    // iota 응용
    // golang.org/doc/effective_go.html#constants
    type ByteSize float64
    const (
        _ = iota // ignore first value
        KB ByteSize = 1 << (10*iota)
        MB
        GB
        TB
        ...
    )

    // another method
        KB ByteSize = 1 << (10*(1+iota))
        MB
        ...

1. 테이블 기반 테스트

    - go는 generic 지원이 안되서 assertion을 이용한 unit test가 안된다.
    - if문을 이용하거나, 테이블 기반 테스트를 권장

    ```go
    // fibonaci test
    testCases := []struct {
        in int
        out int
    }{
        {0, 0},
        (5, 5),
        {6, 8},
    }

    func TestFib(t *testing.T) {
        for i, c := range cases {
            dest := seq.fib(c.in)
            if dest != c.out {
                t.Errorf("Case %d: Fib(%d) == %d, want %d", i, c.in, dest, c.out)
            }
        }
    }
    ```

1. 구조체 내장

    - go에서 structure는 여러 자료형의 필드를 가질 수 있다는 점이 가장 중요.

    ```go
    type Deadline time.Time // define deadline type

    // OverDue returns true if the deadline is before the current time
    func (d *Deadline) OverDue() bool {
        return d != nil && time.Time(*d).Before(time.Now())
    }
    
    // Task structure에 활용
    type Task struct {
        title string
        status status
        Deadline *Deadline
        // *Deadline // Task에 대하여 OverDue method를 또 작성할 필요 없다.
    }
    ```

    - 필드가 내장되어 있으면 내장된 필드가 구조체 전체의 직렬화 결과를 바꿔버리는 문제가 생길 수 있다.

    ```go
    // 필드 내장 문제를 해결하기 위해 Deadline을 structure로
    type Deadline struct {
        time.Time
    }

    func NewDeadline(t time.time) *Deadline {
        return &Deadline{t}
    }

    type Task struct {
        title string
        status status
        Deadline *Deadline
    }
    ```

    - 구조체를 내장하면 내장된 구조체에 들어 있는 필드도 바로 접근 가능

## 2 (De)Serialize

- serialization: 객체의 상태를 보관이나 전송이 가능한 상태로 변환하는 것.
- de~: 다시 객체로 복원하는 것

1. JSON

    - serialization
        - json package: 대문자로 시작하는 필드들만 json으로 직렬화

    ```go
    byteSlice, err := json.Marshal(data)
    ```

    - deserialization
    
    ```go
    err := json.Unmarshal(byteSlice, &data)
    ```

    - tag
        - js는 8byte 실수형이라서 int64를 그대로 넘겨주면 정확도 문제가 생긴다. 따라서 string으로 넘겨주는 습관을 들이자.

    ```go
    type Struct struct {
        Title string `json:"title"` // "title"로 표현
        Internal string `json:"-"` // json에서 무시
        Value int64 `json:",omitempty"` // 0인 경우에 json에 결과 나타나지 않음
        ID int64 `json:",string"` // int지만 json에서는 string으로 표현
    }
    ```

    - json serialization user definition
        - status가 int형이지만, string으로 사용하고 싶을 때.
    
    ```go
    func (s status) MarshalJSON() ([]byte, error) {
        switch s {
        case UNKNOWN:
            return []byte(`"UNKNOWN"`), nil
        case TODO:
            return []byte(`"TODO"`), nil
        case DONE:
            return []byte(`"DONE"`), nil
        default:
            return nil, errors.New("status.MarchalJSON; unkown value")
        }
    }

    func (s *status) UnmarchalJSON(data []byte) error {
        switch string(data){
        case `"UNKNOWN"`:
            *s = UNKNOWN
        case `"TODO"`:
            *s = TODO
        case `"DONE"`:
            *s = DONE
        default:
            return errors.New("status.UnmarchalJSON; unkown value")
        }
        return nil
    }
    ```

    - 구조체가 아닌 자료형 처리
        - 배열도 가능
        - 여러 필드가 있을 때는 구조체 또는 맵을 활용하자
        - 맵을 활용하는 경우, json으로 바뀌면 키가 정렬되서 출력된다.

    - json 필드 조작
        - 구조체에서 특정 필드를 빼고 직렬화하고 싶은 경우

        ```go
        // 구조체 내장을 활용
        type Fields struct {
            VisibleField string
            InvisibleField string
        }

        func OmitFields() {
            f := &Fields{"a","b"}
            b, _ := json.Marshal(struct{
                *Fields
                InvisibleField string `json:",omitempty"`
                Additional string
            }{
                Fields: f, Additional: "c"
            })
            
        }
        ```

        - json 필드명은 소문자로 만들어준다.

1. gob

    - Gob는 go에서만 읽고 쓸 수 있는 형태
    
    ```go
    func Gob() {
        var b bytes.Buffer
        enc := gob.NewEncoder(&b)
        data := map[string]string{"N":"J"}
        if err := enc.Encode(data); err != nil {
            fmt.Println(err)
        }

        dec := gob.NewDecoder(&b)
        var restored map[string]string
        if err := dec.Decode(&restored); err != nil {
            fmt.Println(err)
        }
        fmt.Prnitln(restored)
    }
    ```

## 3 Interface

- interface: set of method, 무언가를 할 수 있는 것 (...??)
- eg) ```io.Reader```: Read method를 정의하는 자료형들을 받을 수 있다, 읽을 수 있는 것.
- eg) Named Type에 String method를 만들어주면 문자열로 표현할 수 있는 것이 되고, ```fmt.Print```와 같은 함수에서 사용가능 해지고, ```fmt.Stringer```로 볼 수 있다.

1. 인터페이스의 정의

    ```go
    // InterfaceName: Loader
    // Method: Load
    type Loader interface {
        Load(filename string) error
    }

    // interface merge
    type ReadWriter{
        io.Reader
        io.Writer
    } // now, io.Reader & io.Writer 산하의 모든 method는 ReadWriter가 된다.
    ```

1. Custom Printer

    ```go
    func (t Task) String() string {
        check := "v"
        if t.Status != DONE {
            check = " "
        }
        return fmt.Sprint("[%s] %s %s", check, t.Title, t.Deadline)
    }
    ```

1. sort, heap

    - ```sort.Sort```: quick sort
    - ```sort.Interface```
    - http://golang.org/pkg/sort/

    1. 정렬 인터페이스의 구현

    1. 정렬 알고리즘

        - quicksort
            - variation: random pivot, sample pivot

        - n <= 7: insertion sort

    1. 힙

        - https://golang.org/pkg/container/heap/
        - 힙정렬은 선택 정렬과 더불어 정렬이 실행되는 도중에 첫 자료를 받아볼 수 있다.
        
1. 외부 의존성 줄이기

1. 빈 인터페이스와 형 단언

1. 인터페이스 변환 스위치