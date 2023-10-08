-- 作成したDBに接続

CREATE DATABASE test_db;

\c test_db;

-- テーブル作成

DROP TABLE IF EXISTS sample;

CREATE TABLE
    sample (
        id integer NOT NULL PRIMARY KEY,
        name varchar(100) NOT NULL,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- IDシーケンス作成

CREATE SEQUENCE sample_seq START 1;

-- サンプルデータ登録

INSERT INTO sample(id, name) VALUES (nextval('sample_seq'), 'test1');