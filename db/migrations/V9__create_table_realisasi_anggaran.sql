CREATE TABLE `realisasi_anggaran` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `kode_subkegiatan` varchar(255) NOT NULL,
  `realisasi_anggaran` int NOT NULL,
  `kode_opd` varchar(30) NOT NULL,
  `rencana_aksi` text,
  `faktor_pendorong` text,
  `faktor_penghambat` text,
  `rekomendasi_tl` text,
  `bukti_dukung` text,
  `bulan` varchar(30) NOT NULL,
  `tahun` varchar(30) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);