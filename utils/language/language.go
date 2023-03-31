package language

import (
	languageMaker "golang.org/x/text/language"
)

type SupportedLanguageType int

const (
	EnUS SupportedLanguageType = iota
	ZhCN
)

var SupportedLanguages = []languageMaker.Tag{
	languageMaker.AmericanEnglish,
	languageMaker.Chinese,
}

type LocaleLanguage struct {
	Language languageMaker.Tag
}

func DefaultLocaleLanguage() *LocaleLanguage {
	return &LocaleLanguage{SupportedLanguages[1]}
}

type MessageMap map[string]string

func (m MessageMap) SetMsg(languageType SupportedLanguageType, msg string) MessageMap {
	m[SupportedLanguages[languageType].String()] = msg
	return m
}

func NewMsgMap(msg string) MessageMap {
	return MessageMap{
		SupportedLanguages[0].String(): msg,
	}
}
