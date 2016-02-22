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

q = open("points2.txt","r")

t = q.read()

q.close()

t = t.split("\n")

origins = ""
destinations = ""

for i in range(len(t)):
    t[i] = t[i][1:]
    print(t[i])
    t[i] = t[i].split("\t: ")
    t[i][0] = int(t[i][0])
    t[i][1] = t[i][1].split(" ")[0] + t[i][1].split(" ")[1]
    origins += t[i][1][:-1] + "|"
    destinations += t[i][1][:-1] + "|"

origins = origins[:-1]
print(origins)

destinations = destinations[:-1]

#a = urllib.request.urlopen("https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + origins + "&destinations=" + destinations+ "&key=AIzaSyAutbb8OTCgYiBZhZCTHa0E9I-8AWdQ21g").read()
#print(a)


##a = "39.155975,-77.310878"
##
##b = "39.155654,-77.309140"
##
##origins = a + "|" + b
##destinations = b + "|" + a

a = urllib.request.urlopen("https://maps.googleapis.com/maps/api/distancematrix/json?mode=walking&origins=" + origins + "&destinations=" + destinations+ "&key=AIzaSyAutbb8OTCgYiBZhZCTHa0E9I-8AWdQ21g").read()
print(str(a))

print(len(t))
print(len(destinations))

o = interpret(a)
l = len(t)
q = open("out.txt", "w")
q.write("start,end,distance,duration \n")
for i in range(l):
    for j in range(l):
        v = i * l + j
        c = o[v]
        print(c,i,j,len(t))
        q.write(str(t[i][0]) + "," + str(t[j][0]) + "," + str(c[0]) + "," + str(c[1]) + "\n")

q.close()


