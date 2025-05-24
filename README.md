# Todo CLI

シンプルなTodo管理CLIツールです。

## インストール

```bash
go install github.com/TakahashiShuuhei/todo@latest
```

## 使い方

### 初期化

```bash
todo init
```

### タスクの追加

```bash
todo add "タスクのタイトル" ["タスクの説明"]
```

説明は省略可能です。

### タスクの一覧表示

```bash
todo list
```

アーカイブされていないタスクのみを表示します。

アーカイブされたタスクも含めて表示する場合：

```bash
todo list all
```

### タスクの完了

IDを指定して完了状態をトグル：

```bash
todo complete <ID>
```

または、インタラクティブに選択：

```bash
todo complete
```

アーカイブされていないタスクのみが選択候補として表示されます。

### タスクのアーカイブ

IDを指定してアーカイブ：

```bash
todo archive <ID>
```

または、インタラクティブに選択：

```bash
todo archive
```

すべてのタスクをアーカイブ：

```bash
todo archive all
```

アーカイブされていないタスクのみがアーカイブされます。

## 表示形式

タスクの表示形式は以下の通りです：

- 完了状態: `✓`（完了）または` `（未完了）
- アーカイブ状態: `🗄`（アーカイブ済み）または` `（未アーカイブ）

例：
- `✓🗄 [1] タスク` - 完了済みでアーカイブされたタスク
- ` ✓ [2] タスク` - 未完了でアーカイブされたタスク
- `✓  [3] タスク` - 完了済みで未アーカイブのタスク
- `   [4] タスク` - 未完了で未アーカイブのタスク

## 開発

```bash
# 依存関係のインストール
go mod download

# ビルド
go build -o todo cmd/todo/main.go
``` 