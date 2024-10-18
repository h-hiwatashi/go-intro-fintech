SQL 文サンプル
【SELECT 文とは】
SELECT 文とは先ほども伝えた通り「データベースから、データを取得する命令」です。この SELECT 文をしようすることで、データベースからデータを取得することが出来ます。

【SELECT 文の基本構文】
SELECT 文の基本的な構文は以下になります。

SELECT [取得したい要素] FROM [使用テーブル];
欲しい情報と、どこから取ってくるかを指定すれば取ってくることが出来ます。

【特定要素だけ表示する】
基本的な構文は以下になります。

実行命令:

SELECT name FROM users;
意味:

「users テーブル」から全データの「name」だけを取得

【「,」で複数指定】
「,」で区切ることで、複数の要素を取得可能です。次の例では「名前」と「レベル」のみを取得・表示しています。

実行命令:

SELECT name,email FROM users;
意味:

「users テーブル」から全データの「name, email」を取得

【* で全部指定】
取得したい要素の項目に*を入れることで、全ての要素を取得することが可能です。

実行命令:

SELECT \* FROM users
意味:

「users テーブル」から全データの「全要素」を取得。

【SELECT の基本構文(まとめ)】
SELECT [取得したい要素] FROM [使用テーブル];
この一文だけでデータを取ってくることが可能です。まずはこれを覚えておきましょう。

【「WHERE」で条件を絞る】
WHERE で条件設定を行うことができる。使い方は簡単で、先ほどの SELECT 文の後ろ WHERE [条件文]を追加するだけです。

SELECT [取得したい要素] FROM [使用テーブル] WHERE [条件文];
例えば先ほどあげた例「level30 以上のキャラクターだけ取得したり」を実行してみましょう。

実行命令:

SELECT \* FROM users WHERE name = "mike";
意味

「users テーブル」の「name が mike のデータ」の「全要素」を取得

【INSERT 文で新規追加】
データの追加は INSERT 文で行います。

INSERT INTO [データを追加したいテーブル名] (要素名 A,要素名 B) VALUES(要素 A に入れるデータ,要素 B に入れるデータ);
例えば今回 user を増やすために INSERT を行うなら、以下のようになる。

実行命令:

INSERT INTO users (id, name) VALUES(1,"joe");
users テーブルに id=1, name="joe"のデータを追加。

実行命令:

SELECT \* FROM users;
を実行すると、データが追加されているのが確認できる。

【UPDATE 文で変更】
データの更新は UPDATE 文を使う。

実行命令:

UPDATE users SET name = "nancy" WHERE id = 1;
意味:

「users テーブル」の「id = 1 のデータ」の「name を"nancy"」に更新

【DELETE 文で削除】
データの削除は DELETE 文で使う。

実行命令:

DELETE FROM users WHERE id = 1;
意味:

「users テーブル」の「id が１のデータ」を削除する
