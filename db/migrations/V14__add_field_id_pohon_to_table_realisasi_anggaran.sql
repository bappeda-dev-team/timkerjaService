ALTER TABLE realisasi_anggaran
ADD COLUMN id_pohon int AFTER id_rencana_kinerja,
ADD COLUMN risiko_hukum VARCHAR(255) AFTER faktor_penghambat;
