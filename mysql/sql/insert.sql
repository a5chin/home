INSERT INTO articles (id, title, content, viewers, favorites) VALUES ('01H0F7PC287Q3C2XH9C575F9ZR', 'タイトル', '# コンテンツ', 1, 1);
INSERT INTO articles (id, title, content, viewers, favorites) VALUES ('home', 'MyHome', '# Who am I?', 2, 0);

INSERT INTO tags (id, name) VALUES ('01H0F7PC2AJBPZJH73Q0EBM45H', '初心者');

INSERT INTO article_tags (article_id, tag_id) VALUES ('01H0F7PC287Q3C2XH9C575F9ZR', '01H0F7PC2AJBPZJH73Q0EBM45H');
