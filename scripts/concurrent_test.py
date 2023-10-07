import requests
import threading

token = 'put your token here'

url = 'http://localhost:5930/rule/create'
data = {'Datasource': 'tem_1{1, Portable, temperature};tem_2{1, Fixed, temperature}', 'Condition': '(tem_2 > 25) & (tem_1 > (tem_2 +3))', 'Action': 'WebSocket: 1,rule Matched, temperature is $tem_2 and $tem_1!', 'CompanyId': '1'}
headers = {'Authorization': 'Bearer ' + token}

concurrent_num = 5000
current_rule = 1

thread_num = 20

def createAndStartRule():
    for i in range(concurrent_num):
        r = requests.post(url, data, headers=headers)
        result = r.json()
        ids = result['data']['ruleId']
        print("rule " + str(ids) + " created")
        urls = 'http://localhost:5930/rule/start?RuleId=' + str(ids) + '&ExecInternal=%40every 5s'
        requests.get(urls, headers=headers)
        print("rule " + str(ids) + " started")


for i in range(thread_num):
    t = threading.Thread(target=createAndStartRule())
    t.start()
