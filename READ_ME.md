# Prizz API

## 概要

Prizz-APIは現在行われている懸賞のデータを提供するAPIです。

 

## 目次

[TOC]

## 更新履歴

[2019/08/27]

* `is_oneclick`, `twitter_way`, `twitter_raw`の追加
* `/way/*`のパラメータにorderを追加

 

## 使用例

例）期限が近い懸賞を50個取得する場合

https://api.prizz.jp/deadline?limit=50

例）応募方法がTwitterの懸賞をTwitterAPIを含めて10個取得する場合

https://api.prizz.jp/way/twitter?limit=10&raw=true

 

 

## フィールド

* `id` : 一意な整数

* `name` : 懸賞の名前

* `winner` : 当選人数

* `image_url` : 画像のURL

  ~~`image_bin` : BASE64形式の画像（廃止）~~

* `created_at` : 登録日

* `updated_at` : 更新日

* `limit_date` : 応募締め切り日時

* `link` : 懸賞フォームのURL

* `provider` : 提供元

* `way` : 応募方法

* `category` : カテゴリ一覧を参照。配列で返る。

* `is_oneclick` : Twitterのワンクリック懸賞に対応しているかどうか。

* `twitter_way` : witterの懸賞に応募する方法。配列で返る。

* `twitter_raw` : Twitter-APIのそのままのjson。

 

 

## エンドポイント

以下、ホストはすべて `api.prizz.jp` です。

 

### レファレンス

エンドポイント



```

```









```

```



このレファレンスを表示できます。

 

 

### 期限が近い順

エンドポイント



```

```









```

```



パラメータ

* `limit` : 返り値の上限を指定します。（デフォルトで10）
* `raw` : `twitter_raw`を返すかどうかを指定します。（デフォルトでfalse）

 

 

### 新着順

エンドポイント



```

```









```

```



パラメータ

* `limit` : 返り値の上限を指定します。（デフォルトで10）
* `raw` : twitter_rawを返すかどうかを指定します。（デフォルトでfalse）

 

 

### 当選人数順

エンドポイント



```

```









```

```



パラメータ

* `limit` : 返り値の上限を指定します。（デフォルトで10）
* `raw` : twitter_rawを返すかどうかを指定します。（デフォルトでfalse）

 

 

### 応募方法別

エンドポイント



```

```









```

```



クエリ

* `way` : 応募方法を指定します。
  * `twitter` : Twitter

パラメータ

* `limit` : 返り値の上限を指定します。（デフォルトで10）
* `order` : 返り値の並びを指定します。（デフォルトでdeadline）
  * `deadline` : 期限が近い順
  * `new` : 新着順
  * `winner` : 当選者順
* `raw` : twitter_rawを返すかどうかを指定します。（デフォルトでfalse）

 

 

### カテゴリ別

エンドポイント



```

```









```

```



クエリ

* `category` : 指定できるカテゴリ名は*カテゴリ一覧*を参照してください。

パラメータ

* `limit` : 返り値の上限を指定します。（デフォルトで10）
* `order` : 返り値の並びを指定します。（デフォルトでdeadline）
  * `deadline` : 期限が近い順
  * `new` : 新着順
  * `winner` : 当選者順
* `raw` : twitter_rawを返すかどうかを指定します。（デフォルトでfalse）

 

 

## カテゴリ一覧

|      |      |
| :--- | :--- |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |
|      |      |

 

 

## 注意事項

* このAPIは趣味で作成したものであり、安定性は保証されません。
* SQLインジェクション等はやめてください。脆弱性を見つけてしまった場合は直ちに[みちのすけ](https://twitter.com/Michin0suke)にご報告ください。
* 仕様は予告なしに変更されることがあります。バージョン管理はされていません。