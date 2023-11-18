CREATE DATABASE IF NOT EXISTS contents;
USE contents;

CREATE TABLE IF NOT EXISTS lp(
  viewer INT UNSIGNED NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS articles(
  id         VARCHAR(32)  PRIMARY KEY,
  title      VARCHAR(128) NOT NULL,
  content    TEXT         NOT NULL,
  viewer     INT UNSIGNED NOT NULL DEFAULT 0,
  favorites  INT UNSIGNED NOT NULL DEFAULT 0,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL
);

CREATE TABLE IF NOT EXISTS tags(
  id         VARCHAR(32) PRIMARY KEY,
  name       VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS article_tags(
  FOREIGN KEY (article_id) REFERENCES articles (id),
  FOREIGN KEY (tag_id)     REFERENCES tags     (id),
  PRIMARY KEY (article_id, tag_id)
);
