package Ganancias

type CalculoSACPrimerCuotaOtrosEmpleos struct {
	CalculoGanancias
}

func (cg *CalculoSACPrimerCuotaOtrosEmpleos) getResultInternal() float64 {
	var importeTotal float64 = 0
	if getfgMes(&cg.Liquidacion.Fechaperiodoliquidacion) <= 6 {
		importeTotal = cg.getfgImporteGananciasOtroEmpleoSiradig("sac")
	}
	return importeTotal
}

func (cg *CalculoSACPrimerCuotaOtrosEmpleos) getResult() float64 {
	return cg.getResultOnDemandTemplate("SAC_PRIMER_CUOTA_OTROS_EMPLEOS", 10, cg)
}

func (cg *CalculoSACPrimerCuotaOtrosEmpleos) getTope() *float64 {
	return nil
}

func (cg *CalculoSACPrimerCuotaOtrosEmpleos) getNombre() string {
	return "SAC Primer Cuota Otros Empleos (+)"
}

func (cg *CalculoSACPrimerCuotaOtrosEmpleos) getEsMostrable() bool {
	return true
}
