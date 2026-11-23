# 対訳表

ユビキタス言語（日本語）と、Go の実装名の対応を定めます。型名・フィールド名・関数名は、
この表に従って記述します。表に無い用語が必要になった場合は、名前を定める前に確認してください。

対象は、`specs/features/coupon.feature` のうち、背景と
「購入金額が最低購入金額に 1 円足りない場合はクーポンを使えない」シナリオに必要な範囲です。
シナリオを追加するたびに、この表も更新します。

| ユビキタス言語 | 分類 | 型 | フィールド・関数 | 備考 |
|---|---|---|---|---|
| 金額 | 値オブジェクト | `Money` | `NewMoney(int64)` | 負の値を許さない。コンストラクタ関数は `(Money, error)` を返し、引数の名前は `amount` |
| 会員ランク | 値オブジェクト | `MembershipRank` | `MembershipRankOf(string)` | 一般 `General`・ゴールド `Gold`・プラチナ `Platinum` の3つを `MembershipRank` 型のパッケージ変数として持つ。`MembershipRankOf` は日本語の名称から値を得る。戻り値は `(MembershipRank, error)`、引数の名前は `displayName` |
| 有効期限 | 値オブジェクト | `ExpirationDate` | `NewExpirationDate(time.Time)` | 当日まで有効。引数の名前は `date` |
| 券種コード | 値オブジェクト | `CouponCode` | `NewCouponCode(string)` | クーポンを識別する。戻り値は `(CouponCode, error)`、引数の名前は `value` |
| 適用結果 | 値オブジェクト | `CouponApplicationResult` | | 成否・エラーメッセージ・割引額・支払金額 |
| クーポン | エンティティ | `Coupon` | `NewCoupon(CouponCode, Money, Money, ExpirationDate)` | 識別子は券種コード。ポインタ `*Coupon` で扱う。引数の名前は `couponCode`、`minimumPurchaseAmount`、`maximumDiscountAmount`、`expirationDate` |
| 買い物かご | エンティティ | `Cart` | `NewCart(Money, MembershipRank)` | ポインタ `*Cart` で扱う。引数の名前は `totalAmount`、`membershipRank` |
| 割引計算 | ドメインサービス | `DiscountCalculationService` | `NewDiscountCalculationService()` | ポインタ `*DiscountCalculationService` で扱う |
| クーポンを適用する | メソッド | `DiscountCalculationService` | `Apply(*Coupon, *Cart, time.Time)` | 戻り値は `CouponApplicationResult`、引数の名前は `coupon`、`cart`、`now`、ステップ定義はこれを呼び出す |
| クーポンが使えるか | メソッド | `DiscountCalculationService` | `CanApply(*Coupon, *Cart, time.Time)` | 戻り値は `bool`、引数の名前は `coupon`、`cart`、`now` |
| 業務ルール違反 | エラー | `DomainError` | `NewDomainError(string)`、`WrapDomainError(string, error)`、`Error()`、`Unwrap()` | `error` インタフェースを実装する。値オブジェクトやエンティティの制約に反する値を渡された場合に、`*DomainError` を `error` として返す。原因となるエラーを包む場合は `WrapDomainError` を使用し、`Unwrap` で原因を返す |
| 券種コード | フィールド | `Coupon` | `couponCode`、`CouponCode()` | 型は `CouponCode` |
| 最低購入金額 | フィールド | `Coupon` | `minimumPurchaseAmount`、`MinimumPurchaseAmount()` | 型は `Money` |
| 割引上限額 | フィールド | `Coupon` | `maximumDiscountAmount`、`MaximumDiscountAmount()` | 型は `Money` |
| 有効期限 | フィールド | `Coupon` | `expirationDate`、`ExpirationDate()` | 型は `ExpirationDate` |
| 同一のクーポンか | メソッド | `Coupon` | `Equals(*Coupon)` | 券種コードが等しければ同一とみなす。戻り値は `bool`、引数の名前は `other` |
| 購入金額 | フィールド | `Cart` | `totalAmount`、`TotalAmount()` | 型は `Money` |
| 会員ランク | フィールド | `Cart` | `membershipRank`、`MembershipRank()` | 型は `MembershipRank` |
| 会員ランクの名称 | フィールド | `MembershipRank` | `displayName`、`DisplayName()` | 型は `string`（一般・ゴールド・プラチナ）、`MembershipRankOf` はこの値で探索する |
| 成否 | フィールド | `CouponApplicationResult` | `successful`、`Successful()` | 型は `bool` |
| エラーメッセージ | フィールド | `CouponApplicationResult` | `errorMessage`、`ErrorMessage()` | 型は `string`（成功時は空文字列） |
| 割引額 | フィールド | `CouponApplicationResult` | `discountAmount`、`DiscountAmount()` | 型は `Money` |
| 支払金額 | フィールド | `CouponApplicationResult` | `paymentAmount`、`PaymentAmount()` | 型は `Money` |
| 金額の値 | フィールド | `Money` | `amount`、`Amount()` | 型は `int64`（円単位） |
| 券種コードの値 | フィールド | `CouponCode` | `value`、`Value()` | 型は `string` |
| 有効期限の日付 | フィールド | `ExpirationDate` | `date`、`Date()` | 型は `time.Time`（時刻は 0 時、タイムゾーンは `time.UTC`） |

フィールドの行の「フィールド・関数」列は、非公開のフィールド名と、その値を読み取る公開メソッド名です。

## ステップ定義

ステップ定義は、`specs/steps` パッケージの `CouponSteps` 構造体のメソッドとして記述し、
同じパッケージの `InitializeScenario(ctx *godog.ScenarioContext)` で godog に登録します。
登録には、ステップの種別に応じて `ctx.Given`、`ctx.When`、`ctx.Then` を使用します。
正規表現とメソッド名は次の表に従います。`"([^"]*)"` と `(\d+)` は godog が引数に変換する箇所で、
それぞれ `string` と `int` の引数になります。メソッドの戻り値は `error` とします。

| 種別 | 正規表現 | メソッド名 | 引数 |
|---|---|---|---|
| 前提 | `^クーポン "([^"]*)" が存在する$` | `CouponExists` | `couponCode string` |
| 前提 | `^クーポン "([^"]*)" の最低購入金額は (\d+) 円である$` | `CouponHasMinimumPurchaseAmount` | `couponCode string`、`minimumPurchaseAmount int` |
| 前提 | `^クーポン "([^"]*)" の割引上限額は (\d+) 円である$` | `CouponHasMaximumDiscountAmount` | `couponCode string`、`maximumDiscountAmount int` |
| 前提 | `^クーポン "([^"]*)" の有効期限は "([^"]*)" である$` | `CouponExpiresOn` | `couponCode string`、`expirationDate string` |
| 前提 | `^割引額の端数は切り捨てる$` | `DiscountFractionIsTruncated` | なし |
| 前提 | `^ポイントは (\d+) ポイントを (\d+) 円として計算する$` | `PointRateIs` | `points int`、`yen int` |
| 前提 | `^現在日時は "([^"]*)" である$` | `CurrentDateTimeIs` | `now string` |
| 前提 | `^ポイントを利用しない$` | `PointsAreNotUsed` | なし |
| 前提 | `^会員ランクは "([^"]*)" である$` | `MembershipRankIs` | `membershipRank string` |
| 前提 | `^カートの購入金額は (\d+) 円である$` | `CartTotalAmountIs` | `totalAmount int` |
| もし | `^クーポン "([^"]*)" を適用する$` | `ApplyCoupon` | `couponCode string` |
| ならば | `^クーポンの適用は失敗する$` | `CouponApplicationFails` | なし |
| ならば | `^エラーメッセージ "([^"]*)" が表示される$` | `ErrorMessageIsShown` | `errorMessage string` |
| ならば | `^割引額は (\d+) 円である$` | `DiscountAmountIs` | `discountAmount int` |
| ならば | `^支払金額は (\d+) 円である$` | `PaymentAmountIs` | `paymentAmount int` |

メソッド名は、前提では状況を表す名前（`CouponExists` など）、もしでは動詞から始まる名前（`ApplyCoupon`）、
ならばでは検証する内容を表す名前（`CouponApplicationFails` など）とします。

`.feature` で「かつ」と記述されたステップは、直前のステップと同じ種別（`ctx.Given` または `ctx.Then`）で
登録します。金額は `(\d+)` で受け取り、`int` の引数とします。`Money` への変換はステップ定義の中で行います。

## 補足

**1つの用語が1つの型になるとは限りません。**
最低購入金額、購入金額、支払金額、割引額は、概念としてはいずれも金額です。違いは、
どのような役割で用いられているかだけです。したがって、型は `Money` のみとし、役割はフィールド名で表します。
別の型に分けるのは、その概念だけが持つ制約がある場合や、取り違えを型で防ぎたい場合に限られます。

**.feature に現れない構成要素もあります。**
「ならば クーポンの適用は失敗する」「かつ エラーメッセージ … が表示される」を検証するには、
適用した結果を受け取る型が必要です。`CouponApplicationResult` がこれにあたります。

**値オブジェクトの内部も、この表で定義します。**
`Money` や `CouponCode` は、型名だけでは実装できません。保持する値まで定義して、
はじめて構造体として記述できます。金額は円単位の整数で扱えるため `int64` とし、
整数のみの計算にしています。

**値オブジェクトとエンティティは、フィールドを公開しない構造体にしています。**
Go にはクラスがなく、構造体と、構造体に付けるメソッドでオブジェクトを表現します。
値オブジェクトは、フィールドを公開しない構造体とし、`NewMoney` のようなコンストラクタ関数で生成し、
`Amount()` のような読み取り用のメソッドで値を取り出します。フィールドを公開しないのは、
生成後に値を書き換えられないようにするためです。値オブジェクトのフィールドはすべて比較可能な型とし、
`==` で同値性を判定できるようにします。エンティティも同じ形の構造体ですが、識別子で同一性を判定するため
ポインタで扱い、`Equals` メソッドで券種コードを比較します。

**会員ランクは、パッケージ変数で表現しています。**
取りうる値は一般・ゴールド・プラチナの3つに限られますが、Go には列挙型がなく、定数で代用すると
日本語の名称を保持できません。そこで、`General`、`Gold`、`Platinum` の3つを `MembershipRank` 型の
パッケージ変数として持ち、コンストラクタ関数を公開しないことで、これ以外の値を生成できないようにします。
`.feature` には「一般」と日本語で記述されていますので、日本語の名称から値を得る `MembershipRankOf(string)` を
用意します。この変換をステップ定義に記述すると、ステップ定義が業務の知識を持つことになるためです。

**現在日時は `Apply` の引数で受け取るようにしています。**
背景には「現在日時は "2026-09-20 10:00:00" である」が記述されています。この日時で有効期限を判定するには、
ドメインサービスが背景で指定された日時を使用する必要があります。ドメインサービスの内部で
`time.Now()` のように実際の現在時刻を取得する実装にすると、背景の指定とは無関係に
実行した日の日時で判定されてしまい、シナリオの結果が実行日によって変わります。

**業務上の失敗と、プログラムの誤りを分けます。**
処理が正常に完了しない場合には、業務上の失敗とプログラムの誤りの2種類があり、この2つを区別して扱います。
クーポンを使えないことは業務上の失敗で、`.feature` にも「ならば クーポンの適用は失敗する」と
記述されているとおり、業務としては正常な結果です。したがってエラーにはせず、`CouponApplicationResult` の
成否とエラーメッセージで呼び出し側に返します。一方、金額に負の値が渡されることはプログラムの誤りで、
仕様のどこにも記述されていません。Go には例外の仕組みがなく、呼び出し側に `error` を返して知らせるのが
標準の方法であるため、コンストラクタ関数が `*DomainError` を `error` として返し、呼び出し側が処理を中断します。
`panic` は使用しません。対訳表に「業務ルール違反」として `DomainError` の行があるのは、このエラーを表すためです。