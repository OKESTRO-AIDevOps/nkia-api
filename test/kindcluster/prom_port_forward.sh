#!/bin/bash


kubectl port-forward svc/prometheus-server 9090:80 --address='0.0.0.0'