# knight-tour-encoder

高信頼ソフトウェアの課題です。knight's tourの問題を、satソルバが解けるDimacs形式ファイルにエンコードします。
[https://github.com/uchijo/knight-tour-decoder](https://github.com/uchijo/knight-tour-decoder)を使うことで、結果をより見やすくすることが可能です。

## 使い方

```sh
go run main.go <グリッドの一辺の長さ> <結果の出力先>
```

## 注意点

- 正方形のグリッドにしか対応しておりません。
- 標準出力はデバッグ用です。無視して構いません。
- 初期位置は指定されていない状態で出力されます。初期位置を指定したい場合は、clauseを追加してください。
  - 例えば、初期位置を1としたい場合は、 `1 0` といった節を追加すればそのようになります。
  - Dimacsファイルの節の数を変更するのを忘れないように注意してください。
