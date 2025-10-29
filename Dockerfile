# =========================
# ビルド用ステージ
# =========================
FROM golang:1.21-alpine AS builder

# 作業ディレクトリを作る
WORKDIR /app

# 先にgo.modだけコピーして、依存ダウンロード（キャッシュ効くように分けるのが定石）
COPY go.mod ./
RUN go mod download

# 残りのソースをコピー
COPY . .

# 実行ファイルをビルド
RUN go build -o main .

# =========================
# 実行用ステージ（軽量イメージ）
# =========================
FROM alpine:3.20

WORKDIR /app

# ビルド成果物だけコピー
COPY --from=builder /app/main /app/main

# デフォルトコマンド
CMD ["/app/main"]
