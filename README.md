# Simple-Authentication

Sebelum menggunakan source code ini, terlebih dahulu buat sebuah database baru dengan nama "go_db" (tanpa tanda kutip)

Kemudian buat sebuah table dengan nama "users" (tanpa tanda kutip) menggunakan query berikut :

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL,
  PRIMARY KEY (id)
);