#!/bin/bash

# Consul Quick Start Script
# This script helps you quickly start and verify your Consul cluster

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== Consul Quick Start ===${NC}\n"

# Check if docker-compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo -e "${RED}Error: docker-compose is not installed${NC}"
    exit 1
fi

# Function to wait for Consul to be ready
wait_for_consul() {
    local max_attempts=30
    local attempt=1
    
    echo -e "${YELLOW}Waiting for Consul to be ready...${NC}"
    while [ $attempt -le $max_attempts ]; do
        if docker exec consul-server1 consul members &> /dev/null; then
            echo -e "${GREEN}Consul is ready!${NC}"
            return 0
        fi
        echo -n "."
        sleep 2
        ((attempt++))
    done
    
    echo -e "${RED}Consul failed to start within timeout${NC}"
    return 1
}

# Start the cluster
echo -e "${YELLOW}1. Starting Consul cluster...${NC}"
docker-compose up -d

# Wait for Consul to be ready
if ! wait_for_consul; then
    exit 1
fi

echo ""
echo -e "${YELLOW}2. Checking cluster members...${NC}"
docker exec consul-server1 consul members

echo ""
echo -e "${YELLOW}3. Checking leader status...${NC}"
docker exec consul-server1 consul operator raft list-peers

echo ""
echo -e "${GREEN}=== Consul Cluster Started Successfully! ===${NC}\n"

echo -e "${GREEN}Available Services:${NC}"
echo "  • Consul UI (Server 1): http://localhost:8500"
echo "  • Consul UI (Server 2): http://localhost:8501"
echo "  • Consul UI (Server 3): http://localhost:8502"
echo "  • Consul Client:        http://localhost:8503"
echo "  • Prometheus:           http://localhost:9090"
echo "  • Grafana:              http://localhost:3000 (admin/admin)"

echo ""
echo -e "${GREEN}Quick Commands:${NC}"
echo "  # View cluster members"
echo "  docker exec consul-server1 consul members"
echo ""
echo "  # Register a test service"
echo "  docker exec consul-server1 consul services register -name=test -port=8080"
echo ""
echo "  # Query services"
echo "  docker exec consul-server1 consul catalog services"
echo ""
echo "  # Check cluster health"
echo "  docker exec consul-server1 consul operator raft list-peers"
echo ""
echo "  # View logs"
echo "  docker logs -f consul-server1"
echo ""
echo "  # Stop cluster"
echo "  docker-compose down"
echo ""
echo -e "${YELLOW}For more examples, check the notes/ directory!${NC}"

