USE clubone;

INSERT INTO wikis (creator, create_time, update_time, title, content, uuid)
VALUES (2, NOW(), NOW(), 'Test Title2', '<p>2Test content for the wiki entry.</p>', UUID());
