-- SELECT * FROM public."TEST"

-- CREATE TABLE IF NOT EXISTS transactions (
--     id SERIAL PRIMARY KEY,
--     description TEXT NOT NULL,
--     amount REAL NOT NULL
-- );

-- INSERT INTO transactions (description, amount) VALUES ('Compra de flores', -19.99);
-- INSERT INTO transactions (description, amount) VALUES ('Venta de producto A', 120.50);
-- INSERT INTO transactions (description, amount) VALUES ('Pago de servicio de internet', -45.00);
-- INSERT INTO transactions (description, amount) VALUES ('Depósito bancario', 300.00);
-- INSERT INTO transactions (description, amount) VALUES ('Cena en restaurante', -75.25);
-- INSERT INTO transactions (description, amount) VALUES ('Compra de libros', -39.90);
-- INSERT INTO transactions (description, amount) VALUES ('Reembolso por devolución', 60.00);
-- INSERT INTO transactions (description, amount) VALUES ('Venta de curso en línea', 199.99);
-- INSERT INTO transactions (description, amount) VALUES ('Suscripción mensual', -10.99);
-- INSERT INTO transactions (description, amount) VALUES ('Bonificación por rendimiento', 500.00);

--SELECT * FROM transactions;

DELETE FROM transactions WHERE id = 1;
