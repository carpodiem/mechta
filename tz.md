## Задание

Написать скрипт для вычисления ежемесячных оплат жителей многоквартирного дома.
Константы:
* количество квартира = 6
* стоимость кубометра гада = 6,87689

Скрипт должен выполнить следующее:
1. Запрашиваем текущее показание счетчика у пользователя.
2. Вычисляем потребление за месяц как разницу между текущим и предыдущим показаниями, где предыдущее показание хранится в базе данных или в файле.
3. Вычисляем стоимость потребленного газа как произведение потребления на стоимость кубометра газа.
4. Делим стоимость потребленного газа на количество квартир (6) и округляем до двух знаков после запятой, чтобы получить среднюю стоимость для каждой квартиры.
5. Для первой квартиры записываем среднюю стоимость
6. Для остальных квартир округляем среднее значение до запятой и вычисляем разницу между этим значением и числом в интервале от -2 до 2. Необходимо предусмотреть ежемесячную ротацию значения числа из интервала для каждой квартиры. То есть на протяжении 5 месяце вычитаемое число для одной квартиры будет варьироваться от -2 до 2.
7. Сохраняем данные о текущем показании счетчика и оплате каждой квартиры в базу данных или файл.
8. Выводим результат для каждой квартиры в формате "Название квартиры - сумма к оплате".