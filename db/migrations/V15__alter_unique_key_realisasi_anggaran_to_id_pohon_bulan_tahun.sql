ALTER TABLE realisasi_anggaran
  DROP INDEX uq_subkeg_bulan_tahun,
  ADD UNIQUE KEY uq_pohon_bulan_tahun (id_pohon, bulan, tahun);;
