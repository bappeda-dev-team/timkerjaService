CREATE TABLE penilaian_kinerja (
  id int PRIMARY KEY AUTO_INCREMENT,
  id_pegawai VARCHAR(255) NOT NULL,
  kode_tim VARCHAR(255) NOT NULL,
  jenis_nilai VARCHAR(255) NOT NULL,
  nilai_kinerja int NOT NULL,
  tahun VARCHAR(30) NOT NULL,
  bulan int NOT NULL,
  kode_opd VARCHAR(30) NOT NULL,
  created_at timestamp DEFAULT (now()),
  updated_at timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  created_by VARCHAR(255)
);
