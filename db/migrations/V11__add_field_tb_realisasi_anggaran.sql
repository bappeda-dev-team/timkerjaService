ALTER TABLE realisasi_anggaran 
ADD COLUMN kode_tim varchar(255) AFTER id,
ADD COLUMN id_rencana_kinerja varchar(255) AFTER kode_tim;