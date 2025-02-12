# 最新のGoイメージを取得
FROM golang:latest

# 環境変数を設定
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 作業ディレクトリを作成
WORKDIR /app

# Goのモジュールファイルをコピー
COPY go.mod go.sum ./

# モジュールの依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# ビルド（REPL実行用）
RUN go build -o my_lisp_interpreter main.go

# コンテナ起動時のデフォルトコマンド（REPLを起動）
CMD ["./my_lisp_interpreter"]
