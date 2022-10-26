TRUNCATE table users CASCADE;
TRUNCATE table cards CASCADE;
TRUNCATE table orders CASCADE;

INSERT INTO public.users
(id, username, "password", created_at, updated_at)
VALUES
    (99, 'an', 'an', '2022-03-14 14:00:00', '2022-03-14 14:00:00'),
    (100, 'nghia', 'nghia', '2022-03-15 14:00:00', '2022-03-15 14:00:00'),
    (101, 'thuy', 'thuy', '2022-03-16 14:00:00', '2022-03-16 14:00:00');
INSERT INTO public.cards
(id, "number", expired_date, user_id, cvv, balance, created_at, updated_at)
VALUES
    (99, '3301223454322203', '2023-03-22 00:00:00', 99, '999', 9999, '2022-03-14 14:00:00', '2022-03-14 14:00:00'),
    (100, '3301223454322204', '2023-03-23 22:00:00', 100, '100', 10000, '2022-03-15 14:00:00', '2022-03-15 14:00:00'),
    (101, '3301223454322205', '2023-03-24 22:00:00', 101, '101', 10001, '2022-03-16 14:00:00', '2022-03-16 14:00:00');
