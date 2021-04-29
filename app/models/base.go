package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_app/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

//_ "github.com/mattn/go-sqlite3" と記述して、ドライバーをインポートしておきます
//以下のコードの中では使用しないので、_ としてインポートします
//ただし、ビルド実行時に、このドライバーを含めないとSQLにアクセスできないので、
//上記のように記述する必要があります

//グローバル変数を定義します
//*sql.DB は、構造体 sql.DB のポインタ型です
var Db *sql.DB

//err を宣言しておきます
var err error

//146 sqlite3 は使わないので以下の const 箇所はコメントアウトしたよ

/*
//テーブル名を宣言しておきます
//115
//tableNameTodo を宣言しておきます
//132
//tableNameSession を宣言しておきます
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

func init() {

	//146
	//Heroku の PostgreSQL のURLが、url に格納されます
	url := os.Getenv("DATABASE_URL")

	//url のリソースを取得して connection に格納します
	connection, _ := pq.ParseURL(url)

	//connection は文字列なのですが、"sslmode=require" を追加します
	connection += "sslmode=require"

	//第1引数 ドライバー名を指定します
	//第2引数 取得したリソースである connection を渡します
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	//146 sqlite3 は使わないので以下の sqlite3 のDB接続の箇所はコメントアウトしたよ

	/*
	//第1引数 ドライバー名を指定します
	//第2引数 DB名を指定します ファイル名を記述すると、現在のディレクトリの配下に、
	//(そのファイル名) が存在すれば、それをOpenします
	//存在しなければ、そのファイル名の SQL ファイルを作成して、Openします
	//[db]
	//driver = sqlite3
	//name = webapp.sql

	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//DB 内にテーブルを作成します
	//`` でコマンドを囲むと以下のように改行して記述できます
	//id INTEGER PRIMARY KEY AUTOINCREMENT, と記述して、
	//id をプライマリーキーに指定します
	//AUTOINCREMENT と記述して、id を自動で増分させていきます
	//uuid STRING NOT NULL UNIQUE, については、
	//uuid の NULL 値を禁止しています。また、値の重複も禁止しています
	//created_at DATETIME と記述して、DATETIME型を指定しています

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	//コマンドを実行します
	//結果を使用しないので、_, err と記述します
	Db.Exec(cmdU)

	//DB 内に todos テーブルを作成します
	//`` でコマンドを囲むと以下のように改行して記述できます
	//id INTEGER PRIMARY KEY AUTOINCREMENT, と記述して、
	//id をプライマリーキーに指定します
	//AUTOINCREMENT と記述して、id を自動で増分させていきます
	//created_at DATETIME と記述して、DATETIME型を指定しています

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	//コマンドを実行します
	//結果を使用しないので、_, err と記述します
	Db.Exec(cmdT)

	//132
	//DB 内に sessions テーブルを作成します
	//`` でコマンドを囲むと以下のように改行して記述できます
	//id INTEGER PRIMARY KEY AUTOINCREMENT, と記述して、
	//id をプライマリーキーに指定します
	//AUTOINCREMENT と記述して、id を自動で増分させていきます
	//uuid STRING NOT NULL UNIQUE, については、
	//uuid の NULL 値を禁止しています。また、値の重複も禁止しています
	//created_at DATETIME と記述して、DATETIME型を指定しています

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	//コマンドを実行します
	//結果を使用しないので、_, err と記述します
	Db.Exec(cmdS)
	*/

}

//110
//UUID を生成する関数を作成します
//NewUUID() 関数で、ユニークIDを生成する
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//110
//パスワードをハッシュ値に変換する関数を作成します
//[]byte(plaintext) と記述して、パスワードを、byte型の配列に変換します
//それを、sha1.Sum() に渡して、ハッシュ値を生成します
//ハッシュ値は 以下のようなbyte型の配列となります
//[93 76 160 113 87 169 180 16 96 118 48 108 207 44 227 144 23 84 227 199]
//そのハッシュ値を fmt.Sprintf() に渡して、人間が読みやすい 16 進数の文字列に変換します
//以下のような文字列になります
//cf23df2207d99a74fbe169e3eba035e633b65d94
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
