CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  usuario VARCHAR(50) NOT NULL,
  correo VARCHAR(60) UNIQUE NOT NULL,
  telefono VARCHAR(10) UNIQUE NOT NULL,
  contrasena VARCHAR(50) NOT NULL
);