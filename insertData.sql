insert into articles (title, contents, username, nice, created_at) values ('firstPost', 'This is my first blog', 'fujihara', 2, now());

insert into articles (title, contents, username, nice) values ('2nd', 'Second blog post', 'fujihara', 4);

insert into comments (article_id, message, created_at) values (1, '1st comment yeah', now());

insert into comments (article_id, message) values (1, 'welcome');