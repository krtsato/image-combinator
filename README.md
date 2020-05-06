# image-combinator

複数の画像を連結して１枚の画像を生成するプログラム

## 前提

- 入力画像１枚は 500 × 500 以下の正方形である
- 入力画像は jpeg 形式である  
- 入力画像の整形は今後対応

## 機能

- 入力画像の整形（予定）
- 出力画像サイズ比の可変化（予定）
- 出力画像のファイル形式変更（予定）

## 出力画像のユースケース

- Twitter
  - ヘッダー 1500px : 500px = 3 : 1
  - １枚投稿
    - sm & pc 1024px : 576px = 16 : 9
- YouTube
  - 再生画面 1920px : 1080px = 16 : 9
  - サムネイル 1280px : 720px = 16 : 9

## 出力画像パターン

- 3 : 1
  - 画像素材 3 * 1 枚，padding なし
  - 画像素材 15 * 5 枚，padding なし
- 16 : 9
  - 画像素材 5 * 2 枚，padding あり
  - 画像素材 144 枚，padding なし
