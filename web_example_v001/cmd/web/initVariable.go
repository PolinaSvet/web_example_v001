package cmd

func InitVariable() error {

	var var_ ListTasks
	//task_001
	var_.Image = "/ui/static/img/task_001.jpg"
	var_.Name = "Вариант №1<br/>ЗАДАНИЕ 5.7.1(SC-03)"
	var_.Describe = "Модуль 5: Основы Go<br/>"
	var_.Link = "/z_task_001_01/"
	var_.PageHTML = "ui/html/z_task_001_01.html"
	var_.Task = "Создайте новый проект с именем exercise. " +
		"Инициализируйте модуль, напишите в него код, который объявляет пять нетипизированных целочисленных констант с идентификаторами месяцев на английском языке с января по май так, чтобы эти константы имели значения от 1 до 5 — соответственно порядковым номерам месяцев в году. " +
		"Выведите каждую из них в консоль. Проверьте, что выводятся правильные номера месяцев. Установите программу с помощью команды go install и проверьте, что ваша программа вызывается отовсюду. Рекомендуем проверить доступность программы хотя бы из двух разных папок. " +
		"Попробуйте использовать множественное объявление констант и идентификатор iota, причём только один раз."
	ListTasksAll = append(ListTasksAll, var_) //0

	var_.Image = "/ui/static/img/task_001.jpg"
	var_.Name = "Вариант №2<br/>ЗАДАНИЕ 5.7.1(SC-03)"
	var_.Describe = "Модуль 5: Основы Go<br/>"
	var_.Link = "/z_task_001_02/"
	var_.PageHTML = "ui/html/z_task_001_02.html"
	ListTasksAll = append(ListTasksAll, var_) //1

	var_.Image = "/ui/static/img/task_001.jpg"
	var_.Name = "Вариант №3<br/>ЗАДАНИЕ 5.7.1(SC-03)"
	var_.Describe = "Модуль 5: Основы Go<br/>"
	var_.Link = "/z_task_001_03/"
	var_.PageHTML = "ui/html/z_task_001_03.html"
	ListTasksAll = append(ListTasksAll, var_) //2

	//task_002
	var_.Image = "/ui/static/img/task_002.jpg"
	var_.Name = "Вариант №1<br/>Задание 6.8.1(HW-01): Реализация простого калькулятора."
	var_.Describe = "Модуль 6. Ветвления, циклы, работа с вводом/выводом<br/>"
	var_.Link = "/z_task_002_01/"
	var_.PageHTML = "ui/html/z_task_002_01.html"
	var_.Task = "Напишите программу, которая считывает первое число, затем арифметический оператор (+, -, *, /), " +
		"затем второе число, после чего, в зависимости от арифметического оператора, " +
		"производит нужное действие и выводит строку с результатом."
	ListTasksAll = append(ListTasksAll, var_) //3

	var_.Image = "/ui/static/img/task_002.jpg"
	var_.Name = "Вариант №2<br/>Задание 6.8.1(HW-01): Реализация простого калькулятора."
	var_.Describe = "Модуль 6. Ветвления, циклы, работа с вводом/выводом<br/>"
	var_.Link = "/z_task_002_02/"
	var_.PageHTML = "ui/html/z_task_002_02.html"
	ListTasksAll = append(ListTasksAll, var_) //4

	var_.Image = "/ui/static/img/task_002.jpg"
	var_.Name = "Вариант №3<br/>Задание 6.8.1(HW-01): Реализация простого калькулятора."
	var_.Describe = "Модуль 6. Ветвления, циклы, работа с вводом/выводом<br/>"
	var_.Link = "/z_task_002_03/"
	var_.PageHTML = "ui/html/z_task_002_03.html"
	ListTasksAll = append(ListTasksAll, var_) //5

	//task_003
	var_.Image = "/ui/static/img/task_003.jpg"
	var_.Name = "Задание 8.8.1(HW-01): Создайте пакет electronic и добавьте в него интерфейсы."
	var_.Describe = "Модуль 8. Обработка ошибок и паник, рекурсия, интерфейсы<br/>"
	var_.Link = "/z_task_003_01/"
	var_.PageHTML = "ui/html/z_task_003_01.html"
	var_.Task = "Создайте пакет electronic и добавьте в него интерфейсы."
	ListTasksAll = append(ListTasksAll, var_) //6

	//task_004
	var_.Image = "/ui/static/img/task_004.jpg"
	var_.Name = "Задание 8.8.2(HW-01): Опишите 2 интерфейса: Auto и Dimensions."
	var_.Describe = "Модуль 8. Обработка ошибок и паник, рекурсия, интерфейсы<br/>"
	var_.Link = "/zTask004/"
	var_.PageHTML = "ui/html/zTask004.html"
	var_.Task = "Опишите 2 интерфейса: Auto и Dimensions."
	ListTasksAll = append(ListTasksAll, var_) //7

	//task_005
	var_.Image = "/ui/static/img/task_005.jpg"
	var_.Name = "Задание 9.8.1(HW-01): Создайте структуру Man, представляющую человека."
	var_.Describe = "Модуль 9. Массивы, словари, слайсы, строки, руны и слайсы байт<br/>"
	var_.Link = "/zTask005/"
	var_.PageHTML = "ui/html/zTask005.html"
	var_.Task = "Создайте структуру Man, представляющую человека."
	ListTasksAll = append(ListTasksAll, var_) //8

	//task_006
	var_.Image = "/ui/static/img/task_006.jpg"
	var_.Name = "Задание 12.8.1(HW-03): Сортировка (англ. sorting — классификация, упорядочение) — последовательное расположение или разбиение на группы чего-либо в зависимости от выбранного критерия."
	var_.Describe = "Модуль 12. Алгоритмы сортировки<br/>"
	var_.Link = "/zTask006/"
	var_.PageHTML = "ui/html/zTask006.html"
	var_.Task = "Реализация всех изученных видов сортировок."
	ListTasksAll = append(ListTasksAll, var_) //9

	//task_007
	var_.Image = "/ui/static/img/task_007.jpg"
	var_.Name = "Задание 13.11.1(HW-03): Реализовать структуры двоичного дерева, неориентированного графа, ориентированного графа."
	var_.Describe = "Модуль 13. Деревья, алгоритмы поиска, графы<br/>"
	var_.Link = "/zTask007/"
	var_.PageHTML = "ui/html/zTask007.html"
	var_.Task = "Двоичное дерево, неориентированный граф, ориентированный граф"
	ListTasksAll = append(ListTasksAll, var_) //10

	//task_008
	var_.Image = "/ui/static/img/task_008.jpg"
	var_.Name = "Задания 14.6.1(HW-03): Реализуйте программу, которая будет находить общие элементы в двух массивах и использовать map для этой цели. Значениями массивов являются строки."
	var_.Describe = "Модуль 14. Хеш-мап, Хеш-функции<br/>"
	var_.Link = "/zTask008/"
	var_.PageHTML = "ui/html/zTask008.html"
	var_.Task = "Реализуйте программу, которая будет находить общие элементы в двух массивах и использовать map для этой цели."
	ListTasksAll = append(ListTasksAll, var_) //11

	//task_009
	var_.Image = "/ui/static/img/task_009.jpg"
	var_.Name = "Задание 16.6.2 (HW): Напишите структуру, которая будет реализовывать клиент для клиента банковского приложения."
	var_.Describe = "Модуль 16. Горутины и базовая синхронизация<br/>"
	var_.Link = "/zTask009/"
	var_.PageHTML = "ui/html/zTask009.html"
	var_.Task = "Напишите структуру, которая будет реализовывать клиент для клиента банковского приложения."
	ListTasksAll = append(ListTasksAll, var_) //12

	//task_010
	var_.Image = "/ui/static/img/task_010.jpg"
	var_.Name = "Задания № 17.3.1, 17.3.3, 17.6.1, 17.6.2, 17.6.3, 17.7.1"
	var_.Describe = "Модуль 17. Синхронизация с использованием атомиков и каналов<br/>"
	var_.Link = "/zTask010/"
	var_.PageHTML = "ui/html/zTask010.html"
	var_.Task = "-"
	ListTasksAll = append(ListTasksAll, var_) //13
	return nil
}
