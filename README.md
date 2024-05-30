counter appを作成する

## 仕様
現在のcountが表示されている
increment/decrement buttonがある
各buttonにより現在のcountが更新される


## MVPモデルについて

### 概要
MVP(Model-View-Presenter)

### Model
指示を受けて，自身を更新。
変更後，コールバックなどを行う。

### View
画面の描画担当。

### Presenter
Viewから移譲を受け、Modelに変更指示を送る。Modelからのコールバックを受けて，Viewを更新。


## 今sampleについて

