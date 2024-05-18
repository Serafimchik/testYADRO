инструкция по запуску:
docker build -t task.exe .
docker run task.exe test1.txt
docker run task.exe test2.txt
...
docker run task.exe test<N>.txt

пример запуска на ОС Windows
PS C:\testYADRO> go build
PS C:\testYADRO> .\testYADRO.exe test1.txt
PS C:\testYADRO> .\testYADRO.exe test2.txt
...
PS C:\testYADRO> .\testYADRO.exe test<N>.txt

Программа работает исправно.
Программа требует рефакторинга.
В дальнейшем можно выделить отдельные действия в самостоятельные функции. 
Например, проверка наличия места в очереди, постановки в очередь, удаление из очереди, проверка наличия свободного стола, занятие стола, освобождение стола.
