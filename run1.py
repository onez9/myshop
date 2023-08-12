#!/usr/bin/python3
import re
import socket
import subprocess
import threading
import os

#print(os.getcwd())
#print(os.path.abspath(__file__))
#print(os.path.realpath(__file__))
#print(os.path.dirname(__file__))

clear = lambda: os.system('clear')

status = os.system('systemctl is-active --quiet postgresql')
if status:
    print('postgresql.service не запущен - запустите!')
    os.system('systemctl start postgresql.service')
else:
    print('Начинаю запуск')
    os.chdir(os.path.dirname(__file__))
    try:
        with open('index.html', 'r') as fi:
            text=fi.read()
            #s=re.search(r'\d+.\d+.\d+.\d+', text)
            result=re.findall(r'\d+.\d+.\d+.\d+', text)
            hostname=socket.gethostname()
            IPAddr=socket.gethostbyname(hostname)
            #print(s.group(0))
            with open('index.html', 'w') as fo:
                print("Вывод прошлых ip-адресов:")
                for item in result:
                    print(item)
                    text = text.replace(item, IPAddr)
                    #print(text)
                fo.write(text)
            print(f'Перейдите по адресу: http://{IPAddr}:1323')

        def f1():
            subprocess.call(['go', 'run', 'server.go'])
        def f2():
            subprocess.call(['npm', 'run', 'dev'])

        t1=threading.Thread(target=f1)
        t2=threading.Thread(target=f2)

        t1.start()
        t2.start()
        clear()
        
    except KeyboardInterrupt:
        print('Завершаю работу программы')



