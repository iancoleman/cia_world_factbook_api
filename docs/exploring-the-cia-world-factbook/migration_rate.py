import datetime
import json
import os

for r, ignore, fs in os.walk("/home/ian/git/cia_world_factbook_api/weekly"):
    fs.sort()
    for f in fs:
        d = f.split("_")[0]
        dt = datetime.datetime.strptime(d, "%Y-%m-%d")
        o = open(os.path.join(r, f))
        content = o.read()
        o.close()
        j = json.loads(content)
        value = j["countries"]["italy"]["data"]["people"]["net_migration_rate"]["migrants_per_1000_population"]
        print("%s,%s" % (d, value))
