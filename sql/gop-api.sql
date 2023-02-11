-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Waktu pembuatan: 11 Feb 2023 pada 18.07
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
  `user_id` int(20) NOT NULL,
  `user_uuid` varchar(255) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `date_created` datetime(3) DEFAULT NULL,
  `date_updated` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`user_id`, `user_uuid`, `name`, `email`, `password`, `date_created`, `date_updated`) VALUES
(1, '56e1c9c6-a9e9-11ed-8350-60599c074d6b', 'Habib Irfan Mahaasin', 'admin@mahaasin.com', '$2a$10$Qyi2SwAX6HyqPe/RVTJeUOe0x0CjrIxBtBrs0k5SOxeKMa8KnzxZG', '2023-02-11 15:35:22.000', '2023-02-12 01:06:46.669'),
(5, 'a38f5510-7dcf-412c-bfcc-6b47c6cb1bc8', 'Super Admin', 'superadmin@gmail.com', '$2a$10$pgjAffXeyEhNwgZK004keOctCHbaTJOAKgTUL4Kt2lNhQbvL6wxHu', '2023-02-11 19:57:34.602', '2023-02-11 19:57:34.602');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
