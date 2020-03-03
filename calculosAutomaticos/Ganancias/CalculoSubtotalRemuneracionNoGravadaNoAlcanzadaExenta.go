package Ganancias

type CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta struct {
	CalculoGanancias
}

func (cg *CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta) getResultInternal() float64 {
	var arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta []float64
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoRemuneracionNoAlcanzadaExentaSinHorasExtras{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoHorasExtrasRemuneracionExenta{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoMovilidadYViaticosRemuneracionExenta{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoMaterialDidacticoPersonalDocenteRemuneracionExenta{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoRemuneracionNoAlcanzadaExentaSinHorasExtrasOtrosEmpleos{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoHorasExtrasRemuneracionExentaOtrosEmpleos{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoMovilidadYViaticosRemuneracionExentaOtrosEmpleos{cg.CalculoGanancias}).getResult())
	arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta = append(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta, (&CalculoMaterialDidacticoPersonalDocenteRemuneracionExentaOtrosEmpleos{cg.CalculoGanancias}).getResult())

	return Sum(arraySubtotalRemuneracionNoGravadaNoAlcanzadaExenta)
}

func (cg *CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta) getResult() float64 {
	return cg.getResultOnDemandTemplate("SUBTOTAL_REMUNERACION_NO_GRAVADA_NO_ALCANZADA_EXENTA", 0, cg)
}

func (cg *CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta) getTope() *float64 {
	return nil
}

func (cg *CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta) getNombre() string {
	return "Subtotal Remuneración No Gravada, No Alcanzada, Exenta"
}

func (cg *CalculoSubtotalRemuneracionNoGravadaNoAlcanzadaExenta) getEsMostrable() bool {
	return false
}