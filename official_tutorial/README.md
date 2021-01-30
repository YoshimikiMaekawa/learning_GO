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
func FuncName(input_name input_type) return_type {}
```
#### 返り値が2つ以上の場合
```go
func FuncName(input_name input_type) (return_type_1, return_type_2, ...) {}
```

### errorの処理
標準装備のerrors moduleを使用する．
'errors.New()'や'nil'の型は`error`なので注意．
#### error ruturn
```go
return "", errors.New("error message")
```
#### successfull return
`nil`をreturnすることで，callerは関数の処理が成功したことを知ることができる．
```go
return return_value, nil
```

### logの処理
標準装備のlog moduleを使用する．
#### 接頭辞をつける
```go
log.SetPrefix("prefix")
```
#### フラグの初期化(?)
```go
log.SetFlags(0)
```
#### logをコンソールに表示し，プログラムを終了する
```go
log.Fatal("error message")
```