# systemd-watchdog
using systemd watchdog

# english
Using a systemd watchdog in a service.
specify the check time in the .service file
[Service]
WatchdogSec=100

or set the time via the function
watchdog.SetInterval(time.Duration(100 * time.Second))

You need to feed the dog more often than the established time (it is recommended twice as often)
watchdog.Refresh()

when the timer fires, systemd kills the process with the SIGABRT signal

# cyrillic 
использование стророжевого таймера systemd в сервисе.
указать время проверики в файле .service
[Service]
WatchdogSec=100

или установить время через функцию
watchdog.SetInterval(time.Duration(100 * time.Second))

кормить собаку надо чаще чем установлено время (рекомендуют в два раза чаще)
watchdog.Refresh()

при срабатывании таймера systemd убивает процесс с сигналом SIGABRT

