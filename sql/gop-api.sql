-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Waktu pembuatan: 12 Feb 2023 pada 18.28
-- Versi server: 5.7.39
-- Versi PHP: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gop-api`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `user_id` varchar(36) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `date_created` datetime(3) DEFAULT NULL,
  `date_updated` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`user_id`, `name`, `email`, `password`, `date_created`, `date_updated`) VALUES
('a5f7ae2a-aaf9-11ed-8350-60599c074d6b', 'Habib Irfan Mahaasin', 'admin@mahaasin.com', '$2a$10$BHZ3hC8RnNXRwmZOT83E4.enMiY9ARlmOjxd/C0sR2Uu5HNA18QQq', '2023-02-13 00:21:12.638', '2023-02-13 00:21:12.638'),
('b98566f8-aaf9-11ed-8350-60599c074d6b', 'Joko Joki Juko', 'user@gmail.com', '$2a$10$TaUR/st8rHtvW1HA.hd7xezOGLvHGUXx8gloXdSjB2viQ/G8sjE3u', '2023-02-13 00:21:45.445', '2023-02-13 00:30:31.865');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
