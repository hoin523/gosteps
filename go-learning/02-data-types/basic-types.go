// basic-types.go - Go의 기본 데이터 타입
package main

import (
	"fmt"
	"unsafe" // 메모리 크기 확인용
)

func main() {
	fmt.Println("=== Go 기본 데이터 타입 ===")

	// 1. 정수형 (Integer Types)
	fmt.Println("\n1. 정수형:")

	// 부호 있는 정수
	var int8Val int8 = 127          // -128 ~ 127
	var int16Val int16 = 32767      // -32,768 ~ 32,767
	var int32Val int32 = 2147483647 // -2,147,483,648 ~ 2,147,483,647
	var int64Val int64 = 9223372036854775807

	// 플랫폼 의존적 정수 (32bit 시스템에서는 32bit, 64bit 시스템에서는 64bit)
	var intVal int = 12345

	fmt.Printf("int8: %d (크기: %d bytes)\n", int8Val, unsafe.Sizeof(int8Val))
	fmt.Printf("int16: %d (크기: %d bytes)\n", int16Val, unsafe.Sizeof(int16Val))
	fmt.Printf("int32: %d (크기: %d bytes)\n", int32Val, unsafe.Sizeof(int32Val))
	fmt.Printf("int64: %d (크기: %d bytes)\n", int64Val, unsafe.Sizeof(int64Val))
	fmt.Printf("int: %d (크기: %d bytes)\n", intVal, unsafe.Sizeof(intVal))

	// 부호 없는 정수
	var uint8Val uint8 = 255     // 0 ~ 255
	var uint16Val uint16 = 65535 // 0 ~ 65,535
	var uint32Val uint32 = 4294967295
	var uint64Val uint64 = 18446744073709551615
	var uintVal uint = 12345

	fmt.Printf("uint8: %d (크기: %d bytes)\n", uint8Val, unsafe.Sizeof(uint8Val))
	fmt.Printf("uint16: %d (크기: %d bytes)\n", uint16Val, unsafe.Sizeof(uint16Val))
	fmt.Printf("uint32: %d (크기: %d bytes)\n", uint32Val, unsafe.Sizeof(uint32Val))
	fmt.Printf("uint64: %d (크기: %d bytes)\n", uint64Val, unsafe.Sizeof(uint64Val))
	fmt.Printf("uint: %d (크기: %d bytes)\n", uintVal, unsafe.Sizeof(uintVal))

	// 2. 부동소수점형 (Floating Point Types)
	fmt.Println("\n2. 부동소수점형:")

	var float32Val float32 = 3.14159
	var float64Val float64 = 3.141592653589793

	fmt.Printf("float32: %f (크기: %d bytes)\n", float32Val, unsafe.Sizeof(float32Val))
	fmt.Printf("float64: %f (크기: %d bytes)\n", float64Val, unsafe.Sizeof(float64Val))
	fmt.Printf("float32 정밀도: %.10f\n", float32Val)
	fmt.Printf("float64 정밀도: %.15f\n", float64Val)

	// 3. 복소수형 (Complex Types)
	fmt.Println("\n3. 복소수형:")

	var complex64Val complex64 = 1 + 2i
	var complex128Val complex128 = 1.5 + 2.5i

	fmt.Printf("complex64: %v (크기: %d bytes)\n", complex64Val, unsafe.Sizeof(complex64Val))
	fmt.Printf("complex128: %v (크기: %d bytes)\n", complex128Val, unsafe.Sizeof(complex128Val))

	// 복소수 연산
	fmt.Printf("복소수 합: %v\n", complex64Val+complex64(complex128Val))
	fmt.Printf("실수부: %f, 허수부: %f\n", real(complex128Val), imag(complex128Val))

	// 4. 불린형 (Boolean Type)
	fmt.Println("\n4. 불린형:")

	var boolVal bool = true
	var falseBool bool = false

	fmt.Printf("bool true: %t (크기: %d bytes)\n", boolVal, unsafe.Sizeof(boolVal))
	fmt.Printf("bool false: %t\n", falseBool)

	// 불린 연산
	fmt.Printf("true && false: %t\n", boolVal && falseBool)
	fmt.Printf("true || false: %t\n", boolVal || falseBool)
	fmt.Printf("!true: %t\n", !boolVal)

	// 5. 문자열형 (String Type)
	fmt.Println("\n5. 문자열형:")

	var stringVal string = "안녕하세요, Go!"
	var emptyString string = ""

	fmt.Printf("문자열: %s (크기: %d bytes)\n", stringVal, unsafe.Sizeof(stringVal))
	fmt.Printf("문자열 길이: %d\n", len(stringVal))
	fmt.Printf("빈 문자열: '%s' (길이: %d)\n", emptyString, len(emptyString))

	// 문자열 연산
	greeting := "안녕" + "하세요!"
	fmt.Printf("문자열 연결: %s\n", greeting)

	// 6. 바이트형 (Byte Type) - uint8의 별칭
	fmt.Println("\n6. 바이트형:")

	var byteVal byte = 'A' // byte는 uint8의 별칭
	fmt.Printf("바이트: %c (숫자: %d)\n", byteVal, byteVal)

	// 7. 룬형 (Rune Type) - int32의 별칭, 유니코드 문자
	fmt.Println("\n7. 룬형:")

	var runeVal rune = '가' // rune은 int32의 별칭
	fmt.Printf("룬: %c (유니코드: %d)\n", runeVal, runeVal)

	// 문자열을 룬 슬라이스로 변환
	koreanText := "안녕하세요"
	runes := []rune(koreanText)
	fmt.Printf("한글 문자열 길이: %d\n", len(runes))
	for i, r := range runes {
		fmt.Printf("인덱스 %d: %c (유니코드: %d)\n", i, r, r)
	}

	// 8. 포인터형 (Pointer Type)
	fmt.Println("\n8. 포인터형:")

	var x int = 42
	var ptr *int = &x // x의 주소

	fmt.Printf("변수 x: %d\n", x)
	fmt.Printf("포인터 ptr: %p\n", ptr)
	fmt.Printf("포인터가 가리키는 값: %d\n", *ptr)
	fmt.Printf("포인터 크기: %d bytes\n", unsafe.Sizeof(ptr))

	// 포인터를 통한 값 변경
	*ptr = 100
	fmt.Printf("포인터로 값 변경 후 x: %d\n", x)
}

/*
Go의 타입 특징:
1. 강타입 언어 - 타입 안전성 보장
2. 정적 타입 - 컴파일 타임에 타입 결정
3. 타입 추론 - := 연산자로 자동 타입 추론
4. 제로값 - 모든 타입은 기본값을 가짐
   - 숫자: 0
   - 불린: false
   - 문자열: ""
   - 포인터, 슬라이스, 맵, 채널, 함수: nil

타입 변환:
- Go는 암시적 타입 변환을 지원하지 않음
- 명시적 변환 필요: int(floatValue)
*/
