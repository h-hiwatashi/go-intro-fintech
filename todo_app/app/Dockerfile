# alpineはバージョンを実質的に固定にできないためDebianベースを使用する
FROM golang:1.19-bullseye
WORKDIR /work
# アップデートとgitのインストール
# goを扱うにはgitが必須であるがalpyneやDebianは軽量化のためgitが入っていない
RUN apt update && apt install -y git