# learning_GO
A learning repository of golang.

## Installation Guide
+ [Download and install](https://golang.org/doc/install)

## 基本コマンド
### packageの実行
packageがあるディレクトリに移動したうえで，
```go
go run .
```
または
```go
go run [package名]
```
***
### moduleを新規作成する
```go
go mod init
```
***
### packageのexeを作成
```go
go build
```
***
### test codeを実行する
```go
go test
```
+ 末尾が`_test.go`のファイルに存在する関数を実行
***
### packageをinstallする
```go
go install
```
+ このコマンドで，packageをインストールしたら，package名をコマンドラインで実行するだけで処理が開始される
    + コマンド化されるみたいな感じ...？．
***

## プログラムのひな型
```go
// package名
package hoge

// moduleのimport
import (
    "fmt"
)

func main() {

}
```

## 基本文法
### 変数の宣言
```go
member := value
```
***
### 文字列処理
#### 変数交じりの文字列の宣言
```go
str := fmt.Sprintf("Hello, %v.", value)
```
+ `%v`はvalueのv
***
### 配列の処理
#### 配列の宣言
```go
array := []value_type{
    value_0,
    value_1,
    value_2,
}
```
+ 最終indexの要素の後に，カンマをつけてもOK
#### 配列の(N + 1)番目の要素を取得
```go
array[N]
```
#### 配列のsizeを取得
```go
len(array)
```
***
### mapの処理
#### mapの宣言
```go
map_name := make(map[key_type]value_type)
```
***
### for文
#### rangeを使用した場合
```go
for index, value := range array {}
```
+ 二つある変数のうち，どちらかが必要なかったら，"_"で明示的に無視することができる
```go
for _, value := range array {}
```
***
### if文
#### 基本
```go
if hoge != fuga {}
```
***
### 関数の宣言
#### 返り値が1つの場合
```go
func FuncName(input_name input_type) return_type {}
```
#### 返り値が2つ以上の場合
```go
func FuncName(input_name input_type) (return_type_1, return_type_2, ...) {}
```
***
### errorの処理
+ 標準装備のerrors moduleを使用する．
+ 'errors.New()'や'nil'の型は`error`なので注意．
#### error ruturn
```go
return "", errors.New("error message")
```
#### successfull return
```go
return return_value, nil
```
+ `nil`をreturnすることで，callerは関数の処理が成功したことを知ることができる．
***
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
***
### random値の処理
#### 現在時刻に基づいて初期化
```go
rand.Seed(time.Now().UnixNano())
```
+ プログラム実行時に自動的に実行される`init()`の中に記述しておくのが無難
#### int型，0～Nまでのrandom値を取得
```go
rand.Intn(N)
```
***
### init関数
```go
init() {}
```
+ プログラム実行後，グローバル変数の宣言が終わったのちに自動で呼ばれる
***
### test code
+ `_test.go`と最後につくファイルを用意する
#### test function
##### ひな型
```go
func TestTestName(t *testing.T) {
    want := ideal_value
    value := Function()
    if value == want {
        t.Fatalf("debug message")
    }
}
```
+ 適当に書いていますが，理想の結果を`want`で宣言し，実際の結果と比較する，みたいなプロセスを踏むみたいな感じかな...？
***