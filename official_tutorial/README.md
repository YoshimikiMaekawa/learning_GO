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
#### 返り値が1つの場合
```go
func func_name(input_name input_type) return_type {}
```
#### 返り値が2つ以上の場合
```go
func func_name(input_name input_type) (return_type_1, return_type_2, ...) {}
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