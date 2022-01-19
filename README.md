## GorineeWeb 웹프레임워크 프로젝트
1. 나만의 웹프레임쿼으 만들기 프로젝트입니다.

## 사용 라이브러리
- github.com/valyala/fasthttp
- github.com/fasthttp/router
- github.com/prometheus/client_golang/prometheus/promhttp

## 해야할 것
1. 라우터 생성하기
2. 서버 실행하기
3. GET 메소드 만들기
4. POST 메소드 만들기
5. PUT 메소드 만들기
6. DELETE  메소드 만들기
7. 라우터 그룹핑 만들기
8. 속도 + 동시접속 가능하게 개선해보기
9. 웹소켓 코드 만들기
10. SSL 인증 관련 코드 만들기
11. 토큰 발급 코드 만들기

## 진행상황 및 피드백
- 현재 New를 통해 서버 실행되는 거 확인
- 서버 실행 시, GorineeWeb과 버전 출력 뛰어주기 완료
- 메소드와 연결되는 핸들러 만들기 완료
- Get Mapping 완료
- query and param 구현 완료
- middleware 구성 진행중
- 라우터 그룹핑 필요 