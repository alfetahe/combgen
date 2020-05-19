package combgen

func buildCombinationStruct(clientInput []string, maxAllowedChars int, charRepeat bool) combinations {
	combinationStruct := combinations{
		clientInput: clientInput,
		maxAllowedCharacters: maxAllowedChars,
	}
	combinationStruct.setMaxAllowedCharSize(charRepeat)

	return combinationStruct
}