from huawei_lte_api.Client import Client
from huawei_lte_api.Connection import Connection


modem = Client(Connection('http://192.168.8.1/')).net
try:
    params = ["3FFFFFFF", "03", ("40", "4")]
    params.insert(0, params[2][modem.net_mode()["LTEBand"] == "40"])
    modem.set_net_mode(*params[:3])
    print(f"Set LTE band to {params[0]}")
except Exception as e:
    print(f"LTE band value cannot be set: {e}")
