CREATE TABLE petugas_tim (
  id int PRIMARY KEY AUTO_INCREMENT,
  id_program_unggulan int not null,
  kode_tim varchar(255) not null,
  pegawai_id varchar(255) not null,
  tahun int not null,
  bulan int,
  created_at timestamp DEFAULT (now()),
  updated_at timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);
