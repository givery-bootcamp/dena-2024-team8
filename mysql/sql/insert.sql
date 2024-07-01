INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');

INSERT INTO users (name, password) VALUES ('taro', '$2a$12$2cv5O.FaITyNSBqlS./zKOotzB8AuWIojHVjLvz/e3jiEn0xYAnBS');
INSERT INTO users (name, password) VALUES ('hanako', '$2a$12$4bTKrmghvw8FFfTyb/8S/O2GGD6WsKHOnv/DUuFKhg3TO8dHxkj.W');

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body) VALUES (2, 'test2', '質問2\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, '今日の料理：簡単パスタレシピ', '今日は忙しい人のための簡単パスタレシピを紹介します。\n材料は、パスタ、ツナ缶、玉ねぎ、オリーブオイルだけ！');
INSERT INTO posts (user_id, title, body) VALUES (1, '週末の映画感想：「星空の向こう側」', '先週末に見た「星空の向こう側」という映画の感想です。\n予想以上に感動的で、特に主演の演技が素晴らしかったです。');
INSERT INTO posts (user_id, title, body) VALUES (1, '初心者でも簡単！ガーデニングのコツ', 'ガーデニング初心者の方へ。\n今回は、失敗しにくい植物の選び方と、水やりの頻度について解説します。');
INSERT INTO posts (user_id, title, body) VALUES (1, '我が家の猫、モモの近況報告', 'モモが我が家に来て1年が経ちました。\n最近は窓際の日向ぼっこが日課になっているようです。癒されます。');
INSERT INTO posts (user_id, title, body) VALUES (1, '効率的な在宅ワークのコツ', '在宅ワークを始めて半年が経ちました。\n今日は、集中力を保つためのタイムマネジメント術について共有します。');
INSERT INTO posts (user_id, title, body) VALUES (1, '旅行記：京都の隠れた名所', '先月の京都旅行で見つけた、観光客があまり知らない素敵なスポットを紹介します。\n特に、夕暮れ時の東山の景色は絶景でした。');
INSERT INTO posts (user_id, title, body) VALUES (1, 'プログラミング学習日記：Day 30', 'プログラミング学習を始めて1ヶ月が経ちました。\n今日は初めて自力でWebアプリを作成できて、とても嬉しいです！');
INSERT INTO posts (user_id, title, body) VALUES (1, '健康的な朝食のすすめ', '忙しい朝でも簡単に作れる、栄養満点の朝食レシピを紹介します。\n今日のテーマは、オートミールを使ったアレンジレシピです。');
INSERT INTO posts (user_id, title, body) VALUES (1, '読書感想：「明日への一歩」', '最近読んだ自己啓発本「明日への一歩」の感想です。\n特に印象に残ったのは、小さな習慣の積み重ねの大切さについての章でした。');
INSERT INTO posts (user_id, title, body) VALUES (1, '初心者ランナーの挑戦：5km走破！', '今日、ついに5km走破に成功しました！\n3ヶ月前は1kmも走れなかったのに、少しずつ距離を伸ばしてきて良かったです。');

INSERT INTO comments (post_id, user_id, body) VALUES (1, 2, '回答1\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (1, 1, '回答2\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 2, '回答3\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 1, '回答4\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 2, '回答5\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 1, '回答6\n改行');
INSERT INTO comments (post_id, user_id, body) VALUES (2, 1, '回答6\n改行');
