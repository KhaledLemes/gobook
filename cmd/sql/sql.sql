CREATE DATABASE IF NOT EXISTS gobook;
USE gobook;

TRUNCATE TABLE IF EXISTS reservas;
TRUNCATE TABLE IF EXISTS quartos;
TRUNCATE TABLE IF EXISTS propriedades;
TRUNCATE TABLE IF EXISTS usuarios;

DROP TABLE IF EXISTS reservas;
DROP TABLE IF EXISTS quartos;
DROP TABLE IF EXISTS propriedades;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(20) not null,
    nome_do_meio varchar(30),
    ultimo_nome varchar(30) not null,
    tel varchar(19) not null,
    nascimento timestamp not null,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    role varchar(10) not null,
    numero_reservas int,
    cadastrado timestamp DEFAULT current_timestamp()
    credito int DEFAULT 0
) ENGINE=INNODB

CREATE TABLE propriedades(
    id int auto_increment primary key,
    nome varchar(200) unique not null,
    descricao varchar(1000) not null,
    estado varchar(50) not null,
    cidade varchar(50),
    petfriendly bit not null default 0,
    categoria varchar(15),
    dono_id int not null,
    foreign key (dono) references usuarios(id),
    endereco varchar(264) not null,
    numero varchar(5)
) ENGINE=INNODB


CREATE TABLE quartos(
    id int auto_increment primary key,
    nome varchar(50) not null,
    valor_noite int not null,
    disponivel bit not null default 0,
    propriedade_id int not null,
    foreign key (propriedade_id) references propriedades(id),
    qt_disponivel int not null
) ENGINE=INNODB



CREATE TABLE reservas(
    id int auto_increment primary key,
    quarto_id int not null,
    foreign key (quarto_id) REFERENCES quartos(id),
    comodidades varchar(1024) not null,
    reembolsavel bit not null default 0,
    valor_por_noite double precision not null,
    valor_total double precision not null,
    taxas double precision,
    noites smallint,
    marcado timestamp DEFAULT current_timestamp(),
    data_checkin timestamp not null,
    data_checkout timestamp not null
) ENGINE=INNODB

