# Chapter8: 실무 패턴

- 새로운 언어를 배우는 잘못된 방법: 다른 언어로 하던 것을 Go로 어떻게 구현하지??
- 어떤 문제를 풀려고 하는가?!?!
- 문제를 푸는 가장 효과적인 방법은 언어마다 다를 수 있다.

## 1 Overloading

- 오버로딩이 필요한 경우
    - 자료형에 따라 다른 이름 붙이기
        - Go에서는 자료형에 따라 다른 함수 이름을 붙이자
    - 동일한 자료형의 자료 개수에 따른 오버로딩
        - Go에서는 가변인자를 사용하자
    - Go에서는 자료형 스위치 활용하기
        - 인터페이스로 인자를 받고, 메서드 내에서 자료형 스위치로 다른 자료형에 맞추어 다른 코드가 수행되게끔 한다.
    - 다양한 인자 넘기기: 기본값을 포함한 여러 설정을 넘기는 경우
        - Go에서는 설정들을 묶은 구조체로 넘겨서 해결해보자
    - Go에서는 인터페이스 활용하자

1. 연산자 오버로딩

    - Go는 지원해주지 않는다.
    - 편의를 위한 기능은 인터페이스로 해결

## 2 Template And Generic Programming

- Generic: 알고리즘을 표현하면서 자료형을 배제할 수 있는 프로그래밍 패러다임

1. Unit Test

    ```go
    // assertEqual
    func assertEqualString(t *testing.T, expected, actual string) {
        if expected != actual {
            t.Error("%s != %s", expected, actual)
        }
    }
    
    // reflect.DeepEqual 이용
    func assertEqual(t *testing.T, expected, actual interface{}) {
        if !reflect.DeepEqual(expected, actual) {
            t.Errorf("%v != %v", expected, actual)
        }
    }

    // table based Test

    // Example Test
    ```

1. 컨테이너 알고리즘: 인터페이스 활용

1. 자료형 메타 데이터

1. go generate

## 3 OOP

1. 다형성

1. 인터페이스

1. 상속

    - 메서드 추가

    - 오버라이딩

    - 서브 타입

1. 캡슐화

## 4 Design Pattern

1. 반복자 패턴

1. 추상 팩토리 패턴

1. 비지터 패턴