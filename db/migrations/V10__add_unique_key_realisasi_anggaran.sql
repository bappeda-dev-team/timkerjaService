ALTER TABLE realisasi_anggaran
  ADD UNIQUE KEY uq_subkeg_bulan_tahun (kode_subkegiatan, bulan, tahun);