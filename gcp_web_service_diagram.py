from diagrams import Diagram
from diagrams.gcp.compute import ComputeEngine
from diagrams.gcp.network import LoadBalancing
from diagrams.gcp.storage import Storage

with Diagram("GCP Web Service", show=False):
    lb = LoadBalancing("load balancer")
    web_server = ComputeEngine("web server")
    storage = Storage("cloud storage")

    lb >> web_server >> storage