#!/bin/bash

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts


helm install prometheus prometheus-community/prometheus --version 22.6.2
