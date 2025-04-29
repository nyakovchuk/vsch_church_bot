package event

func EnInvalidCoordinatesFormat() string {
	return `
<b>🚫 Invalid coordinates format</b>
Example of the correct format:
<code>50.4228 30.3145</code><i> (latitude, longitude separated by space)</i>
<code>50.4228, 30.3145</code><i> (latitude, longitude separated by comma)</i>.
`
}

func UkInvalidCoordinatesFormat() string {
	return `
<b>🚫 Некоректний формат координат</b>
Приклад правильного формату координат:
<code>50.4228 30.3145</code><i> (широта, довгота через пробіл)</i>
<code>50.4228, 30.3145</code><i> (широта, довгота через кому)</i>.
`
}

func RuInvalidCoordinatesFormat() string {
	return `
<b>🚫 Некорректный формат координат</b>
Пример правильного формата координат:
<code>50.4228 30.3145</code><i> (широта, долгота через пробел)</i>
<code>50.4228, 30.3145</code><i> (широта, долгота через запятую)</i>.
`
}

func EnInvalidLatitude() string {
	return `
<b>🚫 Invalid latitude</b> (must be between -90 and 90), but you entered <b><i>%s</i></b>
`
}

func UkInvalidLatitude() string {
	return `
<b>🚫 Некоректна широта</b> (має бути в діапазоні від -90 до 90), а у вас <b><i>%s</i></b>
`
}

func RuInvalidLatitude() string {
	return `
<b>🚫 Некорректная широта</b> (должна быть в диапазоне -90...90), а у вас <b><i>%s</i></b>
`
}

func EnInvalidLongitude() string {
	return `
<b>🚫 Invalid longitude</b> (must be between -180 and 180), but you entered <b><i>%s</i></b>
`
}

func UkInvalidLongitude() string {
	return `
<b>🚫 Некоректна довгота</b> (має бути в діапазоні від -180 до 180), а у вас <b><i>%s</i></b>
`
}

func RuInvalidLongitude() string {
	return `
<b>🚫 Некорректная долгота</b> (должна быть в диапазоне -180...180), а у вас <b><i>%s</i></b>
`
}
