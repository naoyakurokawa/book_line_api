package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/naoyakurokawa/book_line_api/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("failed to terminate server: %v", url)
	s := &http.Server{
		// 引数で受け取ったnet.Listenerを利用するのでAddrフィールドは指定しない
		// Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		// ListenAndServeメソッドではなく、Serveメソッドに変更する
		if err := s.Serve(l); err != nil &&
			// http.ErrServerClosed は http.Server.Shutdown()が正常に終了したことを示すため異常ではない
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		// if err := s.ListenAndServe(); err != nil &&
		// 	err != http.ErrServerClosed {
		// 	log.Printf("failed to close: %+v", err)
		// 	return err
		// }
		return nil
	})

	// チャネルから通知(終了通知)を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別ゴルーチンの終了を待つ
	return eg.Wait()
}
