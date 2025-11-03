-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Nov 03, 2025 at 03:38 AM
-- Server version: 8.4.3
-- PHP Version: 8.3.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `evermos_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `alamats`
--

CREATE TABLE `alamats` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `nama_penerima` varchar(100) NOT NULL,
  `no_hp_penerima` varchar(20) NOT NULL,
  `provinsi_id` longtext,
  `kota_id` longtext,
  `kecamatan_id` longtext,
  `kelurahan_id` longtext,
  `detail_alamat` text,
  `kode_pos` longtext,
  `is_default` tinyint(1) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `alamats`
--

INSERT INTO `alamats` (`id`, `user_id`, `nama_penerima`, `no_hp_penerima`, `provinsi_id`, `kota_id`, `kecamatan_id`, `kelurahan_id`, `detail_alamat`, `kode_pos`, `is_default`, `created_at`, `updated_at`) VALUES
(1, 1, 'Sarah Syakira', '08123456789', '31', '3173', '3173080', '3173080005', 'Jl. Melati No. 12, RT 03 RW 04', '13410', 1, '2025-11-03 08:54:51.046', '2025-11-03 08:54:51.046');

-- --------------------------------------------------------

--
-- Table structure for table `detail_trxes`
--

CREATE TABLE `detail_trxes` (
  `id` bigint UNSIGNED NOT NULL,
  `id_trx` bigint UNSIGNED DEFAULT NULL,
  `id_produk` bigint UNSIGNED DEFAULT NULL,
  `id_log_produk` bigint UNSIGNED DEFAULT NULL,
  `id_toko` bigint UNSIGNED DEFAULT NULL,
  `kuantitas` bigint DEFAULT NULL,
  `harga_total` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `foto_produks`
--

CREATE TABLE `foto_produks` (
  `id` bigint UNSIGNED NOT NULL,
  `id_produk` bigint UNSIGNED DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `kategoris`
--

CREATE TABLE `kategoris` (
  `id` bigint UNSIGNED NOT NULL,
  `nama` varchar(100) NOT NULL,
  `deskripsi` text,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `kategoris`
--

INSERT INTO `kategoris` (`id`, `nama`, `deskripsi`, `created_at`, `updated_at`) VALUES
(1, 'Elektronik', 'Kategori barang elektronik', '2025-11-03 09:52:12.937', '2025-11-03 09:52:12.937');

-- --------------------------------------------------------

--
-- Table structure for table `log_produks`
--

CREATE TABLE `log_produks` (
  `id` bigint UNSIGNED NOT NULL,
  `nama_produk` longtext,
  `slug` longtext,
  `harga_reseller` bigint DEFAULT NULL,
  `harga_konsumen` bigint DEFAULT NULL,
  `stok` bigint DEFAULT NULL,
  `id_toko` bigint UNSIGNED DEFAULT NULL,
  `id_kategori` bigint UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `produks`
--

CREATE TABLE `produks` (
  `id` bigint UNSIGNED NOT NULL,
  `nama_produk` varchar(255) NOT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `harga_reseller` bigint DEFAULT NULL,
  `harga_konsumen` bigint DEFAULT NULL,
  `stok` bigint DEFAULT NULL,
  `id_toko` bigint UNSIGNED DEFAULT NULL,
  `id_kategori` bigint UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `produks`
--

INSERT INTO `produks` (`id`, `nama_produk`, `slug`, `harga_reseller`, `harga_konsumen`, `stok`, `id_toko`, `id_kategori`, `created_at`, `updated_at`) VALUES
(1, 'Produk A', 'produk-a', 10000, 12000, 50, 1, 1, '2025-11-03 10:29:21.444', '2025-11-03 10:29:21.444'),
(2, 'TV', 'tv-elektronik', 1000000, 1200000, 50, 1, 1, '2025-11-03 10:31:04.038', '2025-11-03 10:31:04.038');

-- --------------------------------------------------------

--
-- Table structure for table `tokos`
--

CREATE TABLE `tokos` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `nama_toko` varchar(255) NOT NULL,
  `url_toko` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `tokos`
--

INSERT INTO `tokos` (`id`, `user_id`, `nama_toko`, `url_toko`, `created_at`, `updated_at`) VALUES
(1, 1, 'Toko Sarah Syakira Rambe', 'https://toko.com/Sarah Syakira Rambe', '2025-11-03 08:16:54.529', '2025-11-03 08:16:54.529');

-- --------------------------------------------------------

--
-- Table structure for table `trxes`
--

CREATE TABLE `trxes` (
  `id` bigint UNSIGNED NOT NULL,
  `id_user` bigint UNSIGNED DEFAULT NULL,
  `alamat_pengiriman` bigint UNSIGNED DEFAULT NULL,
  `harga_total` bigint DEFAULT NULL,
  `kode_invoice` longtext,
  `method_bayar` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `nama` longtext,
  `password` longtext,
  `no_telp` longtext,
  `tanggal_lahir` longtext,
  `jenis_kelamin` longtext,
  `tentang` longtext,
  `pekerjaan` longtext,
  `email` varchar(191) DEFAULT NULL,
  `id_provinsi` longtext,
  `id_kota` longtext,
  `is_admin` tinyint(1) DEFAULT '0',
  `token` longtext,
  `refresh_token` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `password`, `no_telp`, `tanggal_lahir`, `jenis_kelamin`, `tentang`, `pekerjaan`, `email`, `id_provinsi`, `id_kota`, `is_admin`, `token`, `refresh_token`, `created_at`, `updated_at`) VALUES
(1, 'Sarah Syakira Rambe', '$2a$10$HRCIzI7hf6u8RK5WpB3pIOp/ThooTDbQ43/LJerDL3NSD2tSPZcJ.', '081234567812', '2003-08-03', 'Perempuan', 'Fresh Graduate', 'Belum Pernah Bekerja', 'ssykra03@gmail.com', '31', '3137', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InNzeWtyYTAzQGdtYWlsLmNvbSIsImlzX2FkbWluIjp0cnVlLCJleHAiOjE3NjIyMjY4ODN9.GNcB314UU8qcb0GSNsAU-8gE4lb8iIHiviQlB4eiu2w', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJlbWFpbCI6IiIsImlzX2FkbWluIjpmYWxzZSwiZXhwIjoxNzYyNzQ1MjgzfQ.e0UP9vcsVl8SsLft1Ss_d16-OdfVLHMPl0pXEILsL00', '2025-11-03 08:16:54.474', '2025-11-03 10:28:03.570');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `alamats`
--
ALTER TABLE `alamats`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `detail_trxes`
--
ALTER TABLE `detail_trxes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_trxes_detail_trx` (`id_trx`);

--
-- Indexes for table `foto_produks`
--
ALTER TABLE `foto_produks`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_produks_foto_produk` (`id_produk`);

--
-- Indexes for table `kategoris`
--
ALTER TABLE `kategoris`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_kategoris_nama` (`nama`);

--
-- Indexes for table `log_produks`
--
ALTER TABLE `log_produks`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `produks`
--
ALTER TABLE `produks`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_produks_slug` (`slug`);

--
-- Indexes for table `tokos`
--
ALTER TABLE `tokos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `trxes`
--
ALTER TABLE `trxes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `alamats`
--
ALTER TABLE `alamats`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `detail_trxes`
--
ALTER TABLE `detail_trxes`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `foto_produks`
--
ALTER TABLE `foto_produks`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `kategoris`
--
ALTER TABLE `kategoris`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `log_produks`
--
ALTER TABLE `log_produks`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `produks`
--
ALTER TABLE `produks`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `tokos`
--
ALTER TABLE `tokos`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `trxes`
--
ALTER TABLE `trxes`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `detail_trxes`
--
ALTER TABLE `detail_trxes`
  ADD CONSTRAINT `fk_trxes_detail_trx` FOREIGN KEY (`id_trx`) REFERENCES `trxes` (`id`);

--
-- Constraints for table `foto_produks`
--
ALTER TABLE `foto_produks`
  ADD CONSTRAINT `fk_produks_foto_produk` FOREIGN KEY (`id_produk`) REFERENCES `produks` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
