# Adapter Pattern이란
 - 이전 화면에있는 GoF 디자인 패턴에서의 구조(Structural) 패턴에 해당하는 패턴중 하나이다.
 - Adapter Pattern은  서로 다른 구조체간의 연결을 하기 위한 패턴으로 interface를 활용하여 연결하게 된다.

## 장점
 - Adapter 패턴의 가장 큰 장점을 기존 코드를 변경하지 않아도 된다는 것이다.
 - 또한 기존 코드를 변경하지 않기 때문에 클래스 재활용성을 증가시킬 수 있다.
 
## 단점
 - 구성요소를 위해 클래스를 증가시켜야 하기 때문에 복잡도가 증가할 수 있다.
 - 클래스 Adapter 의 경우 상속을 사용하기 때문에 유연하지 않다.
 - 반면 객체 Adapter 의 경우는 대부분의 코드를 다시 작성해야 하기 때문에 효율적이지 않다.

## 메모
 - go의 <code>http</code> 관련 내부 패키지중 <code>HandlerFunc</code>라는 타입은 어댑터 패턴으로 구현되어 있다.
 -  <code>HandlerFunc</code>의 경우 <code>ServeHTTP</code>가 구현되어있으므로 <code>HandlerFunc</code>로 감싼 타입은 <code>Handler</code> 타입이 적용되기 때문에 어댑터의 역할을 하게 된다.
 - <code>http.Handler</code>는 <code>middleware</code>를 만들 때 많이 사용하게 된다.
 - <code>middleware</code>의 예제는 아래의 사이트를 참고하면 좋을것같다.
 ###### <code>http.Handler</code>를 이용한 <code>middleware</code> 예제 :  https://up-to-date-items.tistory.com/226
