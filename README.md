# bdd-ai-coupon-go

note の連載「BDD × AI」のサンプルコードです。Gherkin で記述した仕様（`.feature`）を、
Go と godog で実行します。

同じ `.feature` を他の言語でも実行します。言語ごとにリポジトリを分けていますが、
`.feature` の記述はすべて同一です。

| 言語 | リポジトリ |
|------|-----------|
| Java | [bdd-ai-coupon-java](https://github.com/fullness-MFurukawa/bdd-ai-coupon-java) |
| C# | [bdd-ai-coupon-cs](https://github.com/fullness-MFurukawa/bdd-ai-coupon-cs) |
| Go | bdd-ai-coupon-go（本リポジトリ） |
| TypeScript | bdd-ai-coupon-ts（予定） |

---

## 連載記事

| 記事 | タイトル |
|------|---------|
| 第2回 | [仕様がそのまま品質基準になる——AIと作るGherkin（ガーキン）入門](https://note.com/fullness_dev/n/n8afa12c67ee3) |
| Go編-① | [仕様から実装へ（Go編-①）——.feature を実行するコードを AI に生成させる](https://note.com/fullness_dev/n/nf6d92680c1e4) |

記事の公開時点のコードは、タグから取得できます。記事の手順に沿って実行する場合は、
対応するタグをチェックアウトしてください。

```bash
git clone https://github.com/fullness-MFurukawa/bdd-ai-coupon-go.git
cd bdd-ai-coupon-go
git checkout v1   # Go編-①の公開時点
```

| タグ | 状態 |
|------|------|
| `v1` | ドメインモデルのスケルトンとステップ定義のひな型。15ステップが定義済みで、最初のステップが pending で停止する |

---

## 題材

ECサイトのクーポン機能を題材としています。購入金額や会員ランク、有効期限、
ポイントとの併用可否によって、クーポンを使えるかどうかと割引額が決まります。

要件とシナリオの詳細は第2回の記事を参照してください。

---

## 実行環境

| 項目 | バージョン |
|------|-----------|
| Go | 1.27 |
| godog | 0.16.0 |

### VS Code の拡張機能

| 拡張機能 | 発行元 |
|---------|-------|
| Go | Go Team at Google |
| Cucumber | Cucumber（CucumberOpen） |

Go の拡張機能には、言語サポート、補完、テストの実行が含まれます。
Cucumber の拡張機能は、`.feature` の構文強調と補完を提供します。

---

## 実行方法

```bash
go test ./specs/ -v -count=1
```

`-v` を付けると、godog が出力する `.feature` の各ステップと実行結果が表示されます。
`-count=1` は、`go test` の結果のキャッシュを使わずに毎回実行するための指定です。
ソースコードと `.feature` のどちらも変更していない場合、この指定が無いと前回の結果がそのまま表示されます。

`specs/specs_test.go` は、`.feature` を読み込んで実行するテストの入口です。godog はテストクラスを
生成しないため、このファイルを手で用意しています。`Strict: true` の指定により、未定義や未実装の
ステップを含むシナリオは失敗として扱われます。この指定が無い場合、godog はそれらのシナリオを
失敗とせず、テストは成功として報告されます。

ビルドと静的検査は次のコマンドで行います。

```bash
go build ./...
go vet ./...
```

---

## ディレクトリ構成

```
bdd-ai-coupon-go/
├── docs/
│   └── glossary.md                       対訳表（ユビキタス言語と実装名の対応）
├── internal/
│   └── domain/
│       ├── model/                        値オブジェクト・エンティティ・DomainError
│       └── service/                      ドメインサービス
├── specs/
│   ├── features/                         .feature
│   ├── steps/                            ステップ定義
│   └── specs_test.go                     .feature を実行するテストの入口
├── CLAUDE.md                             Claude Code への指示（規約）
├── go.mod
└── go.sum
```

`internal` の下に置いたパッケージは、同じモジュールの中からしか参照できません。
ドメインモデルを外部のモジュールから直接参照させないための、Go の標準の仕組みです。

### 設計方針

詳細は [CLAUDE.md](CLAUDE.md) に記述しています。主な方針は次のとおりです。

- 値オブジェクトは、フィールドを公開しない構造体とし、コンストラクタ関数と読み取り用のメソッドを持たせます
- エンティティも、フィールドを公開しない構造体とし、ポインタで扱い、識別子を比較する `Equals` で同一性を判定します
- エンティティには、状態を変更するメソッドではなく、業務上の判断を表すメソッドを持たせます
- クーポンを使えないことは業務上の正常な結果として `CouponApplicationResult` で返し、
  業務上の制約に反する値は `DomainError` を `error` として返して報告します。`panic` は使用しません
- ドメイン層は Web フレームワークやデータベースのライブラリに依存させません

---

## 対訳表

ユビキタス言語は日本語、Go の型名は英語になります。訳し方が一つに決まらないため、
どの用語をどの名前で実装するかを [docs/glossary.md](docs/glossary.md) にまとめています。

他の言語のリポジトリにも同じ表を配置しています。言語ごとに変わるのは実装名の列だけです。

---

## ライセンス

MIT License

本リポジトリのコードは、自由に複製・改変してご利用いただけます。
