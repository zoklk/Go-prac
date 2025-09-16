package main

import (
	// 1. 우리가 만든 로컬 패키지
	"github.com/zoklk/Go-prac/helloGo/greeting"

	// 2. go get으로 설치한 외부 라이브러리
	"github.com/rs/zerolog/log"
)

func main() {
	// greeting 패키지의 Hello 함수 호출
	message := greeting.Hello()

	// zerolog를 사용해 구조화된 로그 출력
	log.Info().Msg(message)
}
