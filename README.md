# go-clean-architecture

I practice the clean architecture with Go.

## Description of each layer

**Summary**  
Direction of dependence:  
Controller -> UseCase <- Service <- Repository

- Model
  domain model

- DB
  DBconnection, DBmock and so on.

えー、ここからは日本語で書きます

### Controller

外部と内部の変換を行う層

もし必要であればミドルウェアに関する処理を新しい層で用意してもいいし、この層でパッケージを切ってもいい

ちなみに、`gin.Context` を下の層に持っていっていないのはフレームワークへの依存をこの層だけで抑えるため  
仮にコンテキストが必要なら標準のコンテキストを下の層に持たせて、変換処理を書く想定

### UseCase

ドメインロジックを書く層  
アプリケーションサービスとも呼ばれる

ここでは以下のような定義とする  
ドメインロジック := 「問題解決となる処理をそのままプログラミング言語に変換したような処理」  
本当ならエラー処理もなるべく行いたくない

Service層のパーツを使って構築するイメージ

### Service

ドメインサービスを書く層

仕様上これ以上分割できない単位の操作を置く

正直、削ってもいい層ではあるがあったほうが整理できて嬉しい感じがする

命名が少しむずかしい

### Repository

APIを叩いたり、DBにクエリを投げたりする層

repository/api, repository/db みたいにすべきかもしれない

### Model

ドメインモデルを置く層

ここでは以下のような定義とする  
ドメインモデル := 「問題解決のために、物事の特定の側面を抽象化したもの」

例えば、ユーザーだと以下のような情報を持つとする

- ID
- 名前

ただ、実際には他にもいろんな情報が存在する

- 年齢
- ユーザーを作成した日付
- ユーザーの気持ち

これらの情報を取捨選択することが、「抽象化」

ドメインモデル貧血症などに注意

#### なぜ `*model.User` で扱っていないのか

- Q. 返り値などで `model.User` とし、存在判定を `user.IsEmpty()` としているのはなぜか  
  ポインタ型にすれば `user == nil` ともでき、良いのではないか

- A. ポインタ型のほうがGCとの兼ね合いなどにより遅くなるため  
  ただし、大きな構造体や大きな配列はポインタで返すべき  
  詳しくはググる

### DB

DBとのコネクションだったり、モックだったりをここに書いて、Repositoryにある構造体のフィールドに持たせる予定

また、DBから作成したモデルなどがここに入る予定

[gorm](https://github.com/xo/xo) は 遅い&&複雑なクエリが難しそう という理由から却下

ORM としては [sqlx](https://github.com/jmoiron/sqlx) あたりを使いたい  
ただ、テーブル定義から自分でモデルやINSERTクエリを書くのはちょっと大変そうなのでそれらだけ自動生成したい

一応想定している選択肢としては、[xo](https://github.com/xo/xo) を使ってDBからモデル生成して、それをsqlxと合わせて使う
