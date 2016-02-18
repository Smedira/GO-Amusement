import urllib.request


def interpret(a):
    a = str(a)
    print(a)
    a = a.split('"value" : ')
    print(len(a))
    l = []
    for i in range(1, len(a),2):
        l += [[int(a[i].split("\\")[0]),int(a[i+1].split("\\")[0])]]
    return l

origins="41.43206,-81.38992"

destinations = "31.43206,-81.38992"

#a = urllib.request.urlopen("https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + origins + "&destinations=" + destinations+ "&key=AIzaSyAutbb8OTCgYiBZhZCTHa0E9I-8AWdQ21g").read()
#print(a)


a = "39.155975,-77.310878"

b = "39.155654,-77.309140"

origins = a + "|" + b
destinations = b + "|" + a

a = urllib.request.urlopen("https://maps.googleapis.com/maps/api/distancematrix/json?mode=walking&origins=" + origins + "&destinations=" + destinations+ "&key=AIzaSyAutbb8OTCgYiBZhZCTHa0E9I-8AWdQ21g").read()
print(str(a))

print(interpret(a))
