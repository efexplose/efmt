package efmt

func (p *Printer) Errorfc(color Color, format string, a ...any) error {
	return Errorf(colorizeToPrint(color, strings.Split(fmt.Sprintf(format, a...), "\n")))
}

func (p *Printer) Errorlnc(color Color, a ...any) error {
	return Errorln(colorizeToPrint(color, strings.Split(fmt.Sprint(a...), "\n")))
}