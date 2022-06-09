## Comandos para funcionamento do projeto

#### Cria a base de dados 
CREATE DATABASE IF NOT EXISTS poo_avancado_web

#### Seleciona a base de dados
USE poo_avancado_web

#### Cria a tabela
CREATE TABLE IF NOT EXISTS carros (
	id 				INT AUTO_INCREMENT PRIMARY KEY,
	marca 		VARCHAR(50) NOT NULL,
	nome			VARCHAR(50) NOT NULL,
	placa			VARCHAR(50) NOT NULL,
	preco 		FLOAT NOT NULL,
  vendido   TINYINT NOT NULL
)

#### Inicia a imagem docker do projeto
docker run -p 3306:3306 --name golang_sql -e MYSQL_ROOT_PASSWORD=1234 -d mysql:8.0.29
