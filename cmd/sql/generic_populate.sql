TRUNCATE TABLE reserva;
TRUNCATE TABLE quartos;
TRUNCATE TABLE propriedades;
TRUNCATE TABLE usuarios;


INSERT INTO usuarios (nome, nome_do_meio, ultimo_nome, nascimento, email, senha, role, numero_reservas) VALUES
('usuario', 'do', 'mal', '1990-05-15 08:30:00', 'joao.silva@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'forbidden', 0),
('Maria', 'Eduarda', 'Santos', '1985-10-22 14:15:00', 'maria.santos@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 2),
('Pedro', NULL, 'Oliveira', '1992-03-10 09:00:00', 'pedro.oliveira@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 1),
('Ana', 'Luiza', 'Costa', '1998-11-05 18:45:00', 'ana.costa@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 0),
('Lucas', 'Rafael', 'Pereira', '1988-07-19 20:20:00', 'lucas.pereira@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 5),
('Juliana', NULL, 'Almeida', '1995-01-30 11:10:00', 'juliana.almeida@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 1),
('Marcos', 'Antonio', 'Rodrigues', '1982-12-12 16:05:00', 'marcos.rodrigues@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 3),
('Camila', 'Fernandes', 'Gomes', '1991-08-25 07:50:00', 'camila.gomes@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 0),
('Felipe', NULL, 'Martins', '1993-04-14 13:40:00', 'felipe.martins@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'guest', 2),
('Khaled', 'Borges', 'Lemes', '1997-09-08 22:30:00', 'beatriz.carvalho@email.com', 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', 'owner', 4);

INSERT INTO propriedades(nome, descricao, estado, cidade, petfriendly, categoria, dono) VALUES
('Shopping penha', 'Melhor lugar do mundo', 'São Paulo', 'São Paulo',1, 'hotel', 20),
('Casa do Khaled', 'Só pode momo', 'São Paulo', 'São Paulo',1, 'casa', 19);

INSERT INTO quartos(nome, valor_noite, disponivel, propriedade_id, qt_disponivel) VALUES
('mcdonalds', 320, 1, 1, 67),
('americanas', 3, 1, 1, 3),
('loja', 999, 1, 1, 9),
('casa do feioso', 9850, 1, 2, 9);


INSERT INTO reservas(quarto_id, comodidades, valor_por_noite, valor_total, taxas, noites, data_checkin, data_checkout) VALUES
(5, 'Aqui no mcdonalds vc pode comer mt bemn com nosso cafe da manha e idoso da penha aproveite', 30.55, 35.55, 5, 1, '2026-12-31 22:30:00', '2027-01-01 22:30:00'),
(6, 'Lojas americanas todo mudno vai', 2, 20, 0, 10, '2026-12-31 22:30:00', '2027-01-20 22:30:00');



