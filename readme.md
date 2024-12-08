# greeting

あいさつするサービス。https://github.com/naoyafurudono/clio-go を用いるサンプルCLI。

greetサービス https://connectrpc.com/docs/go/getting-started がここで紹介されている。
チュートリアル通りに作成したのがこのgreetingサービス。

チュートリアルではconnectサーバを作成できるが、ここではそれをCLIにする。ビジネスロジックを実装するサービスには手を加えず、インターフェースをconnectサーバからCLIに置き換える。
