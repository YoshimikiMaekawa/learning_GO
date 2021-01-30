# GO Official Tutorial
## 基本コマンド
### moduleを新規作成する
```go
go mod init
```
### `go.mod`の必要moduleを更新する(?)
```go
go build
```

## 基本文法
### 関数の宣言
```go
func func_name(member_name member_type) return_type {

}
```

### errorの処理
#### error ruturn
```go
return "", errors.New("error message")
```
#### successfull return
`nil`をreturnすることで，callerは関数の処理が成功したことを知ることができる．
```go
return return_value, nil
```