package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

// 期待通りにHTTPサーバーが起動しているか
// テストコードから意図通りに終了するか

func TestRun(t *testing.T) {
	t.Skip("リファクタリング中")
	// ポート番号に0を指定すると利用可能なポートを動的に選択してくれる
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	// 1.キャンセル可能な「context.Context」のオブジェクトを作る
	ctx, cancel := context.WithCancel(context.Background())
	// 2.別ゴルーチンでテスト対象の「run」関数を実行してHTTPサーバーを起動する
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	// 3.エンドポイントに対してGETリクエストを送信する
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// どんなポート番号でリッスンしているのか確認
	t.Logf("try request to %q", url)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// 6.GETリクエストで取得したレスポンスボディが期待する文字列であることを検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// 4.「cancel」関数を実効する run関数に終了通知を送信する
	cancel()
	// 5.「*errgroup.Group.Wait」メソッド経由で「run」関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
