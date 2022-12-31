- アクセストークンの署名用の秘密鍵、公開鍵を保存するディレクトリです
- 以下の鍵生成コマンド(macOS)を実行し、ディレクトリ配下に設置して下さい
```
openssl genrsa 4096 > secret.pem
openssl rsa -pubout < secret.pem > public.pem
``` 