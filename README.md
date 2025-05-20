# Todo CLI

シンプルなタスク管理CLIツールです。

## 機能

- タスクの追加
- タスクの一覧表示
- タスクの完了
- タスクの削除

## インストール

```bash
go install github.com/TakahashiShuuhei/todo/cmd/todo@latest
```

## 使用方法

```bash
# タスクの追加
todo add "タスクのタイトル" "タスクの説明"

# タスクの一覧表示
todo list

# タスクの完了
todo complete <タスクID>

# タスクの削除
todo delete <タスクID>
```

## 開発

```bash
# 依存関係のインストール
go mod download

# ビルド
go build -o todo cmd/todo/main.go
``` 