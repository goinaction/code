// counters 패키지를 정의한다.
package counters

// 알림 횟수를 저장하기 위한 정수 값을 저장하는
// alertCounter 타입을 비노출 타입으로 선언한다.
type alertCounter int

// 비노출 타입인 alertCounter 타입의
// 값을 생성하여 리턴하는 New 함수를 정의한다.
func New(value int) alertCounter {
	return alertCounter(value)
}
