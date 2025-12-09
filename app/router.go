package app

import (
	"timkerjaService/controller"

	myMiddleware "timkerjaService/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(timKerjaController controller.TimKerjaController, susunanTimController controller.SusunanTimController, jabatanTimController controller.JabatanTimController, realisasiAnggaranController controller.RealisasiAnggaranController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(myMiddleware.SessionIDMiddleware)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/swagger/doc.json", echoSwagger.WrapHandler)

	e.POST("/timkerja", timKerjaController.Create)
	e.PUT("/timkerja/:id", timKerjaController.Update)
	e.DELETE("/timkerja/:id", timKerjaController.Delete)
	e.GET("/timkerja/:id", timKerjaController.FindById)
	e.GET("/only_timkerja", timKerjaController.FindAll)
	e.GET("/timkerja", timKerjaController.FindAllTm)
	e.GET("/timkerja-non-sekretariat", timKerjaController.FindAllTimNonSekretariat)
	e.GET("/timkerja-sekretariat", timKerjaController.FindAllTimSekretariat)
	// Program unggulan
	e.POST("/timkerja/:kodetim/program_unggulan", timKerjaController.AddProgramUnggulan)
	// TODO: TARGET PERBAIKAN
	e.GET("/timkerja/:kodetim/program_unggulan", timKerjaController.FindAllProgramUnggulanTim)
	e.DELETE("/timkerja/:kodetim/program_unggulan/:id", timKerjaController.DeleteProgramUnggulan)
	// POST simpan realisasi anggaran by subkegiatan
	// post simpan faktor pendrong, penghambat, rtl, bukti dukung
	e.POST("/timkerja/:kodetim/realisasi_pokin", timKerjaController.SaveRealisasiPokin)
	// patch simpan url bukti dukung
	// post simpan rencana kinerja dari sekret
	e.POST("/timkerja_sekretariat/:kodetim/rencana_kinerja", timKerjaController.AddRencanaKinerja)
	// get rekin by tim kerja sekret
	e.GET("/timkerja_sekretariat/:kodetim/rencana_kinerja", timKerjaController.FindAllRencanaKinerjaTim)
	// hapus rencana kinerja dari sekret
	e.DELETE("/timkerja_sekretariat/:kodetim/rencana_kinerja/:id", timKerjaController.DeleteRencanaKinerjaTim)
	// response sret

	e.POST("/susunantim", susunanTimController.Create)
	e.PUT("/susunantim/:id", susunanTimController.Update)
	e.DELETE("/susunantim/:id", susunanTimController.Delete)
	e.GET("/susunantim/:id", susunanTimController.FindById)
	e.GET("/susunantim", susunanTimController.FindAll)

	e.POST("/jabatantim", jabatanTimController.Create)
	e.PUT("/jabatantim/:id", jabatanTimController.Update)
	e.DELETE("/jabatantim/:id", jabatanTimController.Delete)
	e.GET("/jabatantim/:id", jabatanTimController.FindById)
	e.GET("/jabatantim", jabatanTimController.FindAll)

	// Realisasi Anggaran
	e.DELETE("/realisasianggaran/:id", realisasiAnggaranController.Delete)
	e.GET("/realisasianggaran/detail/:id", realisasiAnggaranController.FindById)
	e.GET("/realisasianggaran/:kode_subkegiatan/:kode_tim/:id_rencana_kinerja/:bulan/:tahun", realisasiAnggaranController.FindAll)
	e.POST("/realisasianggaran", realisasiAnggaranController.Upsert)

	// Penilaian Kinerja
	e.GET("/penilaian_kinerja/:kode_tim/:bulan/:tahun", penilaianKinerjaController.FindByKodeTim) // filter by bulan tahun
	e.POST("/penilaian_kinerja", penilaianKinerjaController.Create)
	e.PUT("/penilaian_kinerja/:id", penilaianKinerjaController.Update)
	e.DELETE("/penilaian_kinerja/:id", penilaianKinerjaController.Delete)

	return e
}
