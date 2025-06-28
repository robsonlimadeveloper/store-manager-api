-- Criação da tabela de estabelecimentos
CREATE TABLE IF NOT EXISTS establishments (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    corporate_name VARCHAR(150),
    address VARCHAR(200),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(20),
    street_number VARCHAR(20)
);

-- Criação da tabela de lojas
CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    corporate_name VARCHAR(150),
    address VARCHAR(200),
    city VARCHAR(100),
    state VARCHAR(2),
    zip_code VARCHAR(20),
    street_number VARCHAR(20),
    establishment_id INTEGER NOT NULL,
    CONSTRAINT fk_establishment
        FOREIGN KEY(establishment_id)
        REFERENCES establishments(id)
        ON DELETE RESTRICT
);

-- Dados de exemplo: Estabelecimentos
INSERT INTO establishments (number, name, corporate_name, address, city, state, zip_code, street_number)
VALUES
('EST-001', 'Matriz São Paulo', 'Empresa Matriz LTDA', 'Rua A', 'São Paulo', 'SP', '01000-000', '100'),
('EST-002', 'Filial Rio', 'Empresa Filial RJ LTDA', 'Rua B', 'Rio de Janeiro', 'RJ', '20000-000', '200');

-- Dados de exemplo: Lojas
INSERT INTO stores (number, name, corporate_name, address, city, state, zip_code, street_number, establishment_id)
VALUES
('LOJ-001', 'Loja Centro SP', 'Loja Centro SP LTDA', 'Rua C', 'São Paulo', 'SP', '01100-000', '300', 1),
('LOJ-002', 'Loja Copacabana', 'Loja Copacabana LTDA', 'Rua D', 'Rio de Janeiro', 'RJ', '22000-000', '400', 2);
