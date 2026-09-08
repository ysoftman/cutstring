# cutstring

http 파라미터를 컬러로 구분하여 표시해줍니다. :smile:

## build

```bash
git clone https://github.com/ysoftman/cutstring
go get -u ./...
cd cutstring
go build
```

## installation

```bash
go install github.com/ysoftman/cutstring@latest
```

## example

```bash
cutstring "https://www.google.co.kr/search?q=%EC%B9%B4%EC%B9%B4%EC%98%A4&oq=%EC%B9%B4%EC%B9%B4%EC%98%A4&aqs=chrome..69i57j69i61j69i60l2j0j69i59.2266j0j7&sourceid=chrome&ie=UTF-8"
```
