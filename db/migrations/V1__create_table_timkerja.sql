CREATE TABLE `tim_kerja` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `kode_tim` varchar(255) UNIQUE NOT NULL,
  `nama_tim` varchar(255) NOT NULL,
  `keterangan` varchar(255),
  `tahun` varchar(30) NOT NULL,
  `is_active` boolean NOT NULL DEFAULT true,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `susunan_tim` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `kode_tim` varchar(255) NOT NULL,
  `pegawai_id` varchar(25) NOT NULL COMMENT 'langsung ambil dari service',
  `nama_jabatan_tim` varchar(255) NOT NULL,
  `is_active` boolean NOT NULL DEFAULT true,
  `keterangan` varchar(255),
  `created_at` timestamp DEFAULT (now()),
 `updated_at` timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `jabatan_tim` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `nama_jabatan` varchar(255) NOT NULL,
  `level_jabatan` int NOT NULL COMMENT '1 -> penanggung jawab, 2 -> koordinator, 3 -> ketua',
  `created_at` timestamp DEFAULT (now()),
 `updated_at` timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `susunan_tim` ADD FOREIGN KEY (`kode_tim`) REFERENCES `tim_kerja` (`kode_tim`)
ON UPDATE CASCADE ON DELETE CASCADE;
