import urllib.request

for i in range(1, 1125):
    try:
        urllib.request.urlretrieve("https://pireon.pro/data/database/monster/" + str(i) + ".png", str(i) + ".png")
    except Exception as e:
        pass
