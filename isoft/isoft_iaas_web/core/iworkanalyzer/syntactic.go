package iworkanalyzer

func AnalysisSyntactic(metas []string, lexers []string) (err error) {
	return GetExecuteOrder(metas, lexers)
}

func GetExecuteOrder(metas, lexers []string) error {
	err := RenderFuncCallers(metas, lexers)
	return err
}
