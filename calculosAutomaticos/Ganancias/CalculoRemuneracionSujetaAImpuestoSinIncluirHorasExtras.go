package Ganancias

import "github.com/xubiosueldos/conexionBD/Liquidacion/structLiquidacion"

type CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras struct {
	CalculoGanancias
}

func (cg *CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras) getResultInternal() float64 {
	importeRemuneracionSujetaAImpuestos := (&CalculoRemuneracionSujetaAImpuesto{cg.CalculoGanancias}).getResult()

	cuotaSindical := (&CalculoCuotaSindical{cg.CalculoGanancias}).getResult()
	obraSocial := (&CalculoAportesObraSocial{cg.CalculoGanancias}).getResult()
	aportesJubilatorios := (&CalculoAportesJubilatoriosRetirosPensionesOSubsidios{cg.CalculoGanancias}).getResult()
	remunerativosMenosDescuentos := importeRemuneracionSujetaAImpuestos
	cuotaSindicalOtros := (&CalculoCuotaSindicalOtrosEmpleos{cg.CalculoGanancias}).getResult()
	obraSocialOtros := (&CalculoAportesObraSocialOtrosEmpleos{cg.CalculoGanancias}).getResult()
	aportesJubilatoriosOtros := (&CalculoAportesJubilatoriosRetirosPensionesOSubsidiosOtrosEmpleos{cg.CalculoGanancias}).getResult()
	remunerativosOtros := cg.obtenerRemunerativosOtros()

	var uno float64 = 1
	var porcentaje float64 = 0
	var porcentajeOtrosEmp float64 = 0

	if remunerativosMenosDescuentos > 0 {
		porcentaje = uno - (cuotaSindical/remunerativosMenosDescuentos + obraSocial/remunerativosMenosDescuentos + aportesJubilatorios/remunerativosMenosDescuentos)
	}
	if remunerativosOtros != 0 {
		porcentajeOtrosEmp = uno - (cuotaSindicalOtros/remunerativosOtros + obraSocialOtros/remunerativosOtros + aportesJubilatoriosOtros/remunerativosOtros)
	}

	var importeTotalHorasExtrasGravadas float64 = 0
	var importeTotalHorasExtrasGravadasOtrosEmpleos float64 = 0

	liquidaciones := *cg.obtenerLiquidacionesIgualAnioLegajoMenorMes()

	importeTotalHorasExtrasGravadas = importeTotalHorasExtrasGravadas + (&CalculoHorasExtrasGravadas{cg.CalculoGanancias}).getResult()
	importeTotalHorasExtrasGravadasOtrosEmpleos = importeTotalHorasExtrasGravadasOtrosEmpleos + (&CalculoHorasExtrasGravadasOtrosEmpleos{cg.CalculoGanancias}).getResult()

	for i := 0; i < len(liquidaciones); i++ {
		itemGanancias := obtenerItemGananciaFromLiquidacion(&liquidaciones[i])
		if itemGanancias == nil {
			itemGanancias = &structLiquidacion.Liquidacionitem{}
		}
		calculoGananciasAnterior := CalculoGanancias{itemGanancias, &liquidaciones[i], cg.Db, true}
		importeTotalHorasExtrasGravadas = importeTotalHorasExtrasGravadas + (&CalculoHorasExtrasGravadas{calculoGananciasAnterior}).getResult()
		importeTotalHorasExtrasGravadasOtrosEmpleos = importeTotalHorasExtrasGravadasOtrosEmpleos + (&CalculoHorasExtrasGravadasOtrosEmpleos{calculoGananciasAnterior}).getResult()
	}

	importeTotal := importeRemuneracionSujetaAImpuestos - (importeTotalHorasExtrasGravadas*porcentaje + importeTotalHorasExtrasGravadasOtrosEmpleos*porcentajeOtrosEmp)
	return importeTotal
}

func (cg *CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras) getResult() float64 {
	return cg.getResultOnDemandTemplate("REMUNERACION_SUJETA_A_IMPUESTO_SIN_INCLUIR_HORAS_EXTRAS", 0, cg)
}

func (cg *CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras) getTope() *float64 {
	return nil
}

func (cg *CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras) getNombre() string {
	return "Remuneracion sujeta a impuesto (sin incluir horas extras)"
}

func (cg *CalculoRemuneracionSujetaAImpuestoSinIncluirHorasExtras) getEsMostrable() bool {
	return false
}
