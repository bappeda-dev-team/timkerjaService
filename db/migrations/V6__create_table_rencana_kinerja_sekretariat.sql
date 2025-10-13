CREATE TABLE rencana_kinerja_sekretariat (
       id int PRIMARY KEY AUTO_INCREMENT,
       kode_tim VARCHAR(255) NOT NULL,
       id_rencana_kinerja VARCHAR(255) NOT NULL,
       tahun VARCHAR(30) NOT NULL,
       kode_opd VARCHAR(30) NOT NULL
);
