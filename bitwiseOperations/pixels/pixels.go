package pixels

/*Реализовать функции, которые позволяют получить отдельные цветовые компоненты
(Alpha, Red, Green, Blue) из 32-битного значения пикселя, закодированного в формате 0xAARRGGBB*/
type Pixel uint32

// Маски для извлечения цветовых байтов
const (
	BlueMask Pixel = 255 << (iota * 8)
	GreenMask
	RedMask
	AlphaMask
)

// Ипользуя маску Alpha получаем байт прозрачности и сдвигаем младшие 3 байта вправо
func Alpha(pixel Pixel) uint8 {
	return uint8(pixel & AlphaMask >> 24)
}

//Ипользуя маску Red получаем байт красного и сдвигаем младшие 2 байта вправо
func Red(pixel Pixel) uint8 {
	return uint8(pixel & RedMask >> 16)
}

//Ипользуя маску Green получаем байт зеленого и сдвигаем младший 1 байт вправо
func Green(pixel Pixel) uint8 {
	return uint8(pixel & GreenMask >> 8)
}

//Используя маску Blue получаем байт синего (он уже занимает младшие 8 бит => ничего не сдвигаем)
func Blue(pixel Pixel) uint8 {
	return uint8(pixel & BlueMask)
}
