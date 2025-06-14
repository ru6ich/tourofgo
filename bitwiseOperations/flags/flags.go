package flags

/*Реализовать тип, описывающий набор прав доступа, и создать функции
для управления флагами с помощью побитовых операций.*/

type Flags uint8

const (
	Read    Flags = 1 << iota //00000001
	Write                     //00000010
	Execute                   //00000100
	Delete                    //00001000
	Share                     //00010000
)

const allFlags = Read | Write | Execute | Delete | Share

//Установка указанного права
func (flag *Flags) Set(flagValue Flags) {
	*flag = *flag | flagValue
}

//Очистить права
func (flag *Flags) Clear() {
	*flag = 0
}

//Инвертировать права
func (flag *Flags) Toggle() {
	*flag = *flag ^ allFlags
}

//Проверить наличие прав
func (flag *Flags) Has(flagValue Flags) bool {
	return *flag&flagValue == flagValue
}
