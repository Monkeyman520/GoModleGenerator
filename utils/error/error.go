package error

import (
	"fmt"

	"github.com/Monkeyman520/GoModleGenerator/utils/language"
)

type InternalError struct {
	Msg    string
	MsgMap language.MessageMap
}

func (ie *InternalError) Error() string {
	return fmt.Sprintf("error: %s", ie.Msg)
}

func (ie *InternalError) Message(languageTag string) string {
	if ie.MsgMap == nil {
		return ie.Msg
	}

	if msg, ok := ie.MsgMap[languageTag]; ok {
		return msg
	}
	return ie.Msg
}

func (ie *InternalError) SetMsg(lang language.SupportedLanguageType, msg string) *InternalError {
	if ie.MsgMap == nil {
		ie.MsgMap = language.MessageMap{}
	}
	ie.MsgMap[language.SupportedLanguages[lang].String()] = msg
	return ie
}

func NewPreDefineInternalError(lang language.SupportedLanguageType, defaultMsg string, msg string) *InternalError {
	return (&InternalError{Msg: defaultMsg}).SetMsg(lang, msg)
}

func NewPreDefineZhCNInternalError(defaultMsg string, msg string) *InternalError {
	return NewPreDefineInternalError(language.ZhCN, defaultMsg, msg)
}
