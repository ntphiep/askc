from diagrams import Diagram
from diagrams.azure.compute import AppServices
from diagrams.azure.network import LoadBalancers
from diagrams.azure.storage import StorageAccounts

with Diagram("Azure Web Service", show=False):
    lb = LoadBalancers("load balancer")
    web_server = AppServices("web server")
    storage = StorageAccounts("azure storage")

    lb >> web_server >> storage