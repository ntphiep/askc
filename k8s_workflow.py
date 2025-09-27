from diagrams import Cluster, Diagram, Edge
from diagrams.k8s.clusterconfig import HPA
from diagrams.k8s.compute import Pod, Deployment
from diagrams.k8s.controlplane import API, Scheduler, ControllerManager
from diagrams.k8s.storage import PV, PVC
from diagrams.k8s.network import Service, Ingress
from diagrams.onprem.client import User

# Diagram attributes for better visualization
graph_attr = {
    "fontsize": "12",
    "bgcolor": "white",
    "pad": "0.5",
    "rankdir": "TB"  # Top to Bottom direction
}

with Diagram("Kubernetes Architecture Workflow", show=False, direction="TB", graph_attr=graph_attr):
    # External User (kubectl or UI)
    user = User("User (kubectl/UI)")

    # Control Plane Cluster
    with Cluster("Control Plane"):
        api_server = API("API Server")
        etcd = PV("etcd")
        scheduler = Scheduler("Scheduler")
        controller_manager = ControllerManager("Controller Manager")

    # Worker Nodes Cluster
    with Cluster("Worker Nodes"):
        with Cluster("Node 1"):
            deployment_1 = Deployment("App Deployment")
            pods_1 = [Pod("Pod 1"), Pod("Pod 2")]
            service_1 = Service("Service")
            hpa_1 = HPA("HPA")

        with Cluster("Node 2"):
            deployment_2 = Deployment("App Deployment")
            pods_2 = [Pod("Pod 3")]
            service_2 = Service("Service")
            hpa_2 = HPA("HPA")

    # Networking
    ingress = Ingress("Ingress")

    # Persistent Storage
    pvc = PVC("PVC")
    pv = PV("Persistent Volume")

    # Workflow Connections
    # Step 1: User sends request to API Server
    user >> Edge(label="API Request") >> api_server

    # Step 2: API Server interacts with etcd to store/retrieve state
    api_server >> Edge(label="Store/Retrieve State") >> etcd

    # Step 3: API Server notifies Controller Manager
    api_server >> Edge(label="Process Request") >> controller_manager

    # Step 4: Controller Manager coordinates with Scheduler
    controller_manager >> Edge(label="Schedule Pods") >> scheduler

    # Step 5: Scheduler assigns Pods to Nodes via Deployments
    scheduler >> Edge(label="Assign to Node 1") >> deployment_1
    scheduler >> Edge(label="Assign to Node 2") >> deployment_2

    # Step 6: Deployments manage Pods
    deployment_1 >> pods_1
    deployment_2 >> pods_2

    # Step 7: Services expose Pods
    pods_1 >> Edge(label="Expose") >> service_1
    pods_2 >> Edge(label="Expose") >> service_2

    # Step 8: Ingress routes traffic to Services
    api_server >> Edge(label="Configure Routing") >> ingress
    ingress >> Edge(label="Route Traffic") >> [service_1, service_2]

    # Step 9: Horizontal Pod Autoscaler adjusts Deployments
    service_1 >> Edge(label="Metrics") >> hpa_1
    service_2 >> Edge(label="Metrics") >> hpa_2
    hpa_1 >> Edge(label="Scale") >> deployment_1
    hpa_2 >> Edge(label="Scale") >> deployment_2

    # Step 10: Pods use Persistent Storage
    pods_1 >> Edge(label="Claim") >> pvc
    pods_2 >> Edge(label="Claim") >> pvc
    pvc >> Edge(label="Bind") >> pv

    # Step 11: Pods report status back to API Server
    pods_1 >> Edge(label="Status Update") >> api_server
    pods_2 >> Edge(label="Status Update") >> api_server