DELETE FROM categories WHERE id IN (1, 2, 3);
ALTER SEQUENCE categories_id_seq RESTART WITH 1;