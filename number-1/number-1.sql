-- a. Membuat Tabel-tabel Database
-- Tabel untuk Jenis Kain
CREATE TABLE fabric_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) NOT NULL
);

-- Tabel untuk Nama Kain
CREATE TABLE fabrics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Tabel untuk Kualitas
CREATE TABLE qualities (
    id SERIAL PRIMARY KEY,
    level INT NOT NULL,
    name VARCHAR(50) NOT NULL
);

-- Tabel utama untuk Produk Kain
CREATE TABLE proiducts (
    id SERIAL PRIMARY KEY,
    type_id INT,
    fabric_id INT,
    quality_id INT,
    price BIGINT NOT NULL,
    FOREIGN KEY (type_id) REFERENCES fabric_types(id),
    FOREIGN KEY (fabric_id) REFERENCES fabrics(id),
    FOREIGN KEY (quality_id) REFERENCES qualities(id)
);

-- b. Script Insert Data
-- Insert data ke tabel jenis_kain
INSERT INTO fabric_types (name) VALUES 
('STB'),
('NTB')
ON CONFLICT DO NOTHING;

-- Insert data ke tabel nama_kain
INSERT INTO fabrics (name) VALUES 
('Sutra'),
('Katun')
ON CONFLICT DO NOTHING;

-- Insert data ke tabel kualitas
INSERT INTO qualities (level, name) VALUES 
(1, 'Sangat Bagus'),
(2, 'Bagus'),
(3, 'Cukup Bagus')
ON CONFLICT DO NOTHING;

-- Insert data ke tabel produk_kain
INSERT INTO products (type_id, fabric_id, quality_id, price) VALUES 
(1, 1, 1, 10000000),
(1, 1, 2, 9000000), 
(1, 1, 3, 8000000), 
(2, 2, 1, 10000000),
(2, 2, 2, 10000000),
(2, 2, 3, 10000000) 
ON CONFLICT DO NOTHING;

-- c. Script Query untuk menampilkan data
SELECT 
    ft.name AS 'Jenis Kain',
    f.name AS 'Nama Kain',
    q.level AS 'Kualitas',
    q.name AS 'Nama Kualitas',
    CONCAT('Rp ', FORMAT(pk.harga, 0)) AS 'Harga'
FROM products p
JOIN fabric_types ft ON p.type_id = ft.id
JOIN fabrics f ON p.fabric_id = f.id
JOIN qualities q ON p.quality_id = q.id
ORDER BY ft.name, q.level;