CREATE TABLE realisasi_pokin (
    id int PRIMARY KEY AUTO_INCREMENT,
    id_pokin int NOT NULL,
    kode_tim VARCHAR(255) NOT NULL,
    jenis_pohon VARCHAR(25) NOT NULL,
    jenis_item VARCHAR(25),
    kode_item VARCHAR(255),
    nama_item VARCHAR(255),
    pagu INT,
    realisasi INT,
    faktor_pendorong TEXT,
    faktor_penghambat TEXT,
    rtl TEXT,
    url_bukti_dukung TEXT,
    tahun VARCHAR(30) NOT NULL,
    kode_opd VARCHAR(30) NOT NULL,
    created_at timestamp DEFAULT (now()),
    updated_at timestamp DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP
);
