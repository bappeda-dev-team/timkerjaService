ALTER TABLE realisasi_anggaran 
ADD COLUMN id_program_unggulan int AFTER id_rencana_kinerja,
MODIFY COLUMN bulan int;