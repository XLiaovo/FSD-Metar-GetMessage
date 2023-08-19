import os
import requests
import time

def fetch_and_save_metar():
    timestamp = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    print(f"{timestamp}: 开始获取气象信息，请等待...")

    if os.path.exists("metar.txt"):
        os.remove("metar.txt")

    url = "https://metar.vatsim.net/all"
    response = requests.get(url)
    web_content = response.text

    with open("metar.txt", "w", encoding="utf-8") as file:
        file.write(web_content)

    timestamp = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
    print(f"{timestamp}: 气象信息获取完成，已保存至metar文本中")

try:
    while True:
        fetch_and_save_metar()
        timestamp = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        print(f"{timestamp}: 程序静默，等待下一次执行...")
        time.sleep(1800) # 30分钟执行一次
except KeyboardInterrupt:
    print("Program terminated")

input("Press Enter to exit")
