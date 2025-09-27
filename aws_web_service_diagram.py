from diagrams import Diagram, Cluster, Edge
from diagrams.aws.compute import EC2
from diagrams.aws.database import RDS
from diagrams.aws.network import ELB

# Cấu hình tổng thể cho sơ đồ
graph_attr = {
    "fontsize": "16",
    "bgcolor": "white",
    "pad": "0.5",
    "rankdir": "LR"
}

node_attr = {
    "style": "filled",
    "shape": "box",
    "fillcolor": "lightblue",
    "fontsize": "12"
}

edge_attr = {
    "color": "gray",
    "fontsize": "10"
}

# Bắt đầu vẽ Diagram
with Diagram(
    "Simple Web Service Architecture",
    show=False,
    filename="simple_web_service",
    graph_attr=graph_attr,
    node_attr=node_attr,
    edge_attr=edge_attr
):
    # Load Balancer ngoài Internet
    lb = ELB("Load Balancer")

    # Cụm Public Subnet
    with Cluster("Public Subnet"):
        web_servers = [EC2("Web 1"), EC2("Web 2")]

    # Cụm Database Subnet
    with Cluster("Private Subnet"):
        db = RDS("Primary DB")

    # Kết nối giữa các thành phần
    lb >> Edge(color="blue") >> web_servers
    web_servers >> Edge(color="green", style="dashed") >> db