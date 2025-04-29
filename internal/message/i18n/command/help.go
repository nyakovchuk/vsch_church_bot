package command

func EnCommandHelp() string {
	return `
⛪ The bot supports <b>three ways</b> to share your location for finding nearby churches:

<b>1.</b> <u>"Send Location" button</u> <i>(available on mobile devices only)</i>

<b>2.</b> <u>Via the Telegram menu</u> <i>(works on all devices)</i>
Tap the 📎 <b><i>"Attachments"</i></b> icon in the input field → <b><i>"Location"</i></b> → Choose a point on the map.

<b>3.</b> <u>Manual coordinate input</u>
<b>Send a message in one of the following formats:</b>
<code>50.8096 25.4109</code> <i>(latitude and longitude separated by a space)</i>
<code>50.8096, 25.4109</code> <i>(latitude and longitude separated by a comma)</i>.`
}

func UkCommandHelp() string {
	return `
⛪ Бот підтримує <b>три способи</b> передавання геолокації для пошуку церков:

<b>1.</b> <u>Кнопка "Надіслати місцезнаходження"</u> <i>(лише для мобільних пристроїв)</i>

<b>2.</b> <u>Через меню Telegram</u> <i>(працює на всіх пристроях)</i>
Натисніть на значок 📎 <b><i>"Вкладення"</i></b> у полі введення → <b><i>"Місцезнаходження"</i></b> → Виберіть точку на мапі.

<b>3.</b> <u>Ручне введення координат</u>
<b>Надішліть повідомлення у форматі:</b>
<code>50.8096 25.4109</code> <i>(широта, довгота через пробіл)</i>
<code>50.8096, 25.4109</code> <i>(широта, довгота через кому)</i>.`
}

func RuCommandHelp() string {
	return `
⛪ Бот поддерживает <b>три способа</b> передачи геолокации для поиска церквей:

<b>1.</b> <u>Кнопка "Отправить местоположение"</u> <i>(только для мобильных устройств)</i>

<b>2.</b> <u>Через меню Telegram</u> <i>(работает на всех устройствах)</i>
Нажмите на значок 📎 <b><i>"Вложения"</i></b> в поле ввода → <b><i>"Местоположение"</i></b> → Выберите точку на карте.

<b>3.</b> <u>Ручной ввод координат</u>
<b>Отправьте сообщение в формате:</b>
<code>50.8096 25.4109</code> <i>(широта, долгота через пробел)</i>
<code>50.8096, 25.4109</code> <i>(широта, долгота через запятую)</i>.`
}
