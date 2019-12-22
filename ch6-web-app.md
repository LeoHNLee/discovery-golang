# Chapter6: Web Application

## 1 Hello, 世界!

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello, 世界!")
    }
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 2 To Do List

1. RESTful API

    - Client - Server
    - no state
    - enable Cache
    - URI: 접근 자원 선택
    - get / put / post / delete

1. Data Access Object

    - data access object 패턴
        - db에 필요한 연산들을 추상 인터페이스로 만들어서 사용
        - 사용하는 db와 비즈니스 로직을 구분할 수 있다.
        - 단점...도 많은데 알아서 느껴보자

    ```go
    // ID is a data type to identify a task
    type ID string

    // DataAccess is an interface to access tasks
    type DataAccess interface {
        Get(id ID) (task.Task, error)
        Put(id ID, t task.Task) error
        Post(t task.Task) (ID, error)
        Delete(id ID) error
    }

    // MemoryDataAccess is a simple in-memory database
    type MemoryDataAccess struct {
        tasks map[ID]task.Task
        nextID int64
    }

    // NewMemoryDataAccess returns a new MemoryDataAccess
    func NewMemoryDataAccess() DataAccess {
        return &MemoryDataAccess{
            tasks: map[ID]task.Task{},
            nextID: int64(1),
        }
    }

    // ErrTaskNotExist occurs when the task with the Id was not found.
    var ErrTaskNotExist = errors.New("task does not exist")

    // Get returns a task with a given ID.
    func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
        t, exists := m.tasks[id]
        if !exists {
            return task.Task{}, ErrTaskNotExist
        }
        return t, nil
    }

    // Put updates a task with a given ID with t
    func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
        if _, exists := m.tasks[id]; !exists {
            return ErrTaskNotExist
        }
        m.tasks[id] = t
        return nil
    }

    // Post adds a new task.
    func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
        id := ID(fmt.Sprint(m.nextID))
        m.nextID++
        m.tasks[id] = t
        return id, nil
    }

    // Delete removes the task with a given ID
    func (m *MemoryDataAccess) Delete(id ID) error {
        if _, exists := m.tasks[id]; !exists {
            return ErrTaskNotExist
        }    
        delete(m.tasks, id)
        return nil
    }
    ```
1. RESTful API 핸들러 구현

1. Restful 서버 완성

1. HTML 템플릿 작성

## 3 code refactoring

1. 통일성 있게 파일 나누기

1. 라우터 사용하기

## 4 do more

1. HTTP 파일 서버

    - js, css, html 등 파일 분리
    - http.FileServer

    ```go
    var (
        addr = flag.String(
            "addr",
            ":8080",
            "address of the webserver",
        )
        root = flag.String(
            "root",
            "/var/www",
            "root directory",
        )
    )
    
    func main() {
        flag.Parse()
        log.Fatal(http.ListenAndServe(
            *addr,
            http.FileServer(http.Dir(*root))
        ))
    }
    ```

1. 몽고디비와 연동

    ```shell
    go get gopkg.in/mgo.v2
    ```
    
    ```go
    // MongoAccessor is an Accessor for MongoDB
    type MongoAccessor struct {
        session *mgo.Session
        collection *mgo.collection
    }

    // New returns a new MongoAccessor
    func new(path, db, c string) task.Accessor {
        session, err := mgo.Dial(path)
        if err != nil {
            return nil
        }
        collection := session.DB(db).C(c)
        return &MongoAccessor{
            session: session,
            collection: collection,
        }
    }

    func (m *MongoAccessor) Close() error {
        m.session.Close()
        return nil
    }

    // idToObjectId returns bson.ObjectID converted from id.
    func idToObjectId(id task.ID) bson.ObjectId {
        return bson.ObjectIdHex(string(id))
    }
    
    // objectIdToID returns task.Id converted from objID
    func objectIdToID(objID bson.ObjectId) task.ID {
        return task.ID(objID.Hex())
    }

    // Get returns a task with a given ID
    func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
        return m.collection.UpdateId(idToObjectId(id), t)
    }

    // Post adds a new task
    func (m *MongoAccessor) Post(t task.Task) (task.ID, error) {
        objID := bson.NewObjectId()
        _, err := m.collection.UpsertId(objID, &t)
        return objectIdToID(objID), err
    }

    // Delete removes the task with a given ID
    func (m *MongoAccessor) Delete(id task.ID) error {
        return m.collection.RemoveId(idToObjectId(id))
    }
    ```

1. 에러 처리

    - 에러에 추가 정보 실어서 보내기

    ```go
    // Error interace type
    type error interface {
        Error() string
    }

    // other way
    fmt.Errorf("ID %d is negative", id)
    
    // define error type
    func (e ErrNegativeID) Error() string {
        return fmt.Sprintf("ID %d is negative", e)
    }

    // define error struct
    ```

    - 반복된 에러 처리 피하기

    ```go
    // Must function raise panic when raise error in neccessary working code
    func Must(err error) {
        if err != nil {
            panic(err)
        }
    }
    ```

    - 추가 정보와 함께 반복된 처리 피하기

    ```go
    type ResponseError struct {
        Err error
        Code int
    }

    resp := Response(
        ID: id,
        Task: t,
        Error: respErr,
    )
    json.NewEncoder(w).Encode(resp)
    w.WriteHeader(resp.Code)
    ```

    - panic, recover

        - panic: 호출 스택을 타고 역순으로 올라가서 최종적으로 프로그램 종료
        - recover: 상위 호출자로 panic 전파 방지

    ```go
    func f() (i int) {
        defer func() { // 이후 g() 패닉 발생과 무관하게 동작한다.
            if r:= recover(); r != nil {
                fmt.Println("recovered in f", r)
                i = -1 // panic이 일어난 경우에만 반환값을 -1로 보낼 수 있다.
            }
        }()
        g() // this function panics
        return 100 // 반환코드에 도달하기 전에 panic이 발생했기 때문에 코드가 실행되지 않고 디폴트 반환값이 반환된다.
    }

    func g() {
        panic("I panic!")
    }
    ```