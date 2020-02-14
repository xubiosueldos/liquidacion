package Ganancias

type CalculoPrimasDeSeguroParaCasoDeMuerte struct {
	CalculoGanancias
}

func (cg *CalculoPrimasDeSeguroParaCasoDeMuerte) getResultInternal() float64 {
	return 0
}

func (cg *CalculoPrimasDeSeguroParaCasoDeMuerte) getResult() float64 {
	return cg.getResultOnDemandTemplate("Primas de seguro para el caso de muerte (-)", "PRIMAS_DE_SEGURO_PARA_CASO_DE_MUERTE", 24, cg)
}

func (cg *CalculoPrimasDeSeguroParaCasoDeMuerte) getTope() *float64 {
	//ESTO TIENE TOPE PERO AUN NO SE IMPLEMENTO TODO
	return nil
}
