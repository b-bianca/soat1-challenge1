-- CREATE DATABASE `restaurante`;
-- USE `restaurante`;

-- cliente
DROP TABLE IF EXISTS cliente;
CREATE TABLE cliente (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	nome VARCHAR(250) NOT NULL,
	email VARCHAR(250) NULL,
	cpf VARCHAR(11) NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- pagamento
DROP TABLE IF EXISTS pagamento_status;
CREATE TABLE pagamento_status (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	nome VARCHAR(250) NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS pagamento;
CREATE TABLE pagamento (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	pedido_id BIGINT NOT NULL,
	valor DECIMAL(8, 2) NOT NULL,
	status_id BIGINT NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- produto
DROP TABLE IF EXISTS produto_categoria;
CREATE TABLE produto_categoria (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	nome VARCHAR(250) NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS produto;
CREATE TABLE produto (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	nome VARCHAR(250) NOT NULL,
	descricao VARCHAR(250) NOT NULL,
	categoria_id BIGINT NOT NULL,
	preco DECIMAL(8, 2) NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	atualizado_em DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- pedido
DROP TABLE IF EXISTS pedido_status;
CREATE TABLE pedido_status (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	nome VARCHAR(250) NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS pedido_item;
CREATE TABLE pedido_item (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	pedido_id BIGINT NOT NULL,
	produto_id BIGINT NOT NULL,
	quantidade INT NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS pedido;
CREATE TABLE pedido (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	cliente_id BIGINT NULL,
	confirmado BOOL NOT NULL,
	pago BOOL NOT NULL,
	status_id BIGINT NOT NULL,
	criado_em DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- foreign keys
ALTER TABLE pagamento ADD CONSTRAINT pagamento_status_fk FOREIGN KEY (status_id) REFERENCES pagamento_status(id);
ALTER TABLE pagamento ADD CONSTRAINT pagamento_pedido_fk FOREIGN KEY (pedido_id) REFERENCES pedido(id);
ALTER TABLE produto ADD CONSTRAINT produto_fk FOREIGN KEY (categoria_id) REFERENCES produto_categoria(id);
ALTER TABLE pedido_item ADD CONSTRAINT item_pedido_fk FOREIGN KEY (pedido_id) REFERENCES pedido(id);
ALTER TABLE pedido_item ADD CONSTRAINT item_produto_fk FOREIGN KEY (produto_id) REFERENCES produto(id);
ALTER TABLE pedido ADD CONSTRAINT pedido_cliente_fk FOREIGN KEY (cliente_id) REFERENCES cliente(id);
ALTER TABLE pedido ADD CONSTRAINT pedido_status_fk FOREIGN KEY (status_id) REFERENCES pedido_status(id);


-- data insertions
INSERT INTO cliente (nome, email, cpf) VALUES (N'João', N'j@gmail.com', '12345678901');
INSERT INTO cliente (nome, email, cpf) VALUES (N'Maria', N'm@gmail,com', '12345678902');
INSERT INTO cliente (nome, email, cpf) VALUES (N'José', N'jj@gmail.com', '12345678903');
INSERT INTO cliente (nome, email, cpf) VALUES (N'Ana', N'a@gmail.com', '12345678904');

INSERT INTO produto_categoria (nome) VALUES (N'Lanche');
INSERT INTO produto_categoria (nome) VALUES (N'Acompanhamento');
INSERT INTO produto_categoria (nome) VALUES (N'Bebida');
INSERT INTO produto_categoria (nome) VALUES (N'Sobremesa');

INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'X-Burguer', N'Pão, carne, queijo, alface, tomate e maionese', 1, 10.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'X-Salada', N'Pão, carne, queijo, alface, tomate, maionese e batata palha', 1, 12.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'X-Bacon', N'Pão, carne, queijo, alface, tomate, maionese e bacon', 1, 15.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'X-Tudo', N'Pão, carne, queijo, alface, tomate, maionese, bacon, ovo e batata palha', 1, 18.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Batata Frita', N'Batata frita com molho especial', 2, 5.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Refrigerante', N'Refrigerante lata 350ml', 3, 4.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Suco', N'Suco natural 500ml', 3, 6.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Milk Shake', N'Milk Shake 500ml', 3, 8.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Sorvete', N'Sorvete 500ml', 4, 10.00);
INSERT INTO produto (nome, descricao, categoria_id, preco) VALUES (N'Pudim', N'Pudim de leite', 4, 8.00);

INSERT INTO pedido_status (nome) VALUES (N'Recebido');
INSERT INTO pedido_status (nome) VALUES (N'Em preparação');
INSERT INTO pedido_status (nome) VALUES (N'Pronto');
INSERT INTO pedido_status (nome) VALUES (N'Finalizado');

INSERT INTO pagamento_status (nome) VALUES (N'Pendente');
INSERT INTO pagamento_status (nome) VALUES (N'Pago');
